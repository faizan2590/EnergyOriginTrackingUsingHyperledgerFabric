#PEER_MODE=net
#Command=dev-init.sh -e 
#Generated: Sat Apr 11 14:26:18 UTC 2020 
docker-compose  -f ./compose/docker-compose.base.yaml      -f ./compose/docker-compose.explorer.yaml    up -d --remove-orphans
