package auth

import (
	eu "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/entity/user"
	hlp "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/helper"
)

func (authUsecase *authUsecase) Register(user *eu.User) error {
	hashedPassword, err := hlp.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = authUsecase.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (authUsecase *authUsecase) Login(email, password string) (*eu.User, error) {
	user, err := authUsecase.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = hlp.VerifyPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
