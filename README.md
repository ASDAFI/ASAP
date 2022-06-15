# ASAP

Simple IOT based Farm management system using CQRS pattern.


## Build

### Setup Protobuf
get ready to compile .proto files
```
make docker-build
```

### Build Protobuf
compile .proto files
```
make all-by-docker
```

### Build Farm
compile go files
```
make farm
```

## Command Handler


### Serve
```
./farm serve
```

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

