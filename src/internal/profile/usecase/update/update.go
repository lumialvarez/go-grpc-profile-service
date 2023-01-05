package update

import (
	"context"
	domainError "github.com/lumialvarez/go-grpc-profile-service/src/internal/error"
	"github.com/lumialvarez/go-grpc-profile-service/src/internal/profile"
)

type Repository interface {
	Update(user *profile.Profile) (*profile.Profile, error)
	GetById(id int64) (*profile.Profile, error)
	GetByLanguage(language string) (*profile.Profile, error)
}

type UseCaseSaveProfile struct {
	repository Repository
}

func NewUseCaseUpdateProfile(repository Repository) UseCaseSaveProfile {
	return UseCaseSaveProfile{repository: repository}
}

func (uc UseCaseSaveProfile) Execute(ctx context.Context, domainProfile *profile.Profile) (*profile.Profile, error) {
	dbProfile, err := uc.repository.GetById(domainProfile.Id())
	if err != nil {
		return nil, domainError.NewNotFound("Profile no exists")
	}
	if domainProfile.Language() != dbProfile.Language() {
		return nil, domainError.NewNotFound("Invalid Language")
	}

	createdProfile, err := uc.repository.Update(domainProfile)
	if err != nil {
		return nil, err
	}

	return createdProfile, nil
}
