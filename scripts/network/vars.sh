#!/bin/bash

#node0, validator
export BINARY=./build/appd
export CHAIN_ID=test
export KEYRING=test
# User balance of stake tokens
export USER_COINS="100000000000stake"
# Amount of stake tokens staked
export STAKE="100000000stake"
# Node IP address
export NODE_IP="127.0.0.1"

export ALICE="alice"
export BARBARA="barbara"
export CINDY="cindy"

# Home directory
export HOME_DIR=$HOME

export USERS=($ALICE $BARBARA $CINDY)

# Validator moniker
export MONIKERS=("beacon" "val1" "val2")
export BEACON_MONIKER="beacon"

export NODES_ROOT_DIR=${HOME_DIR}/cosmos/nodes

# Base port. Ports assigned after these ports sequentially by nodes.
export RPC_LADDR_BASEPORT=29170
export P2P_LADDR_BASEPORT=29180
export GRPC_LADDR_BASEPORT=29190
export NODE_ADDRESS_BASEPORT=29200
export PPROF_LADDR_BASEPORT=29210
export CLIENT_BASEPORT=29220

export VALIDATOR_2=val2
export ALICE=alice
export BARBARA=barb
export CINDY=cindy
export DON=don
