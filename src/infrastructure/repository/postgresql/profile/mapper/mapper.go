package mapper

import (
	"encoding/json"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/repository/postgresql/profile/dao"
	"github.com/lumialvarez/go-grpc-profile-service/src/internal/profile"
	"gorm.io/datatypes"
	"log"
	"time"
)

type Mapper struct {
}

func (m Mapper) ToDomain(daoProfile *dao.Profile) (*profile.Profile, error) {
	serializedData := daoProfile.Data

	var profileData profile.Data
	err := json.Unmarshal(serializedData, &profileData)
	if err != nil {
		log.Print("Error to convert DAO to Domain (" + err.Error() + ")")
	}
	domainProfile := profile.NewProfile(daoProfile.Id, daoProfile.Language, profileData, time.Time(daoProfile.LastUpdate), daoProfile.Status)
	return domainProfile, nil
}

func (m Mapper) ToDAO(domainProfile *profile.Profile) (*dao.Profile, error) {

	serializedData, err := json.Marshal(domainProfile.Data())
	if err != nil {
		return nil, err
	}

	daoProfile := dao.Profile{
		Id:         domainProfile.Id(),
		Language:   domainProfile.Language(),
		Data:       datatypes.JSON(serializedData),
		LastUpdate: datatypes.Date(domainProfile.LastUpdate()),
		Status:     domainProfile.Status(),
	}
	return &daoProfile, nil
}
