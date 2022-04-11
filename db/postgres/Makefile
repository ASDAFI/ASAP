
up:
	docker-compose -f postgres-docker-compose.yaml up

down:
	docker-compose -f postgres-docker-compose.yaml down

backup:
	mkdir backup
	bash backup.sh

backup-cronjob:
	bash setup_cronjob.sh


setup:
	bash setup.sh
