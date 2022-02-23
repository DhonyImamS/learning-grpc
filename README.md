# learn_grpc
This repository is my learning process to interact with simple gRPC in go Language

## Folder Structure
For your convenient please put the file based on the folder structure.

```
├── LEARN_GRPC
│   ├── client
│   │   ├── main.go
│   ├── common
│   │   ├── config
│   │       ├── config.go
│   │   └── model
│   │       ├── empty.proto
│   │       ├── garage.proto
│   │       ├── user.proto
│   ├── services
│   │   ├── service-garage
│   │   │   ├── main.go
│   │   ├── service-user
│   │   │   ├── main.go
│   ├── go.mod
│   ├── go.sum
```


## How to Run Service Dummy on your localhost
```
run service garage or service user, via terminal using => go run main.go
```

## How to hit to localhost gRPC servers programmatically in GO
```
run main.go file under folder client using => go run main.go
```

