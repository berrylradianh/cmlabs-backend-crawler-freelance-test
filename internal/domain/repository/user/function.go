package user

import (
	"errors"

	eu "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/entity/user"
)

func (userRepo *userRepo) CreateUser(user *eu.User) error {
	existingUser := &eu.User{}
	err := userRepo.DB.Where("email = ?", user.Email).First(existingUser).Error
	if err != nil {
		err = userRepo.DB.Create(&user).Error
		if err != nil {
			return err
		}
	} else {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Email already exists")
	}

	return nil
}

func (userRepo *userRepo) GetUserByEmail(email string) (*eu.User, error) {
	user := &eu.User{}
	err := userRepo.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, errors.New("Record Not Found")
	}

	return user, nil
}
