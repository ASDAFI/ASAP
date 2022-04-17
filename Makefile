postgres:
	cd ./db/postgres; bash setup.sh

mqtt:
	cd ./mqtt; bash setup.sh

add-device:
	docker exec mqtt mosquitto_passwd -b /mosquitto/config/pwfile $(serial) $(password)
