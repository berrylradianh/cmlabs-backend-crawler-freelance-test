package auth

import (
	ua "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/usecase/auth"
)

type AuthHandler struct {
	authUsecase ua.AuthUsecase
}

func New(authUsecase ua.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase,
	}
}
