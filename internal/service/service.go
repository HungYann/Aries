package service

import (
	"net/http"
	"server/internal/config"
)

func NewServer(applicationConfig config.Application) *http.Server {
	return &http.Server{
		Addr:    applicationConfig.Address,
		Handler: &HttpProxyHandler{},
	}
}
