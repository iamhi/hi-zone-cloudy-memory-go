package dataservice

import (
	"github.com/iamhi/cloudy-memory-go/src/db/redisclient"

	"github.com/iamhi/cloudy-memory-go/src/domain"
)

func  GetData(userData domain.UserData, path string) string {
	return redisclient.GetValue(userData.Uuid, path)
}

func SetData(userData domain.UserData, path string, data string) {
	redisclient.SetValue(userData.Uuid, path, data);
}

func GetPaths(userData domain.UserData) []string {
	return redisclient.GetKeys(userData.Uuid)
}

func DeletePath(userData domain.UserData, path string) string {
	var deleted_value = redisclient.GetValue(userData.Uuid, path)
	
	redisclient.DeleteValue(userData.Uuid, path)

	return deleted_value
}
