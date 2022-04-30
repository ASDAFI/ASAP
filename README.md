# ASAP

## Build Protobuf
compile .proto files
```
make docker-build
make all-by-docker
```

## Build Farm
compile go files
```
make farm
```

## Command Handler


### Migrate DB
```
./farm migrate
```

### Add Farm
```
./farm create farm {farm_name} {production}
```

### Add User
```
./farm create user {username} {password} {farm_id} {first_name} {last_name} {phone} {email}
```

