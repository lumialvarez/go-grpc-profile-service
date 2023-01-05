package list

import (
	"context"
	domainError "github.com/lumialvarez/go-grpc-profile-service/src/internal/error"
	"github.com/lumialvarez/go-grpc-profile-service/src/internal/profile"
)

type Repository interface {
	Save(user *profile.Profile) (*profile.Profile, error)
	GetById(id int64) (*profile.Profile, error)
	GetByLanguage(language string) (*profile.Profile, error)
	GetAll() ([]*profile.Profile, error)
}

type UseCaseListProfile struct {
	repository Repository
}

func NewUseCaseListProfile(repository Repository) UseCaseListProfile {
	return UseCaseListProfile{repository: repository}
}

func (uc UseCaseListProfile) Execute(ctx context.Context, id int64, language string) ([]*profile.Profile, error) {
	if id > 0 {
		var domainProfiles []*profile.Profile
		domainProfile, err := uc.repository.GetById(id)
		if err != nil {
			return nil, err
		}
		if len(language) > 0 && domainProfile.Language() != language {
			return nil, domainError.NewNotFound("User ID and Language mismatch")
		}
		domainProfiles = append(domainProfiles, domainProfile)
		return domainProfiles, nil
	}

	if len(language) > 0 {
		var domainProfiles []*profile.Profile
		domainProfile, err := uc.repository.GetByLanguage(language)
		if err != nil {
			return nil, err
		}
		domainProfiles = append(domainProfiles, domainProfile)
		return domainProfiles, nil
	}

	domainUsers, err := uc.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return domainUsers, nil

}
