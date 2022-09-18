package main

import (
	"context"
	"log"
	"net"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learn_grpc/common/model"
	"learn_grpc/common/config"
)

var localStorage *model.GarageListByUser

func init() {
    localStorage = new(model.GarageListByUser)
    localStorage.List = make(map[string]*model.GarageList)
}

type GaragesServer struct{}

func (GaragesServer) Add(ctx context.Context, param *model.GarageAndUserId) (*empty.Empty, error) {
	userId := param.UserId
	garage := param.Garage

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(model.GarageList)
		localStorage.List[userId].List = make([]*model.Garage, 0)
	}
	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)

	log.Println("Adding garage", garage.String(), "for user", userId)

	return new(empty.Empty), nil
}

 func (GaragesServer) List(ctx context.Context, param *model.GarageUserId) (*model.GarageList, error) {
	userId := param.UserId

	fetchData, ok := localStorage.List[userId]

	if ok != true {
		return nil, status.Errorf(codes.NotFound, "Data Not Found")
	}

	return fetchData, nil
}

 func main() {
    srv := grpc.NewServer()
    var garageSrv GaragesServer
    model.RegisterGaragesServer(srv, garageSrv)

    log.Println("Starting RPC server at", config.SERVICE_GARAGE_PORT)

    l, err := net.Listen("tcp", config.SERVICE_GARAGE_PORT)
    if err != nil {
        log.Fatalf("could not listen to %s: %v", config.SERVICE_GARAGE_PORT, err)
    }

    log.Fatal(srv.Serve(l))
}
