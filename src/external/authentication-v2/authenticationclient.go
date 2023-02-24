package authenticationv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/iamhi/cloudy-memory-go/config/authenticationconfig"
)

var application_json_media_type = "application/json"

func MakeLoginRequest(login_request LoginRequest) LoginResult {
	login_request_json, marshal_error := json.Marshal(login_request)

	if marshal_error != nil {
		log.Fatal(marshal_error)
	}

	login_response, login_error := http.Post(
		authenticationconfig.GetUrl() + "/user/login",
		application_json_media_type,
		bytes.NewBuffer(login_request_json)) 

	if login_error != nil {
		log.Fatal(login_error)
	}

	var login_result LoginResult

	decoding_error := json.NewDecoder(login_response.Body).Decode(&login_result)

	if decoding_error != nil {
		log.Fatal(decoding_error);
	}

	return login_result
}

func StartUp() {
	fmt.Println("Booting authentication client")

	login_result := MakeLoginRequest(LoginRequest{
		authenticationconfig.GetServiceUsername(),
		authenticationconfig.GetServicePassword(),
	})

	fmt.Println(login_result)
}

func MakeDecodeRequest (decode_request DecodeRequest, access_token string) (DecodeResult, error) {
	decode_request_json, marshal_error := json.Marshal(decode_request)

	if marshal_error != nil {
		log.Fatal(marshal_error)
	}

	decode_response, decode_error := http.Post(
			authenticationconfig.GetUrl() + "/token/decode?accessToken=" + access_token,
			application_json_media_type,
			bytes.NewBuffer(decode_request_json))
	
	if decode_error != nil {
		log.Fatal(decode_error)
	}

	if decode_response.StatusCode == http.StatusBadRequest {
		return DecodeResult{}, fmt.Errorf("Decoding invalid token")
	}

	var decode_result DecodeResult

	decoding_error := json.NewDecoder(decode_response.Body).Decode(&decode_result)

	if decoding_error != nil {
		log.Fatal(decoding_error)
	}

	return decode_result, nil
}

