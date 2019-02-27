#!/bin/bash
set -o errexit -o nounset -o pipefail

### TODO this needs better setup

tendermint node --home=$HOME/.credchain --proxy_app=tcp://localhost:11111 > tendermint.log &
credchain start -bind=tcp://localhost:11111 > credchain.log & 

echo "Waiting for blockchain to start up"
sleep 5
tallybox start -port=5005 -remote=http://localhost:26657 > tallybox.log &
