package authenticationv2

type DecodeResult struct {
	
	Uuid string `json:"uuid"`

	Username string `json:"username"`

	Roles []string `json:"roles"`
}

