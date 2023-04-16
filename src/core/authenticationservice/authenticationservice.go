package authenticationservice

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/iamhi/cloudy-memory-go/config/authenticationconfig"
	"github.com/iamhi/cloudy-memory-go/src/domain"
	authenticationv2 "github.com/iamhi/cloudy-memory-go/src/external/authentication-v2"
)

var access_token string

var refresh_token string

func getNewTokens() {
	fmt.Println("Getting new tokens")
	login_result, err := authenticationv2.MakeLoginRequest(authenticationv2.LoginRequest{
		Username: authenticationconfig.GetServiceUsername(),
		Password: authenticationconfig.GetServicePassword(),
	})

	if err != nil {
		fmt.Println("Unable to login", err)
		return;
	}

	access_token = login_result.AccessToken
	refresh_token = login_result.RefreshToken

	fmt.Println("New tokens set")
	fmt.Println(login_result)
}

func StartAuthenticationService() {
	getNewTokens()

	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Every(10).Minutes().Do(getNewTokens)

	scheduler.StartAsync()

	fmt.Println("Authentication service started")
}

func DecodeToken(token string) (domain.UserData, error) {
	if token == "" {
		return domain.UserData{}, fmt.Errorf("Empty token")
	}

	fmt.Println("Attempting to decode tokens")

	decode_result, decode_error := authenticationv2.MakeDecodeRequest(authenticationv2.DecodeRequest{
		 Token: token,
	}, access_token)

	if decode_error != nil {
		return domain.UserData{}, decode_error
	}

	return domain.UserData{
		Uuid: decode_result.Uuid,
		Username: decode_result.Username,
	}, nil
}
