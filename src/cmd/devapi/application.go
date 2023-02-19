package devapi

import (
	"fmt"
	"google.golang.org/grpc"
	"net"

	//"fmt"
	"github.com/lumialvarez/go-grpc-profile-service/src/cmd/devapi/config"
	"log"
	//"net"
)

func Start() {
	appConfig, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at appConfig", err)
	}

	lis, err := net.Listen("tcp", appConfig.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Profile Svc on", appConfig.Port)

	grpcServer := grpc.NewServer()

	ConfigureServers(grpcServer, appConfig)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
