# learn_grpc
This repository is my learning process to interact with simple gRPC in GO Language, using Go 1.16
Owner try to learn test this grpc service using testify using TDD style and Godog using BDD style

# resources
This repository using many source comes from 
```
1. https://dasarpemrogramangolang.novalagung.com/C-golang-protobuf-implementation.html
2. https://go-tour-id2.appspot.com/basics/1 ( read this minimum before concurrency )
3. https://github.com/stretchr/testify
4. https://github.com/cucumber/godog

```

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
│   ├── test_integration
│   │   ├── rpc
│   │   │   ├── garage.go
│   │   ├── rpctest
│   │   │   ├── garage_test.go
│   │   ├── rpctestbdd
│   │   │   ├── assertions
│   │   │   |   ├── custom_assertions.go
│   │   │   ├── features
│   │   │   |   ├── garage.features
│   │   │   ├── garage_step_def_test.go
│   │   │   ├── godogs.go
│   ├── go.mod
│   ├── go.sum
```

## Installation all dependency for this project from go.mod file
```
run => go mod download, from root project
```

## How to Run Service Dummy on your localhost
```
run service garage or service user, via terminal with these step:
1. navigate into folder services/services-garage or services/services-user
2. run => go run main.go

OR

run directly from root project using this command ( sample for run service garage ):
=> go run services/service-garage/main.go
```

## How to hit to localhost gRPC servers programmatically in GO
```
run main.go file under folder client using => go run main.go
```

## How to run integration test GRPC in GO TDD Style
```
1. run your service garage first
2. from main root project , run => go test ./test_integration_/rpctest/... 
```

## How to run integration test GRPC in GO using GODOG for BDD Style
```
1. run your service garage first
2. from main root project , run => go test -v test_integration/rpctestbdd/garage_step_def_test.go
```

