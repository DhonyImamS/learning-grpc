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
type godogsCtxKey struct {}

type responseAndData struct {
	response *emptypb.Empty
	data garageTestData
}

type garageTestData struct {
	GarageData model.Garage
	UserID string
}

func SetRegisterPayload(data garageTestData) model.GarageAndUserId {
	return model.GarageAndUserId{
		UserId: data.UserID,
		Garage: &data.GarageData,
	}
}

func setDataGarage(ctx context.Context) (context.Context, error) {
	dataTest := garageTestData{}

	dataTest.GarageData = model.Garage{
		Id:   "q001",
		Name: "Test Bersama Mba Ayu",
		Coordinate: &model.GarageCoordinate{
		Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

	dataTest.UserID = "u00g1"

	return context.WithValue(ctx, godogsCtxKey{}, dataTest), nil
}

func invokeRPCGarageRegister(ctx context.Context) (context.Context, error) {
	dataTest, ok := ctx.Value(godogsCtxKey{}).(garageTestData)

	if !ok {
		return ctx, errors.New("there is no data new garage to be registered")
	}

	payloadRegisterGarage := model.GarageAndUserId{
		UserId: dataTest.UserID,
		Garage: &dataTest.GarageData,
	}

	respSendRegisterData, err := rpc.SendRegisterGarage(&payloadRegisterGarage)

	if err != nil {
		return ctx, errors.New("failed to invoke RPC Register Garage")
	}
	
	assertContext := responseAndData{respSendRegisterData, dataTest}
	return context.WithValue(ctx, godogsCtxKey{}, assertContext), nil
}

func assertRegisterDataExistInListGarage(ctx context.Context) error {
	assertContext, _ := ctx.Value(godogsCtxKey{}).(responseAndData)
	customMatcher.GodogExpectedAndActual(assert.Equal, "", assertContext.response.String())

	// assertion on fetchList
	payloadFetch := model.GarageUserId{
		UserId: assertContext.data.UserID,
	}
	
	// invoke list garage through grpc
	respFetch, err := rpc.FetchListGarage(&payloadFetch)
	customMatcher.GodogExpectedAndActual(assert.Equal, nil, err)

	for index := range respFetch.List {
		customMatcher.GodogExpectedAndActual(assert.Equal, assertContext.data.GarageData.String(), respFetch.List[index].String())
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