# EnergyOriginTrackingUsingHyperledgerFabric

Virtual environment
--------------------
Vagrant is used for creating a virtual environment.
Fabric will be setup in the Ubuntu 16.04 LTS VM

Steps to launch the VM
======================
Launch VM via::

    ~# vagrant up
Login to VM via::

    ~# vagrant ssh


Scripts needed to install pre-requisite software are under the network folder
-----------------------------------------------------------------------------
1. Change the directory::


    ~# cd ./network/setup
    ~# chmod u+x *.sh

2. Install the pre-requisites (docker, docker-compose & Go)::


    ~# ./install-prereqs.sh
    
3. Install Fabric::


    ~# sudo -E ./install-fabric.sh

4. Install Fabric CA Server::


    ~# ./install-caserver.sh
    
5. Install Go Tools::


    ~# ./install-gotools.sh
    ~# ./ govendor -h  #to validate

6. Install Hyperledger Explorer::


    ~# ./install-explorer.sh
    ~# ./validate-explorer.sh


Launch the environment
----------------------

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
