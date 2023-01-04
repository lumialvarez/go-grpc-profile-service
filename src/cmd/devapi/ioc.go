package devapi

import (
	"github.com/lumialvarez/go-grpc-profile-service/src/cmd/devapi/config"
	errorGrpc "github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/error"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/profile"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/profile/pb"
	repositoryProfile "github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/repository/postgresql/profile"
	"github.com/lumialvarez/go-grpc-profile-service/src/internal/profile/usecase/list"
	"github.com/lumialvarez/go-grpc-profile-service/src/internal/profile/usecase/save"
)

type DependenciesContainer struct {
	ProfileService pb.ProfileServiceServer
}

func LoadDependencies(config config.Config) DependenciesContainer {
	profileRepository := repositoryProfile.Init(config)

	userCaseSave := save.NewUseCaseSaveProfile(&profileRepository)
	userCaseList := list.NewUseCaseListProfile(&profileRepository)
	apiResponseProvider := errorGrpc.NewAPIResponseProvider()

	s := profile.NewHandler(userCaseList, userCaseSave, apiResponseProvider)

	return DependenciesContainer{
		ProfileService: &s,
	}
}
