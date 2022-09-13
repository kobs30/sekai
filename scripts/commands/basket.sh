#!/bin/bash

# queries
sekaid query basket token-basket-by-id 1
sekaid query basket token-basket-by-denom b1/usd
sekaid query basket token-baskets "" false

# transactions
sekaid tx basket mint-basket-tokens 1 1000stake --from=validator --keyring-backend=test --home=$HOME/.sekaid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
sekaid tx basket burn-basket-tokens 1 1000b1/usd --from=validator --keyring-backend=test --home=$HOME/.sekaid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
sekaid tx basket swap-basket-tokens 1 100000stake ueth --from=validator --keyring-backend=test --home=$HOME/.sekaid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
sekaid tx basket basket-claim-rewards 1000b1/usd --from=validator --keyring-backend=test --home=$HOME/.sekaid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
sekaid tx basket disable-basket-deposits 1 --from=validator --keyring-backend=test --home=$HOME/.sekaid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
sekaid tx basket disable-basket-swaps 1 --from=validator --keyring-backend=test --home=$HOME/.sekaid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
sekaid tx basket disable-basket-withdraws 1 --from=validator --keyring-backend=test --home=$HOME/.sekaid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block

# proposals
sekaid tx basket proposal-create-basket --title="title" --description="description" --basket-suffix="usd" --basket-description="usd stable coins collection" --swap-fee="0.001" --slippage-fee-min="0.001" --tokens-cap="10000000" --limits-period="86400" --mints-min="10000" --mints-max="1000000" --mints-disabled=false --burns-min="10000" --burns-max="1000000" --burns-disabled=false --swaps-min="10000" --swaps-max="10000" --swaps-disabled=false --basket-tokens="stake#10#false#false#false,ukex#1#false#false#false,ueth#0.1#false,false,true" --from=validator --keyring-backend=test --home=$HOME/.sekaid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
sekaid tx basket proposal-edit-basket --title="title" --description="description" --basket-id="1" --basket-suffix="usd" --basket-description="usd stable coins collection" --swap-fee="0.001" --slippage-fee-min="0.001" --tokens-cap="10000000" --limits-period="86400" --mints-min="10000" --mints-max="1000000" --mints-disabled=false --burns-min="10000" --burns-max="1000000" --burns-disabled=false --swaps-min="10000" --swaps-max="10000" --swaps-disabled=false --basket-tokens="stake#10#false#false#false,ukex#1#false#false#false,ueth#0.1#false,false,true" --from=validator --keyring-backend=test --home=$HOME/.sekaid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
sekaid tx basket proposal-basket-withdraw-surplus 1,2 $(sekaid keys show -a validator --keyring-backend=test) --title="title" --description="description" --from=validator --keyring-backend=test --home=$HOME/.sekaid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block

