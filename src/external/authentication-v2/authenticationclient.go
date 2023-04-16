package authenticationv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iamhi/cloudy-memory-go/config/authenticationconfig"
)

const application_json_media_type = "application/json"

func MakeLoginRequest(login_request LoginRequest) (LoginResult, error) {
	login_request_json, marshal_error := json.Marshal(login_request)

	if marshal_error != nil {
		return LoginResult{}, fmt.Errorf("Error marshaling the request")
	}

	login_response, login_error := http.Post(
		authenticationconfig.GetUrl() + "/user/login",
		application_json_media_type,
		bytes.NewBuffer(login_request_json)) 

	if login_error != nil {
		fmt.Println(login_error)

		return LoginResult{}, fmt.Errorf("Error connectin to authentication service")
	}

	var login_result LoginResult

	decoding_error := json.NewDecoder(login_response.Body).Decode(&login_result)

	if decoding_error != nil {
		// TODO: Log the error message or the response body as string
		return LoginResult{}, fmt.Errorf("Error decoing the response from authentication service")
	}

	return login_result, nil
}

func StartUp() {
	fmt.Println("Booting authentication client")

	login_result, err := MakeLoginRequest(LoginRequest{
		authenticationconfig.GetServiceUsername(),
		authenticationconfig.GetServicePassword(),
	})

	if err != nil {
		// Should I crash here?
		fmt.Println("Unable to logn", err)
	}

	fmt.Println(login_result)
}

func MakeDecodeRequest (decode_request DecodeRequest, access_token string) (DecodeResult, error) {
	decode_request_json, marshal_error := json.Marshal(decode_request)

	if marshal_error != nil {
		return DecodeResult{}, fmt.Errorf("Unable to marshal the decode request")
	}

	decode_response, decode_error := http.Post(
			authenticationconfig.GetUrl() + "/token/decode?accessToken=" + access_token,
			application_json_media_type,
			bytes.NewBuffer(decode_request_json))
	
	if decode_error != nil {
		return DecodeResult{}, fmt.Errorf("Error while requesting decode")
	}

	if decode_response.StatusCode == http.StatusBadRequest {
		return DecodeResult{}, fmt.Errorf("Decoding invalid token")
	}

	var decode_result DecodeResult

	decoding_error := json.NewDecoder(decode_response.Body).Decode(&decode_result)

	if decoding_error != nil {
		return DecodeResult{}, fmt.Errorf("Error while gettin response from decode requesting")
	}

	return decode_result, nil
}

