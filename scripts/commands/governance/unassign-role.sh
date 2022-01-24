#!/bin/bash

RoleValidator=2

sekaid tx customgov proposal account unassign-role $RoleValidator  --title="title" --description="description" --addr=$(sekaid keys show -a validator --keyring-backend=test --home=$HOME/.sekaid) --from=validator --keyring-backend=test --home=$HOME/.sekaid --chain-id=testing --fees=100ukex --yes

sekaid query customgov proposals
sekaid query customgov proposal 1

sekaid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.sekaid --chain-id=testing --fees=100ukex --yes 

# check roles
sekaid query customgov roles $(sekaid keys show -a validator --keyring-backend=test --home=$HOME/.sekaid)
