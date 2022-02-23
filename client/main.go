package main

import (
	"context"
	"log"
	"google.golang.org/grpc"
	"learn_grpc/common/model"
	"learn_grpc/common/config"
	"github.com/golang/protobuf/ptypes/empty"
	"encoding/json"
  )

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
    conn, err := grpc.Dial(port, grpc.WithInsecure())
    if err != nil {
        log.Fatal("could not connect to", port, err)
    }

    return model.NewUsersClient(conn)
}


func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
 
	return model.NewGaragesClient(conn)
 }


 func main() {
	user1 := model.User{
		  Id:       "n001",
		  Name:     "Noval Agung",
		  Password: "kw8d hl12/3m,a",
		  Gender:   model.UserGender(model.UserGender_value["MALE"]),
	  }
  
	garage1 := model.Garage{
		  Id:   "q001",
		  Name: "Test Bersama Mba Ayu",
		  Coordinate: &model.GarageCoordinate{
			  Latitude:  45.123123123,
			  Longitude: 54.1231313123,
		  },
	  }

	garageUserId := model.GarageUserId{
		UserId: "u00g5",
	}

	garageAndUserId := model.GarageAndUserId{
		UserId: "u00g1",
		Garage: &garage1,
	}
	

	user := serviceUser()
	garage := serviceGarage()
  
	// invoke method for register rpc with data user1
	user.Register(context.Background(), &user1)

	// show all registered users
	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

	// invoke add garage through grpc
	garage.Add(context.Background(), &garageAndUserId)

	// invoke list garage through grpc
	res2, err := garage.List(context.Background(), &garageUserId)
	if err != nil {
		log.Fatal(err.Error())
	}
	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))
  }