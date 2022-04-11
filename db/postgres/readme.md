# POSTGRESQL

## Setup
setup postgres containers and backup.
```
bash setup.sh
```

## Down
delete and stop database containers.
```
make down
```

## Up
pull and run database containers.
```
make up
```

### INFO

POSTGRES
```
Container: postgres
  port: 5432
  host: localhost
  username: asap
  password: asapDB!
```

PGADMIN4
```
Container: pgadmin4
  port: 5050
  username: asap@carriot.ir
  password: asapDB!)
```
