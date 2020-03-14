#!/bin/bash

# run one/many
for arg in "$@"
do
  if [[ $arg == 1 ]]; then
    echo "run server..."
    go build -o ./server github.com/zii/pet-sim/cmd
    ./server -listen=:80
  else
    echo unknown argument: "$arg"
  fi
done
