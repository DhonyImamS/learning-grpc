package rpctestbdd

import (
	"context"
	"errors"
	"testing"
	"github.com/cucumber/godog"
	"learn_grpc/common/model"
	"learn_grpc/test_integration/rpc"
	"github.com/stretchr/testify/assert"
	customMatcher "learn_grpc/test_integration/rpctestbdd/assertions"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// godogsCtxKey is the key used to store the available godogs in the context.Context.
type godogsCtxKey struct{}

func setDataGarage(ctx context.Context) (context.Context, error) {
	garageNew := model.Garage{
		Id:   "q001",
		Name: "Test Bersama Mba Ayu",
		Coordinate: &model.GarageCoordinate{
		Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}
	
	payloadRegisterGarage := model.GarageAndUserId{
		UserId: "u00g1",
		Garage: &garageNew,
	}
	return context.WithValue(ctx, godogsCtxKey{}, payloadRegisterGarage), nil
}

func invokeRPCGarageRegister(ctx context.Context) (context.Context, error) {
	payloadRegisterGarage, ok := ctx.Value(godogsCtxKey{}).(model.GarageAndUserId)

	if !ok {
		return ctx, errors.New("there is no data new garage to be registered")
	}

	respSendRegisterData, err := rpc.SendRegisterGarage(&payloadRegisterGarage)

	if err != nil {
		return ctx, errors.New("failed to invoke RPC Register Garage")
	}

	return context.WithValue(ctx, godogsCtxKey{}, respSendRegisterData), nil
}

func assertRegisterDataExistInListGarage(ctx context.Context) error {
	respSendRegisterData,_ := ctx.Value(godogsCtxKey{}).(*emptypb.Empty)
	customMatcher.GodogExpectedAndActual(assert.Equal, "", respSendRegisterData.String())

	// assertion on fetchList
	garageNew := model.Garage{
		Id:   "q001",
		Name: "Test Bersama Mba Ayu",
		Coordinate: &model.GarageCoordinate{
		Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

    payloadFetch := model.GarageUserId{
		UserId: "u00g1",
	}
	
	// invoke list garage through grpc
	respFetch, err := rpc.FetchListGarage(&payloadFetch)
	customMatcher.GodogExpectedAndActual(assert.Equal, nil, err)

	for index := range respFetch.List {
		customMatcher.GodogExpectedAndActual(assert.Equal, garageNew.String(), respFetch.List[index].String())
	}
    
	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
        ctx.Step(`^Client set data for new Garage$`, setDataGarage)
        ctx.Step(`^Client registered new Garage$`, invokeRPCGarageRegister)
        ctx.Step(`^Client will be able to retrieve new Garage data$`, assertRegisterDataExistInListGarage)
}