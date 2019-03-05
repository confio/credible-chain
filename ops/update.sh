#!/bin/bash

SRC=${GOPATH:-$HOME/go}/src

cd $SRC/github.com/confio/credible-chain
git pull
make deps && make install

echo
echo "Validating versions:"
tallybox version
credchain version