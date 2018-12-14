#!/usr/bin/env python

from datetime import (
    datetime,
)
import json
import logging
import re
import subprocess
import sys
import threading
import os
import time

from utils import (
    connect_nodes,
    ensure_topology,
    get_actual_topology,
    kill_nodes,
    make_barbell_topology,
    make_local_nodes,
)

# --i <image>
# --n <number of nodes>
# --I <interface> eno4
# example:
# os.system("./umba --i sharding --I eno4 --n 20")
# os.system("docker exec -it whiteblock-node0 -port=8080")
# ip eample: 10.1.0.2 would be whiteblock-node0
# 10.1.0.6 would be whiteblock-node1 ++
# just add the flags and commands that are needed
# we will handle logic of collecting data and configuring umba

#TEST SERIES A
#LATENCY


def test_time_broadcasting_data_single_shard():
    num_nodes = 30
    num_collations = 1
    collation_size = 1000000  # 1MB
    collation_time = 50  # broadcast 1 collation every 50 milliseconds

    print("Spinning up {} nodes...".format(num_nodes), end='')
    nodes = make_local_nodes(0, num_nodes)
    print("done")
    print("Connecting nodes...", end='')
    topo = make_barbell_topology(nodes)
    connect_nodes(nodes, topo)
    print("done")
    print("Checking the connections...", end='')
    ensure_topology(nodes, topo)
    print("done")

    for node in nodes:
        node.subscribe_shard([0])
    broadcasting_node = 0
    print(
        "Broadcasting {} collations(size={} bytes) from node{} in the barbell topology...".format(
            num_collations,
            collation_size,
            broadcasting_node,
        ),
        end='',
    )
    nodes[broadcasting_node].broadcast_collation(0, num_collations, collation_size, collation_time)  # 1MB
    print("done")

    print("Gathering time...", end='')
    time_broadcast = nodes[0].get_log_time('rpcserver:BroadcastCollation: finished', 0)
    time_received = nodes[-1].get_log_time(
        'Validating the received message',
        num_collations - 1,
    )
    print("done")
    print(
        "time to broadcast all data to the last node: \x1b[0;37m{}\x1b[0m".format(
            time_received - time_broadcast,
        )
    )
    print("Cleaning up the nodes...", end='')
    kill_nodes(nodes)
    print("done")


def test_joining_through_bootnodes():
    num_bootnodes = 1
    num_normal_nodes = 10
    print("Spinning up {} bootnodes...".format(num_bootnodes), end='')
    bootnodes = make_local_nodes(0, num_bootnodes)
    print("done")
    print("Connecting bootnodes...", end='')
    topo = make_barbell_topology(bootnodes)
    connect_nodes(bootnodes, topo)
    print("done")
    print("Checking the connections...", end='')
    ensure_topology(bootnodes, topo)
    print("done")

    bootnodes_multiaddr = [node.multiaddr for node in bootnodes]
    print("Spinning up {} nodes...".format(num_normal_nodes), end='')
    nodes = make_local_nodes(num_bootnodes, num_bootnodes + num_normal_nodes, bootnodes_multiaddr)
    print("done")

    print("Sleeping for seconds...", end='')
    time.sleep(3)
    print("done")

    all_nodes = bootnodes + nodes
    actual_topo = get_actual_topology(all_nodes)
    print("actual_topo =", actual_topo)

    print("Cleaning up the nodes...", end='')
    kill_nodes(all_nodes)
    print("done")


def test_reproduce_bootstrapping_issue():
    num_bootnodes = 1
    num_normal_nodes = 5
    print("Spinning up {} bootnodes...".format(num_bootnodes), end='')
    bootnodes = make_local_nodes(0, num_bootnodes)
    print("done")
    bootnodes_multiaddr = [node.multiaddr for node in bootnodes]
    print("Spinning up {} nodes...".format(num_normal_nodes), end='')
    nodes = make_local_nodes(num_bootnodes, num_bootnodes + num_normal_nodes, bootnodes_multiaddr)
    print("done")

    print("Sleeping for seconds...", end='')
    time.sleep(3)
    print("done")

    all_nodes = bootnodes + nodes
    for node in all_nodes:
        node.subscribe_shard([1])
    for node in all_nodes:
        peers = node.list_peer()
        print(f"{node}: peers={peers}")
        topic_peers = node.list_topic_peer([])
        print(f"{node}: topic_peers={topic_peers}")

    kill_nodes(bootnodes + nodes)


if __name__ == "__main__":
    test_time_broadcasting_data_single_shard()
    test_joining_through_bootnodes()
    test_reproduce_bootstrapping_issue()
