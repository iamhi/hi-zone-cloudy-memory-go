package dataservice

import (
	"github.com/iamhi/cloudy-memory-go/src/db/redisclient"

	"github.com/iamhi/cloudy-memory-go/src/domain"
)

func createUserPath(userData domain.UserData, path string) string {
	return userData.Uuid + "/" + path
}

func  GetData(userData domain.UserData, path string) string {
	return redisclient.GetValue(createUserPath(userData, path))
}

func SetData(userData domain.UserData, path string, data string) {
	redisclient.SetValue(createUserPath(userData, path), data);
}
