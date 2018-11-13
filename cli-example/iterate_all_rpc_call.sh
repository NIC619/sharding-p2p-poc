#!/bin/bash
# This script is expected to be executed in the root dir of the repo

EXE_NAME="./sharding-p2p-poc"
IP=127.0.0.1
PORT=10000
RPCPORT=13000
LOGLEVEL="DEBUG"

# spinup_node {port_increment} {other_params}
spinup_node() {
    port=$((PORT+$1))
    rpcport=$((RPCPORT+$1))
    p=$@
    params=${@:2}
    $EXE_NAME -port=$port -rpcport=$rpcport -loglevel=$LOGLEVEL $params &
}

cli_prompt() {
    p=$@
    params=${@:2}
    echo "$EXE_NAME -rpcport=$((RPCPORT+$1)) -client $params"
}

# add_peer {port_increment_host} {port_increment_peer} {target_pid}
add_peer() {
    target_pid=$3
    `cli_prompt $1` addpeer $IP $((PORT+$2)) $target_pid
}

# subscribe_shard {port_increment} {shard_id} {shard_id} ...
subscribe_shard() {
    p=$@
    params=${@:2}
    `cli_prompt $1` subshard $params
}

# unsubscribe_shard {port_increment} {shard_id} {shard_id} ...
unsubscribe_shard() {
    p=$@
    params=${@:2}
    `cli_prompt $1` unsubshard $params
}

# get_subscribe_shard {port_increment}
get_subscribe_shard() {
    p=$@
    `cli_prompt $1` getsubshard
}

# broadcast_collation {port_increment} {shard_id} {num_collation} {size} {period}
broadcast_collation() {
    p=$@
    params=${@:2}
    `cli_prompt $1` broadcastcollation $params
}

# stop_server {port_increment}
stop_server() {
    p=$@
    `cli_prompt $1` stop
}

# listpeer {port_increment}
list_peer() {
    p=$@
    `cli_prompt $1` listpeer
}


# listtopicpeer {port_increment} {topic0} {topic1} ...
list_topic_peer() {
    p=$@
    `cli_prompt $1` listtopicpeer
}

# remove_peer {port_increment} peerID
remove_peer() {
    p=$@
    params=${@:2}
    `cli_prompt $1` removepeer $params
}



go build

for i in `seq 0 1`;
do
    spinup_node $i
done

sleep 2

# peer 0 add peer 1
add_peer 0 1

# peer 0 subscribe shard
subscribe_shard 0 1 2 3 4 5

# peer 1 subscribe shard
subscribe_shard 1 2 3 4

# get peer 0's subscribed shard
get_subscribe_shard 0

# peer 0 broadcast collations
broadcast_collation 0 2 2 100 0

# peer 1 broadcast collations
broadcast_collation 1 1 1 100 0
# exit code should be 1
if [ "$?" != "1" ]
then
    exit 1
fi

# peer 0 unsubscribe shard
unsubscribe_shard 0 2 4

get_subscribe_shard 0

list_peer 0

list_topic_peer 0

remove_peer 0 QmexAnfpHrhMmAC5UNQVS8iBuUUgDrMbMY17Cck2gKrqeX

for i in `seq 0 1`;
do
    stop_server $i
done

sleep 1
