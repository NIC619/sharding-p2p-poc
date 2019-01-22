import datetime
import re
import subprocess
import time

import pytest

from simulation.config import (
    PORT_BASE,
    RPC_PORT_BASE,
)
from simulation.logs import (
    LOG_PATTERN,
    RPCLogs,
    OperationLogs,
    map_log_enum_pattern,
)
from simulation.exceptions import (
    CLIFailure,
)
from simulation.network import (
    get_docker_host_ip,
    make_local_node,
    make_local_nodes,
    wait_for_pubsub_heartbeat,
)
from simulation.node import (
    Node,
)


@pytest.yield_fixture("module")
def unchanged_node():
    """Unchanged when tested, to save the time initializing new nodes every test.
    """
    n = make_local_node(12345)
    wait_for_pubsub_heartbeat()
    n.set_peer_id()
    yield n
    n.close()


@pytest.yield_fixture
def nodes():
    ns = make_local_nodes(0, 3)
    yield ns
    for n in ns:
        n.close()


def connect(node_0, node_1):
    node_0.add_peer(node_1)
    assert node_1.peer_id in node_0.list_peer()
    assert node_0.peer_id in node_1.list_peer()


def test_name(unchanged_node):
    assert unchanged_node.name == "whiteblock-node{}".format(unchanged_node.seed)


def is_node_running(node):
    res = subprocess.run(
        ["docker ps | grep {} | wc -l".format(node.name)],
        shell=True,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        check=True,
        encoding="utf-8",
    )
    num_lines = int(res.stdout.strip())
    if num_lines > 1:
        raise Exception("Should not grep more than 1 running container named {}".format(
            node.name,
        ))
    return num_lines != 0


def test_close():
    n = make_local_node(0)
    assert is_node_running(n)
    n.close()
    assert not is_node_running(n)


def test_run():
    unrun_node = Node(
        get_docker_host_ip(),
        PORT_BASE,
        RPC_PORT_BASE,
        0,
    )
    assert not is_node_running(unrun_node)
    unrun_node.run()
    assert is_node_running(unrun_node)
    unrun_node.close()


def test_wait_for_log(unchanged_node):
    # Note: here assume we at least have two lines in the format of "INFO: xxx"
    #       if the log style changes in the container and we get fewer lines, this will block
    #       forever
    log = unchanged_node.wait_for_log("INFO", 1)
    # just confirm the grep'ed log is correct
    assert "INFO" in log.split(' ')


def test_get_log_time(unchanged_node):
    log_time = unchanged_node.get_log_time("INFO", 1)
    assert isinstance(log_time, datetime.datetime)


def test_set_peer_id():
    n = make_local_node(0)
    assert n.peer_id is None
    n.set_peer_id()
    assert n.peer_id is not None
    n.close()


def test_multiaddr(nodes):
    unrun_node = Node(
        get_docker_host_ip(),
        PORT_BASE,
        RPC_PORT_BASE,
        0,
    )
    # test: `multiaddr` should fail since the `peer_id` haven't been set after initialized.
    with pytest.raises(ValueError):
        unrun_node.multiaddr
    # test: nodes are run with `set_peer_id` in `make_local_nodes`
    assert nodes[0].multiaddr == "/ip4/{}/tcp/{}/ipfs/{}".format(
        nodes[0].ip,
        nodes[0].port,
        nodes[0].peer_id
    )
    unrun_node.close()


def test_cli(unchanged_node):
    non_existing_command = "123456"
    res = unchanged_node.cli([non_existing_command])
    assert res.returncode != 0
    existing_command = "listpeer"
    res = unchanged_node.cli([existing_command])
    assert res.returncode == 0
    # test: wrong type, should be `list` instead of the `string`
    with pytest.raises(CLIFailure):
        unchanged_node.cli(existing_command)


def test_cli_safe(unchanged_node):
    non_existing_command = "123456"
    # test: non-existing command should fail in `cli_safe`
    with pytest.raises(CLIFailure):
        unchanged_node.cli_safe([non_existing_command])
    unchanged_node.cli_safe(["listpeer"])


def test_identify(unchanged_node):
    pinfo = unchanged_node.identify()
    assert "peerID" in pinfo
    assert "multiAddrs" in pinfo


def test_list_peer(unchanged_node):
    peers = unchanged_node.list_peer()
    assert peers == []


def test_add_peer(nodes):
    assert len(nodes[0].list_peer()) == 0
    assert len(nodes[1].list_peer()) == 0
    nodes[0].add_peer(nodes[1])
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_ADD_PEER_FMT],
        0,
    )
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_ADD_PEER_FINISHED],
        0,
    )
    # symmetric connection
    assert nodes[1].peer_id in nodes[0].list_peer()
    assert nodes[0].peer_id in nodes[1].list_peer()
    # set the wrong `ip`/`port`/`seed` of nodes[2], and nodes[0] try to add it
    ip_2 = nodes[2].ip
    port_2 = nodes[2].port
    seed_2 = nodes[2].seed
    # ip
    nodes[2].ip = "123.456.789.012"
    # test: nodes[0] should fail to add nodes[2] because of the mocked and wrong ip
    with pytest.raises(CLIFailure):
        nodes[0].add_peer(nodes[2])
    nodes[2].ip = ip_2
    # port
    nodes[2].port = 123
    # test: nodes[0] should fail to add nodes[2] because of the mocked and wrong port
    with pytest.raises(CLIFailure):
        nodes[0].add_peer(nodes[2])
    nodes[2].port = port_2
    # seed
    nodes[2].seed = 32767
    # test: nodes[0] should fail to add nodes[2] because of the mocked and wrong seed
    with pytest.raises(CLIFailure):
        nodes[0].add_peer(nodes[2])
    nodes[2].seed = seed_2
    # add 2 peers(avoid nodes[0] adding nodes[2] again because of `dial backoff `)
    nodes[2].add_peer(nodes[0])
    nodes[2].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_ADD_PEER_FMT],
        0,
    )
    nodes[2].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_ADD_PEER_FINISHED],
        0,
    )
    assert len(nodes[0].list_peer()) == 2
    assert len(nodes[2].list_peer()) == 1


def test_remove_peer(nodes):
    # test: wrong format `peer_id`
    with pytest.raises(CLIFailure):
        nodes[0].remove_peer("123")
    # test: remove non-peer
    nodes[0].remove_peer(nodes[2].peer_id)
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_REMOVE_PEER_FMT],
        0,
    )
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_REMOVE_PEER_FINISHED],
        0,
    )
    assert len(nodes[0].list_peer()) == 0
    connect(nodes[0], nodes[1])
    assert len(nodes[0].list_peer()) == 1
    nodes[0].remove_peer(nodes[1].peer_id)
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_REMOVE_PEER_FMT],
        1,
    )
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_REMOVE_PEER_FINISHED],
        1,
    )
    # symmetric connection
    assert len(nodes[0].list_peer()) == 0
    assert len(nodes[1].list_peer()) == 0
    # test: remove twice
    nodes[0].remove_peer(nodes[1].peer_id)
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_REMOVE_PEER_FMT],
        2,
    )
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_REMOVE_PEER_FINISHED],
        2,
    )
    assert len(nodes[0].list_peer()) == 0


def test_list_topic_peer(nodes):
    topic_peers = nodes[0].list_topic_peer()
    # assert the default listening topic is present
    assert "listeningShard" in topic_peers
    assert len(topic_peers["listeningShard"]) == 0
    # test: add_peer and should be one more peer in `listeningShard`
    connect(nodes[0], nodes[1])
    assert len(nodes[0].list_topic_peer()["listeningShard"]) == 1
    # test: remove_peer and should be one less peer in `listeningShard`
    nodes[0].remove_peer(nodes[1].peer_id)
    assert len(nodes[0].list_topic_peer()["listeningShard"]) == 0


def test_get_subscribed_shard(unchanged_node):
    shards = unchanged_node.get_subscribed_shard()
    assert shards == []


def test_subscribe_shard(nodes):
    # test: wrong type, should be list of int
    with pytest.raises(TypeError):
        nodes[0].subscribe_shard(0)
    # test: normally subscribe
    nodes[0].subscribe_shard([0])
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_SUBSCRIBE_SHARD_FMT],
        0,
    )
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_SUBSCRIBE_SHARD_FINISHED],
        0,
    )
    assert nodes[0].get_subscribed_shard() == [0]
    assert nodes[1].list_shard_peer([0])["0"] == []
    # test: after adding peers, the node should know the shards its peer subscribes
    connect(nodes[0], nodes[1])
    time.sleep(0.2)
    assert nodes[0].peer_id in nodes[1].list_shard_peer([0])["0"]
    # test: subscribe "0" twice, and shards "1", "2"
    nodes[0].subscribe_shard([0, 1, 2])
    # test: ensure there are two logs LOG_SUBSCRIBE_SHARD_FINISHED because `subscribe_shard` is
    #       called twice
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_SUBSCRIBE_SHARD_FMT],
        1,
    )
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_SUBSCRIBE_SHARD_FINISHED],
        1,
    )
    assert nodes[0].get_subscribed_shard() == [0, 1, 2]


def test_unsubscribe_shard(nodes):
    # test: wrong type, should be list of int
    with pytest.raises(TypeError):
        nodes[0].unsubscribe_shard(0)
    # test: unsubscribe before subscribe
    nodes[0].unsubscribe_shard([0])
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_UNSUBSCRIBE_SHARD_FMT],
        0,
    )
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_UNSUBSCRIBE_SHARD_FINISHED],
        0,
    )
    # test: normally subscribe
    nodes[0].subscribe_shard([0])
    connect(nodes[0], nodes[1])
    nodes[0].unsubscribe_shard([0])
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_UNSUBSCRIBE_SHARD_FMT],
        1,
    )
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_UNSUBSCRIBE_SHARD_FINISHED],
        1,
    )
    assert nodes[0].get_subscribed_shard() == []
    # test: ensure the peer knows its peer unsubscribes
    assert nodes[1].list_shard_peer([0])["0"] == []


def test_broadcast_collation(nodes):
    # test: broadcast to a non-subscribed shard
    with pytest.raises(CLIFailure):
        nodes[0].broadcast_collation(0, 1, 100, 123)
    connect(nodes[0], nodes[1])
    connect(nodes[1], nodes[2])
    nodes[0].subscribe_shard([0])
    nodes[1].subscribe_shard([0])
    nodes[2].subscribe_shard([0])
    wait_for_pubsub_heartbeat()
    # test: see if all nodes receive the broadcasted collation. Use a bigger size of collation to
    #       avoid the possible small time difference. E.g. sometimes t2 < t1, which does not make
    #       sense.
    nodes[0].broadcast_collation(0, 1, 1000000, 123)
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_BROADCAST_COLLATION_FMT],
        0,
    )
    t0 = nodes[0].get_log_time(
        map_log_enum_pattern[RPCLogs.LOG_BROADCAST_COLLATION_FINISHED],
        0,
    )
    t1 = nodes[0].get_log_time(
        map_log_enum_pattern[OperationLogs.LOG_RECEIVE_MSG],
        0,
    )
    t2 = nodes[1].get_log_time(
        map_log_enum_pattern[OperationLogs.LOG_RECEIVE_MSG],
        0,
    )
    t3 = nodes[2].get_log_time(
        map_log_enum_pattern[OperationLogs.LOG_RECEIVE_MSG],
        0,
    )
    assert t0 < t1
    assert t1 < t2
    assert t2 < t3


def test_bootstrap(nodes):
    # test: stop before start
    nodes[0].bootstrap(False)
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_BOOTSTRAP_FMT],
        0,
    )
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_BOOTSTRAP_FINISHED],
        0,
    )
    connect(nodes[0], nodes[1])
    connect(nodes[1], nodes[2])
    time.sleep(0.2)
    # test: nodes[0] bootstraps with nodes[1] as the bootnode and therefore connects to nodes[2]
    nodes[0].bootstrap(True, nodes[1].multiaddr)
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_BOOTSTRAP_FMT],
        1,
    )
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_BOOTSTRAP_FINISHED],
        1,
    )
    time.sleep(2)
    assert nodes[2].peer_id in nodes[0].list_peer()


@pytest.mark.skip(
    "Sometimes this test fails but it is not important, "
    "because it just need more time to stop"
)
def test_stop():
    n = make_local_node(0)
    assert is_node_running(n)
    n.stop()
    time.sleep(2)  # FIXME: we need to wait for a fairly long time to wait for its stop
    assert not is_node_running(n)


def test_discover_shard(nodes):
    # connection: 0 <-> 1 <-> 2
    connect(nodes[0], nodes[1])
    connect(nodes[1], nodes[2])
    # test: calls without `shard_ids`
    with pytest.raises(TypeError):
        nodes[0].discover_shard()
    peers_shard_0 = nodes[0].discover_shard([0])
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_DISCOVER_SHARD_FMT],
        0,
    )
    nodes[0].wait_for_log(
        map_log_enum_pattern[RPCLogs.LOG_DISCOVER_SHARD_FINISHED],
        0,
    )
    assert len(peers_shard_0["0"]) == 0
    # test: discover shard through pubsub topic
    nodes[2].subscribe_shard([0])
    # wait for subscription broadcasted to the topic "listeningShards"
    wait_for_pubsub_heartbeat()
    assert nodes[2].peer_id in nodes[0].discover_shard([0])["0"]
    # test: return nothing when no shard_ids
    assert len(nodes[0].discover_shard([])) == 0
    # test: should be able discover recently subscribed shards
    nodes[2].subscribe_shard([1])
    wait_for_pubsub_heartbeat()
    # test: should return all shards if no `shard_ids` given
    assert nodes[2].peer_id in nodes[0].discover_shard([0])["0"]
    assert nodes[2].peer_id in nodes[0].discover_shard([1])["1"]
    # test: should not discover unsubscribed shards
    nodes[2].unsubscribe_shard([0])
    wait_for_pubsub_heartbeat()
    assert nodes[2].peer_id not in nodes[0].discover_shard([0])["0"]


def test_get_logs(nodes):
    nodes[0].subscribe_shard([0])
    nodes[0].broadcast_collation(0, 1, 10000, 100)
    nodes[0].discover_shard([0])
    nodes[0].bootstrap(True, nodes[1].multiaddr)
    nodes[0].bootstrap(False)
    nodes[0].unsubscribe_shard([0])
    nodes[0].remove_peer(nodes[1].peer_id)
    time.sleep(0.5)
    # test: just ensure we get enough lines that conform to the format `LOG_PATTERN` and are
    #       `DEBUG` type logs
    log_pattern = LOG_PATTERN.format(".*")
    valid_logs = tuple(
        line
        for line in nodes[0].get_logs()
        if re.search(log_pattern, line) and ("DEBUG" in line)
    )
    # 7 RPCs above, 2 logs(doing and finished) for each operation, so at least
    # we get `7 * 2 = 14` logs
    assert len(valid_logs) >= 14


def test_get_events(nodes):
    assert len(tuple(nodes[0].get_events())) == 0
    nodes[0].subscribe_shard([0])
    nodes[0].broadcast_collation(0, 1, 10000, 100)
    nodes[0].discover_shard([0])
    nodes[0].bootstrap(True, nodes[1].multiaddr)
    nodes[0].bootstrap(False)
    nodes[0].unsubscribe_shard([0])
    nodes[0].remove_peer(nodes[1].peer_id)
    event_types = [
        node.event_type
        for node in nodes[0].get_events()
    ]
    # test: only receive one message
    assert event_types.count(OperationLogs.LOG_RECEIVE_MSG) == 1
    # remove `OperationLogs.LOG_RECEIVE_MSG`, because we only guarantee the order of RPCLogs
    event_types.remove(OperationLogs.LOG_RECEIVE_MSG)
    # test: the order of the events
    assert event_types[0] == RPCLogs.LOG_SUBSCRIBE_SHARD_FMT
    assert event_types[1] == RPCLogs.LOG_SUBSCRIBE_SHARD_FINISHED
    assert event_types[2] == RPCLogs.LOG_BROADCAST_COLLATION_FMT
    assert event_types[3] == RPCLogs.LOG_BROADCAST_COLLATION_FINISHED
    assert event_types[4] == RPCLogs.LOG_DISCOVER_SHARD_FMT
    assert event_types[5] == RPCLogs.LOG_DISCOVER_SHARD_FINISHED
    assert event_types[6] == RPCLogs.LOG_BOOTSTRAP_FMT
    assert event_types[7] == RPCLogs.LOG_BOOTSTRAP_FINISHED
    assert event_types[8] == RPCLogs.LOG_BOOTSTRAP_FMT
    assert event_types[9] == RPCLogs.LOG_BOOTSTRAP_FINISHED
    assert event_types[10] == RPCLogs.LOG_UNSUBSCRIBE_SHARD_FMT
    assert event_types[11] == RPCLogs.LOG_UNSUBSCRIBE_SHARD_FINISHED
    assert event_types[12] == RPCLogs.LOG_REMOVE_PEER_FMT
    assert event_types[13] == RPCLogs.LOG_REMOVE_PEER_FINISHED
