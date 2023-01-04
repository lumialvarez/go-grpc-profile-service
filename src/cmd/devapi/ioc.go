package devapi

import (
	"github.com/lumialvarez/go-grpc-profile-service/src/cmd/devapi/config"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/auth/pb"
)

type DependenciesContainer struct {
	AuthService pb.AuthServiceServer
}

func LoadDependencies(config config.Config) DependenciesContainer {
	//userRepository := repositoryUser.Init(config)

	//userCaseRegister := register.NewUseCaseRegisterUser(&userRepository)

	//s := auth.NewHandler(userCaseRegister, useCaseLogin, useCaseValidate, useCaseList, useCaseUpdate, apiResponseProvider)

	return DependenciesContainer{
		//AuthService: &s,
	}
}
