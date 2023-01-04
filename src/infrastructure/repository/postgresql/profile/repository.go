package repositoryProfile

import (
	"github.com/lumialvarez/go-grpc-profile-service/src/cmd/devapi/config"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/platform/postgresql"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/repository/postgresql/profile/dao"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/repository/postgresql/profile/mapper"
	"github.com/lumialvarez/go-grpc-profile-service/src/internal/profile"
)

type Repository struct {
	postgresql postgresql.Client
	mapper     mapper.Mapper
}

func Init(config config.Config) Repository {
	return Repository{postgresql: postgresql.Init(config.DBUrl), mapper: mapper.Mapper{}}
}

func (repository *Repository) GetById(id int64) (*profile.Profile, error) {
	var daoProfile dao.Profile
	result := repository.postgresql.DB.Where(&dao.Profile{Id: id}).First(&daoProfile)
	if result.Error != nil {
		return nil, result.Error
	}
	domainProfile, err := repository.mapper.ToDomain(&daoProfile)
	if err != nil {
		return nil, err
	}
	return domainProfile, nil
}

func (repository *Repository) GetByLanguage(language string) (*profile.Profile, error) {
	var daoProfile dao.Profile
	result := repository.postgresql.DB.Where(&dao.Profile{Language: language}).First(&daoProfile)
	if result.Error != nil {
		return nil, result.Error
	}
	domainProfile, err := repository.mapper.ToDomain(&daoProfile)
	if err != nil {
		return nil, err
	}
	return domainProfile, nil
}

func (repository *Repository) Save(domainProfile *profile.Profile) (*profile.Profile, error) {
	daoProfile, err := repository.mapper.ToDAO(domainProfile)
	if err != nil {
		return nil, err
	}
	result := repository.postgresql.DB.Create(&daoProfile)
	if result.Error != nil {
		return nil, result.Error
	}
	savedDomainProfile, _ := repository.mapper.ToDomain(daoProfile)

	return savedDomainProfile, nil
}

func (repository *Repository) GetAll() ([]*profile.Profile, error) {
	var daoProfiles []dao.Profile
	result := repository.postgresql.DB.Find(&daoProfiles)
	if result.Error != nil {
		return nil, result.Error
	}
	var domainProfiles []*profile.Profile
	for _, dao := range daoProfiles {
		domainProfile, err := repository.mapper.ToDomain(&dao)
		if err != nil {
			return nil, result.Error
		}
		domainProfiles = append(domainProfiles, domainProfile)
	}
	return domainProfiles, nil
}