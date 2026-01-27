package config

import (
	"os"
)

// GetPort retorna a porta do ambiente ou a padrão 8080
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}