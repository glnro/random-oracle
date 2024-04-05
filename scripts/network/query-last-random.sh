#!/bin/bash

source ./vars.sh  &> /dev/null

$BINARY q oracle random --node "tcp://127.0.0.1:29170" -o json