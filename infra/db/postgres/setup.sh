mkdir backup
bash setup_cronjob.sh
docker-compose -f docker-compose.yml up -d
