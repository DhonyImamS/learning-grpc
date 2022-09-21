package rpctest

import (
  "testing"
	"learn_grpc/common/model"
	"learn_grpc/test_integration/rpc"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGarageRPC(t *testing.T) {

  t.Run("it should return not found for not found GarageUserId", func(t *testing.T) {
    payloadFetch := model.GarageUserId{
      UserId: "NotFOUND",
    }
  
    // invoke list garage through grpc
    _, err := rpc.FetchListGarage(&payloadFetch)
    
    if assert.Error(t, err) {
      assert.Equal(t, status.Errorf(codes.NotFound, "Data Not Found"), err)
    }
  })

  t.Run("it should return empty body when succeed register new garage", func(t *testing.T) {
    garageNew := model.Garage{
      Id:   "q001",
      Name: "Test Bersama Mba Ayu",
      Coordinate: &model.GarageCoordinate{
      Latitude:  45.123123123,
        Longitude: 54.1231313123,
      },
    }
  
    payload := model.GarageAndUserId{
      UserId: "u00g1",
      Garage: &garageNew,
    }
  
    // invoke add garage through grpc
    resp, err := rpc.SendRegisterGarage(&payload)

    assert.Equal(t, nil, err)
    assert.Equal(t, "", resp.String())
  })

  t.Run("it should return new data garage on fetchList garage after succeed registration", func(t *testing.T) {
    garageNew := model.Garage{
      Id:   "q001",
      Name: "Test Bersama Mba Ayu",
      Coordinate: &model.GarageCoordinate{
      Latitude:  45.123123123,
        Longitude: 54.1231313123,
      },
    }
  
    payload := model.GarageAndUserId{
      UserId: "u00g1",
      Garage: &garageNew,
    }
  
    // invoke add garage through grpc
    resp, err := rpc.SendRegisterGarage(&payload)

    assert.Equal(t, nil, err)
    assert.Equal(t, "", resp.String())

    // assertion on fetchList
    payloadFetch := model.GarageUserId{
      UserId: payload.UserId,
    }
  
    // invoke list garage through grpc
    respFetch, err := rpc.FetchListGarage(&payloadFetch)
    assert.Equal(t, nil, err)

    for index := range respFetch.List {
      assert.Equal(t, garageNew.String(), respFetch.List[index].String())
    }
  })
}