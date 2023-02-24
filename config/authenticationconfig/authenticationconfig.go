package authenticationconfig

import "github.com/spf13/viper"

type AuthenticationConfig struct {

	url string

	serviceUsername string

	servicePassword string
}

var authentication_config = AuthenticationConfig{}

func LoadProperties() {
	authentication_config.serviceUsername = viper.GetString("authentication.v2.username")
	authentication_config.servicePassword = viper.GetString("authentication.v2.password")
	authentication_config.url = viper.GetString("authentication.v2.url")
}

func GetServiceUsername() string {
	return authentication_config.serviceUsername	
}

func GetServicePassword() string {
	return authentication_config.servicePassword
}

func GetUrl() string {
	return authentication_config.url
}
