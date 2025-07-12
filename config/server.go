package config

import (
	"net/http"

	"github.com/assidik12/go-restfull-api/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    ":3000",
		Handler: authMiddleware,
	}
}
