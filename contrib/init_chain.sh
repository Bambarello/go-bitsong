#!/bin/bash

PASSWORD="12345678"

rm -rf ~/.bitsong*

# Initialize the genesis.json file that will help you to bootstrap the network
bitsongd init MyValidator --chain-id=bitsong-dev-network-3

bitsongcli config chain-id bitsong-dev-network-3
bitsongcli config output json
bitsongcli config indent true
bitsongcli config trust-node true

# Change default bond token genesis.json
sed -i 's/stake/ubtsg/g' ~/.bitsongd/config/genesis.json
#sed -i 's/"send_enabled": true/"send_enabled": false/g' ~/.bitsongd/config/genesis.json

# Change gov parameters (2 min)
sed -i 's/"max_deposit_period": "172800000000000"/"max_deposit_period": "240000000000"/g' ~/.bitsongd/config/genesis.json
sed -i 's/"voting_period": "172800000000000"/"voting_period": "240000000000"/g' ~/.bitsongd/config/genesis.json

# Create a key to hold your validator account
echo ${PASSWORD} | bitsongcli keys add validator
echo ${PASSWORD} | bitsongcli keys add faucet

# Add that key into the genesis.app_state.accounts array in the genesis file
# NOTE: this command lets you set the number of coins. Make sure this account has some coins
# with the genesis.app_state.staking.params.bond_denom denom, the default is staking
bitsongd add-genesis-account validator 150000000000ubtsg
bitsongd add-genesis-account faucet 116000000000000ubtsg

# Generate the transaction that creates your validator
echo ${PASSWORD} | bitsongd gentx --name validator --amount=100000000000ubtsg

# Add the generated bonding transaction to the genesis file
bitsongd collect-gentxs
bitsongd validate-genesis

# Now its safe to start `bitsongd`
bitsongd start
