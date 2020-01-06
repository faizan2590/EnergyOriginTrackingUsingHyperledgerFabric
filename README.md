# EnergyOriginTrackingUsingHyperledgerFabric
Launch the environment
======================
Dev  mode:     dev-init.sh   -d  -e

Try out energy_origin_tracking
================
#1  

<<Terminal #1>>  


- Setup the organization context to acme
. set-env.sh acme

- Setup the chaincode environment
set-chain-env.sh   -n  token   -p energy_origin_tracking    -c '{"Args": ["init"]}'  

#2

<<Terminal #2>>

- Start the chaincode 
. set-env.sh acme
cc-run.sh

#3

<<Terminal #1>>

- Install & Instantiate the chaincode
chain.sh    install 
chain.sh    instantiate                             <<Observe terminal#2>>

Checkout explorer - you should see 1 transaction against the chaicode 'energy_origin_tracking'

#4

<<Terminal #1>>

Invoke
======
Add EnergyOrigin
> set-chain-env.sh         -i   '{"Args":["addEnergyOriginCertificate", "0x123", "1234", "2345", "32", "1", "3", "300"]}'
> chain.sh  invoke

Query
=====
Check Agent's EnergyOrigin
> set-chain-env.sh         -q   '{"Args":["getAgentDetail", "0x123"]}'
> chain.sh query


Checkout explorer to see the transaction
