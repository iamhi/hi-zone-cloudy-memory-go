package authenticationv2

type LoginResult struct {

	AccessToken string `json:"accessToken"`

	RefreshToken string `json:"refreshToken"`
}


