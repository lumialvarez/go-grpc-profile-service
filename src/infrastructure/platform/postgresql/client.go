package postgresql

import (
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/repository/postgresql/profile/dao"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

func Init(url string) Client {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Print("Connected to " + url)

	db.AutoMigrate(dao.Profile{})

	return Client{db}
}
