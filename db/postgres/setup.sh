mkdir backup
bash setup_cronjob.sh
docker-compose -f postgres-docker-compose.yaml up
