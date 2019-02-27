#!/bin/bash
set -o errexit -o nounset -o pipefail

systemctl stop tallybox
systemctl stop credchain
systemctl stop tendermint

systemctl start tendermint
systemctl start credchain
systemctl start tallybox