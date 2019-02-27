#!/bin/bash
set -o errexit -o nounset -o pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# install the service files
for file in ${SCRIPT_DIR}/services/*; do
    cp $file /etc/systemd/system
    base=$(basename $file)
    systemctl enable $base
done

# order matters....
systemctl start tendermint.service
systemctl start credchain.service
systemctl start tallybox.service
