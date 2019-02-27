#!/bin/bash
set -o errexit -o nounset -o pipefail

# Tested on Digital Ocean image
# Ubuntu 18.04, $20/mo

# get go
sudo apt install -y golang
cat << 'EOF' >> ~/.bash_aliases
export GOPATH=${HOME}/go
export PATH=${PATH}:${GOPATH}/bin
EOF
source ~/.bash_aliases

# build credible chain code
(
    mkdir -p ${GOPATH}/src/github.com/confio
    cd ${GOPATH}/src/github.com/confio
    git clone https://github.com/confio/credible-chain.git
    cd credible-chain

    # compile
    make deps
    make install

    # verify
    credchain version
    tallybox version
)

# build tendermint code
(
    mkdir -p ${GOPATH}/src/github.com/tendermint
    cd ${GOPATH}/src/github.com/tendermint
    git clone https://github.com/tendermint/tendermint.git
    cd tendermint

    # use v0.29.1 as that is what our code tests against
    git checkout v0.29.1 

    # compile
    make get_vendor_deps
    make install

    # verify
    tendermint version
)


