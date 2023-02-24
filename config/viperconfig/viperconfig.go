package viperconfig

import (
	"fmt"
	"log"

	"github.com/iamhi/cloudy-memory-go/config/authenticationconfig"
	"github.com/iamhi/cloudy-memory-go/config/redisconfig"
	"github.com/spf13/viper"
)

func Setup() {
	viper.SetConfigName("application")
	viper.SetConfigType("properties")
	viper.AddConfigPath("./config")
	
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found")
		} else {
			log.Fatal("Config file bad format")
		}
	}

	authenticationconfig.LoadProperties()
	redisconfig.LoadProperties()

	fmt.Println("Config reading successful")
}
