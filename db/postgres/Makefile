
up:
	docker-compose -f postgres-docker-compose.yaml up

down:
	docker-compose -f postgres-docker-compose.yaml down

backup:
	bash backup.sh

backup-cronjob:
	rm -rf backup.sh
	bash setup_cronjob.sh


setup:
	bash setup.sh