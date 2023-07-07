package handler

import (
	authHandler "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/delivery/http/handler/auth"
)

type Handler struct {
	AuthHandler *authHandler.AuthHandler
}
