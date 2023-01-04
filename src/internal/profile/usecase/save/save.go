package save

import (
	"context"
	domainError "github.com/lumialvarez/go-grpc-profile-service/src/internal/error"
	"github.com/lumialvarez/go-grpc-profile-service/src/internal/profile"
)

type Repository interface {
	Save(user *profile.Profile) (*profile.Profile, error)
	GetByLanguage(language string) (*profile.Profile, error)
}

type UseCaseSaveProfile struct {
	repository Repository
}

func NewUseCaseSaveProfile(repository Repository) UseCaseSaveProfile {
	return UseCaseSaveProfile{repository: repository}
}

func (uc UseCaseSaveProfile) Execute(ctx context.Context, domainProfile *profile.Profile) (*profile.Profile, error) {

	_, err := uc.repository.GetByLanguage(domainProfile.Language())
	if err == nil {
		return nil, domainError.NewAlreadyExists("Language already exists")
	}

	createdProfile, err := uc.repository.Save(domainProfile)
	if err != nil {
		return nil, err
	}

	return createdProfile, nil
}
