package devapi

import (
	"github.com/lumialvarez/go-grpc-profile-service/src/cmd/devapi/config"
	errorGrpc "github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/error"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/profile"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/profile/pb"
	repositoryProfile "github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/repository/postgresql/profile"
	"github.com/lumialvarez/go-grpc-profile-service/src/internal/profile/usecase/list"
	"github.com/lumialvarez/go-grpc-profile-service/src/internal/profile/usecase/save"
	"github.com/lumialvarez/go-grpc-profile-service/src/internal/profile/usecase/update"
)

type DependenciesContainer struct {
	ProfileService pb.ProfileServiceServer
}

func LoadDependencies(config config.Config) DependenciesContainer {
	profileRepository := repositoryProfile.Init(config)

	useCaseSave := save.NewUseCaseSaveProfile(&profileRepository)
	useCaseList := list.NewUseCaseListProfile(&profileRepository)
	useCaseUpdate := update.NewUseCaseUpdateProfile(&profileRepository)
	apiResponseProvider := errorGrpc.NewAPIResponseProvider()

	s := profile.NewHandler(useCaseList, useCaseSave, useCaseUpdate, apiResponseProvider)

	return DependenciesContainer{
		ProfileService: &s,
	}
}
