package user

import "gorm.io/gorm"

type User struct {
	*gorm.Model `json:"-"`

	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}
