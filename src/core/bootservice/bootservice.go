package bootservice

import (
	"github.com/iamhi/cloudy-memory-go/config/viperconfig"
	"github.com/iamhi/cloudy-memory-go/src/core/authenticationservice"
	"github.com/iamhi/cloudy-memory-go/src/db/redisclient"
)

func BootUp() {
	viperconfig.Setup()
	authenticationservice.StartAuthenticationService()
	redisclient.StartUp()
}
