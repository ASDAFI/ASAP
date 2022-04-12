#!/bin/bash

mkdir cronjobs

content="
#!/bin/bash
docker exec -t postgres pg_dump -h 127.0.0.1 -p 5432 -U asap -w postgres -F p > $(pwd)/backup/postgres.sql
docker exec -t postgres pg_dump -h 127.0.0.1 -p 5432 -U asap -w asap -F p > $(pwd)/backup/asap.sql
docker exec -t postgres pg_dump -h 127.0.0.1 -p 5432 -U asap -w farmdb -F p > $(pwd)/backup/farmdb.sql
(date +%s) > $(pwd)/backup/last_recovery.timestamp
date > $(pwd)/backup/last_recovery.date
"


echo "$content" > backup.sh
#crontab -l > mycron
echo "0 */2 * * * bash $(pwd)/backup.sh" > cronjobs/backup.cron
crontab cronjobs/backup.cron

