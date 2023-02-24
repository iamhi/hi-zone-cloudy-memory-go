package userservice

import (
	"github.com/iamhi/cloudy-memory-go/src/core/authenticationservice"
	"github.com/iamhi/cloudy-memory-go/src/domain"
)

func GetUserData(token string) (domain.UserData, error) {
	return authenticationservice.DecodeToken(token)
}
