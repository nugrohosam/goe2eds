package usecases

import (
	"github.com/nugrohosam/services/infrastructure"
)

func AuthorizationValidation(token) (bool, error) {
	return infrastructure.ValidateToken(token)
} 

// GetDataAuth ...
func GetDataAuth(tokenString string) (map[string]interface{}, error) {
	return infrastructure.GetDataAuth(token)
}