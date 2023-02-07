package persistence

import (
	"github.com/hermanowiczpiotr/ola/user/domain/entity"
	"github.com/hermanowiczpiotr/ola/user/domain/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	DSN = os.Getenv("dsn")
)

type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func NewRepositories() *Repositories {
	db, error := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		panic("failed to connect persistence")
	}

	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}
}

func (r *Repositories) AutoMigrate() error {
	return r.db.AutoMigrate(&entity.User{})
}
