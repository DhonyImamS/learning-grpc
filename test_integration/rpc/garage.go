package rpc

import (
	"google.golang.org/grpc"
	"learn_grpc/common/model"
	"log"
	"context"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)


func FetchListGarage(payload *model.GarageUserId)  (*model.GarageList, error) {
	conn, err := grpc.Dial("127.0.0.1:7000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", ":7000", err)
	}

	return model.NewGaragesClient(conn).List(context.Background(), payload)
}

func SendRegisterGarage(payload *model.GarageAndUserId) (*emptypb.Empty, error) {
	conn, err := grpc.Dial("127.0.0.1:7000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", ":7000", err)
	}

	return model.NewGaragesClient(conn).Add(context.Background(), payload)
}


