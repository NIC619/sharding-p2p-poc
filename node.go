package main

import (
	"context"
	"time"

	host "github.com/libp2p/go-libp2p-host"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type Node struct {
	host.Host        // lib-p2p host
	discovery        Discovery
	*RequestProtocol // for peers to request data
	*ShardManager

	ctx context.Context

	dht                 *kaddht.IpfsDHT
	doBootstrapping     bool
	cancelBootstrapping context.CancelFunc
}

// NewNode creates a new node with its implemented protocols
func NewNode(ctx context.Context, h host.Host, dht *kaddht.IpfsDHT, eventNotifier EventNotifier) *Node {
	shardPrefTable := NewShardPrefTable()
	pubsubService, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		logger.Fatalf("Failed to create new pubsub service, err: %v", err)
	}
	node := &Node{
		Host:            h,
		discovery:       NewGlobalTable(ctx, h, pubsubService, shardPrefTable),
		ctx:             ctx,
		dht:             dht,
		doBootstrapping: false,
	}
	node.RequestProtocol = NewRequestProtocol(node)
	node.ShardManager = NewShardManager(ctx, node, pubsubService, eventNotifier, node.discovery, shardPrefTable)

	return node
}

// TODO: should be changed to `Knows` and `HasConnections`
func (n *Node) IsPeer(peerID peer.ID) bool {
	for _, value := range n.Peerstore().Peers() {
		if value == peerID {
			return true
		}
	}
	return false
}

//StartBootstrapping starts the bootstrapping process in dht, with the contact peers
// `bootstrapPeers`.
func (n *Node) StartBootstrapping(ctx context.Context, bootstrapPeers []pstore.PeerInfo) error {
	if n.doBootstrapping {
		return nil
	}
	// try to connect to the chosen nodes
	bootstrapConnect(ctx, n, bootstrapPeers)

	bootstrapCtx, cancel := context.WithCancel(ctx)
	// err := n.dht.Bootstrap(bootstrapCtx)
	var cfg kaddht.BootstrapConfig
	cfg = kaddht.DefaultBootstrapConfig
	cfg.Timeout = time.Duration(1 * time.Second)
	proc, err := n.dht.BootstrapWithConfig(cfg)
	if err != nil {
		return err
	}
	// wait till ctx or dht.Context exits.
	// we have to do it this way to satisfy the Routing interface (contexts)
	go func() {
		defer proc.Close()
		select {
		case <-bootstrapCtx.Done():
		case <-n.dht.Context().Done():
		}
	}()

	n.doBootstrapping = true
	n.cancelBootstrapping = cancel
	return nil
}

//IsBootstrapping indicates if the node is bootstrapping
func (n *Node) IsBootstrapping() bool {
	return n.doBootstrapping
}

//StopBootstrapping stops the bootstrapping process, using the cancel function results from
// `StartBootstrapping`
func (n *Node) StopBootstrapping() error {
	if !n.doBootstrapping {
		return nil
	}
	n.cancelBootstrapping()
	n.doBootstrapping = false
	n.cancelBootstrapping = nil
	return nil
}
