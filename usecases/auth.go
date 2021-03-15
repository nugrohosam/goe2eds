package usecases

import (
	"github.com/nugrohosam/goe2eds/services/infrastructure"
)

func AuthorizationValidation(token string) (bool, error) {
	return infrastructure.ValidateToken(token)
} 

// GetDataAuth ...
func GetDataAuth(token string) (map[string]interface{}, error) {
	return infrastructure.GetDataAuth(token)
}