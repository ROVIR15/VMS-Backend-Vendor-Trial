package delivery

import (
	"context"

	interactor "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/interactor"
	_grpc "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/grpc"
	"google.golang.org/grpc"
)

type GrpcEndpointConfigurations struct {
	_grpc.GrpcServer
}

func NewGrpcEndpointConfigurations(grpc _grpc.GrpcServer) *GrpcEndpointConfigurations {
	return &GrpcEndpointConfigurations{GrpcServer: grpc}
}

func (endpoint *GrpcEndpointConfigurations) RegisterRoutes(
	ctx context.Context,
	usecase *interactor.Usecase,
) {
	endpoint.GrpcServiceBuilder().RegisterRoutes(func(server *grpc.Server) {

		// Register your grpc service here
		// example:
		// exampleService := service.NewExampleService(usecase.ExampleUsecase)
		// example.RegisterExampleServiceServer(server, exampleService)
		// brandHandler := controllers.NewBrandController(usecase.CreateBrandHandler, usecase.FindOneBrandHandler, server, endpoint.Logger())
		// brandHandler.RegisterPath()

		// hotelHandler := controllers.NewHotelController(usecase.CreateHotelHandler, server, endpoint.Logger())
		// hotelHandler.RegisterPath()

	})

}
