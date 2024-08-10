package api

import (
	"log"
	"net"

	"auth-service/internal/items/config"
	"auth-service/internal/items/service"

	pb "auth-service/genproto/auth"

	"google.golang.org/grpc"
)

type API struct {
	service *service.Service
}

func New(service *service.Service) *API {
	return &API{
		service: service,
	}
}

func (a *API) RUN(config *config.Config, service *service.Service) error {
	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()

	pb.RegisterAuthServiceServer(serverRegisterer, service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
