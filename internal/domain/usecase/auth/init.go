package auth

import (
	eu "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/entity/user"
	ru "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/repository/user"
)

type AuthUsecase interface {
	Register(user *eu.User) error
	Login(email, password string) (*eu.User, error)
}

type authUsecase struct {
	userRepository ru.UserRepository
}

func New(userRepository ru.UserRepository) *authUsecase {
	return &authUsecase{
		userRepository,
	}
}
