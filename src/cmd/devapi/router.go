package devapi

import (
	"github.com/lumialvarez/go-grpc-profile-service/src/cmd/devapi/config"
	"google.golang.org/grpc"
)

func ConfigureServers(grpcServer *grpc.Server, config config.Config) {

	handlers := LoadDependencies(config)

	registerServers(grpcServer, handlers)
}

func registerServers(grpcServer *grpc.Server, handlers DependenciesContainer) {
	//pb.RegisterAuthServiceServer(grpcServer, handlers.AuthService)
}
