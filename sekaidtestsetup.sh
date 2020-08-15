#!/bin/bash

rm -rf $HOME/.sekaid/

cd $HOME

sekaid init --chain-id=testing testing
sekaid keys add validator
sekaid add-genesis-account $(sekaid keys show validator -a) 1000000000stake,1000000000validatortoken
sekaid gentx-claim validator
sekaid start
