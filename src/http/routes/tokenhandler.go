package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAccessTokenCookieValue (cookies []*http.Cookie) string {
	for _, cookie := range cookies {
		if cookie.Name == accessTokenCookieName {
			return cookie.Value
		}
	}

	return "" 
}

func getToken(ctx *gin.Context) string  {	
	access_token_header := ctx.GetHeader(accessTokenHeader)
	
	if access_token_header != "" {
		return access_token_header
	}

	access_token_query_param, query_param_exists := ctx.GetQuery(accessTokenQueryParam)

	if query_param_exists {
		return access_token_query_param
	}

	access_token_cookie := getAccessTokenCookieValue(ctx.Request.Cookies())

	if access_token_cookie != "" {
		return access_token_cookie
	}

	return ""
}
