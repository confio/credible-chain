#!/bin/bash
set -o errexit -o nounset -o pipefail

# set up new keys and initialize all configs
KEY=$(tallybox keys)
tendermint init --home=$HOME/.credchain
credchain init $KEY
