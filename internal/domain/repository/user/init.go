package user

import (
	eu "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/entity/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *eu.User) error
	GetUserByEmail(email string) (*eu.User, error)
}

type userRepo struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) UserRepository {
	return &userRepo{
		DB,
	}
}
