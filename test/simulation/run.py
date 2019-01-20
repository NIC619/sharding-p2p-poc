import functools
import logging
import math
import time
from random import randint

from simulation.logs import (
    LOG_BROADCAST_COLLATION_FINISHED,
    LOG_RECEIVE_MSG,
)
from simulation.network import (
    Network,
)


def get_degree_map(topology):
    node_degree_map = {}
    for con in topology:
        for i in range(2):
            cur_degree = node_degree_map.get(con[i])
            if cur_degree is not None:
                node_degree_map[con[i]] = cur_degree + 1
            else:
                node_degree_map[con[i]] = 1
    return sorted(node_degree_map.items())


def visualize(topology, file_name):
    import networkx as nx
    G = nx.Graph()
    for con in topology:
        G.add_edge(con[0], con[1])

    import matplotlib.pyplot as plt
    pos_nodes = nx.spring_layout(G)
    nx.draw(G, pos_nodes, with_labels=True)

    pos_attrs = {node: (coords[0], coords[1] + 0.08) for node, coords in pos_nodes.items()}

    # node_attrs = nx.get_node_attributes(G, 'shard_id')
    # custom_node_attrs = {node: attr for node, attr in node_attrs.items()}
    # nx.draw_networkx_labels(G, pos_attrs, labels=custom_node_attrs)

    plt.savefig(file_name)


def trim_connections(network, target_conns_ratio):
    actual_topo = list(network.get_actual_topology())
    print(len(actual_topo))
    normal_nodes_count = len(network.normal_nodes)
    max_normal_nodes_conns = (normal_nodes_count**2 - normal_nodes_count)/2
    total_normal_nodes_conns = len(actual_topo) - normal_nodes_count
    if total_normal_nodes_conns / max_normal_nodes_conns <= target_conns_ratio:
        # Connections ratio is already below target ratio
        return
    else:
        trimmed_conns_count = 0
        while (total_normal_nodes_conns - trimmed_conns_count) / max_normal_nodes_conns > target_conns_ratio:
            print("ratio:", (total_normal_nodes_conns - trimmed_conns_count) / max_normal_nodes_conns)
            index_to_trim = randint(0, len(actual_topo)-1)
            if actual_topo[index_to_trim][0] == 0:
                continue
            else:
                network.nodes[actual_topo[index_to_trim][0]].remove_peer(network.nodes[actual_topo[index_to_trim][1]].peer_id)
                trimmed_conns_count += 1
                print("trim conn", actual_topo[index_to_trim])
                del actual_topo[index_to_trim]
    print(len(actual_topo))


def static_nodes_subscribe(static_nodes, num_shards, avg_sub_per_node):
    assert avg_sub_per_node <= num_shards
    for i, node in enumerate(static_nodes):
        inc = i % avg_sub_per_node - avg_sub_per_node // 2
        sub_shard_list = []
        for _ in range(avg_sub_per_node + inc):
            shard_id = randint(0, num_shards-1)
            while shard_id in sub_shard_list:
                shard_id = randint(0, num_shards-1)
            sub_shard_list.append(shard_id)
        node.subscribe_shard(sub_shard_list)
        print(node.peer_id, " subscribe to", sub_shard_list)


def test_decor(test_func):
    @functools.wraps(test_func)
    def func(*args, **kwargs):
        total_len = 100
        separator = "="
        if len(test_func.__name__) > (total_len - 2):
            start_kanban = "{} {}".format(separator * 30, test_func.__name__)
        else:
            num_separators = total_len - (len(test_func.__name__) + 2)
            start_kanban = "{} {} {}".format(
                separator * (num_separators // 2),
                test_func.__name__,
                separator * ((num_separators // 2 + 1) if num_separators % 2 != 0 else num_separators // 2),  # noqa: E501
            )
        print(start_kanban)
        t_start = time.time()
        res = test_func(*args, **kwargs)
        print("{} execution time: {} seconds".format(test_func.__name__, time.time() - t_start))
        end_kanban = separator * total_len
        print(end_kanban)
        return res
    return func


@test_decor
def test_time_broadcasting_data_single_shard():
    num_collations = 10
    collation_size = 1000000  # 1MB
    collation_time = 50  # broadcast 1 collation every 50 milliseconds
    percent = 0.9

    n = Network(
        num_bootnodes=0,
        num_normal_nodes=30,
        topology_option=Network.topology_option.BARBELL,
    )

    nodes = n.nodes
    for node in nodes:
        node.subscribe_shard([0])

    print("Sleeping for seconds...", end='')
    time.sleep(3)
    print("done")

    broadcasting_node = 0
    print(
        "Broadcasting {} collations(size={} bytes) from node{} in the barbell topology...".format(
            num_collations,
            collation_size,
            broadcasting_node,
        ),
        end='',
    )
    nodes[broadcasting_node].broadcast_collation(0, num_collations, collation_size, collation_time)
    print("done")

    # TODO: maybe we can have a list of broadcasted data, and use them to grep in the nodes' logs
    #       for precision, instead of using only numbers
    # wait until all nodes receive the broadcasted data, and gather the time
    print("Gathering time...", end='')
    time_broadcast = nodes[broadcasting_node].get_log_time(LOG_BROADCAST_COLLATION_FINISHED, 0)
    time_received_list = []
    for i in range(len(nodes)):
        time_received = nodes[i].get_log_time(LOG_RECEIVE_MSG, num_collations - 1,)
        time_received_list.append(time_received)
    print("done")

    # sort the time, find the last node in the first percent% nodes who received the data
    time_received_sorted = sorted(time_received_list, key=lambda t: t)
    index_last = math.ceil(len(nodes) * percent - 1)

    print(
        "time to broadcast all data to {} percent nodes: \x1b[0;37m{}\x1b[0m".format(
            percent,
            time_received_sorted[index_last] - time_broadcast,
        )
    )


@test_decor
def test_joining_through_bootnodes():
    n = Network(
        num_bootnodes=1,
        num_normal_nodes=10,
        topology_option=Network.topology_option.NONE,
    )

    print("Sleeping for seconds...", end='')
    time.sleep(3)
    print("done")

    actual_topo = n.get_actual_topology()
    print("actual_topo =", actual_topo)


@test_decor
def test_reproduce_bootstrapping_issue():
    n = Network(
        num_bootnodes=1,
        num_normal_nodes=5,
        topology_option=Network.topology_option.NONE,
    )

    all_nodes = n.nodes
    for node in all_nodes:
        node.subscribe_shard([1])

    print("Sleeping for seconds...", end='')
    time.sleep(3)
    print("done")

    for node in all_nodes:
        peers = node.list_peer()
        topic_peers = node.list_topic_peer([])
        print("{}: summary: len(peers)={}, len_per_topic={}".format(
            node,
            len(peers),
            {key: len(value) for key, value in topic_peers.items()},
        ))
        print(f"{node}: peers={peers}")
        print(f"{node}: topic_peers={topic_peers}")


@test_decor
def test_plan():
    """
    Steps:
        - Build network
        - Provision nodes
        - Configure network conditions between nodes according to specified test case
        - Configure actions and behavior between nodes according to specified test case
        - Output performance data in CSV format
        - Aggregate data, parse, & present data
        - Push data to appropriate repo
        - Reset environment
    Metrics:
        - Subscription Time
        - Discovery Time
        - Message Propagation Time
        - Shard Propagation Time
    """
    # variables
    num_validators = 20
    num_static_nodes = 20
    num_shards = 10
    message_size = 0
    propagation_time = 0
    # TODO: the following variables are not controllable now
    bandwidth = 0
    network_latency = 0
    packet_loss = 0

    # Build network
    n = Network(
        num_bootnodes=1,
        num_normal_nodes=num_validators+num_static_nodes,
        topology_option=Network.topology_option.NONE,
    )
    print("Sleeping for seconds...", end='')
    time.sleep(3)
    print("done")
    # Provision nodes
    # FIXME: Configure network conditions between nodes according to specified test case

    actual_topo = n.get_actual_topology()
    print("actual_topo =", actual_topo)
    print(get_degree_map(actual_topo))
    trim_connections(n, 0.3)
    trimmed_topo = n.get_actual_topology()
    print("trimmed_topo =", trimmed_topo)
    print(get_degree_map(trimmed_topo))
    visualize(trimmed_topo, "network.png")

    print("Sleeping for seconds...", end='')
    time.sleep(3)
    print("done")

    # TODO: Configure actions and behavior between nodes according to specified test case
    validators = n.nodes[len(n.bootnodes):len(n.bootnodes)+num_validators]
    static_nodes = n.nodes[len(n.bootnodes)+num_validators:len(n.bootnodes)+num_validators+num_static_nodes]
    static_nodes_subscribe(static_nodes, num_shards, num_shards // 4)

    print("Sleeping for seconds...", end='')
    time.sleep(3)
    print("done")

    pass
    # TODO: [Added] Shutdown the test, according to the event occurred
    pass
    # TODO: Output performance data in CSV format
    #       Should aggregate first, and then output?
    pass
    # TODO: Aggregate data, parse, & present data
    pass
    # Reset environment
    pass


if __name__ == "__main__":
    l = logging.getLogger("simulation.Network")
    h = logging.StreamHandler()
    h.setLevel(logging.DEBUG)
    l.addHandler(h)
    # test_time_broadcasting_data_single_shard()
    # test_joining_through_bootnodes()
    # test_reproduce_bootstrapping_issue()
    test_plan()
