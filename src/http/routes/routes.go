package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamhi/cloudy-memory-go/src/core/dataservice"
	"github.com/iamhi/cloudy-memory-go/src/core/userservice"
)

type Routes struct {}

func getData(ctx *gin.Context) {
	access_token := getToken(ctx)
	path, query_param_exists := ctx.GetQuery(pathQueryParam)

	if !query_param_exists {
		ctx.String(http.StatusBadRequest, "Empty path")

		return
	}

	user_data, user_data_error := userservice.GetUserData(access_token)

	if user_data_error != nil {
		ctx.String(http.StatusBadRequest, user_data_error.Error())
		return
	}

	stored_user_data := dataservice.GetData(user_data, path)

  ctx.IndentedJSON(http.StatusOK, stored_user_data)
}

func postData(ctx *gin.Context) {
	access_token := getToken(ctx)
	path, query_param_exists := ctx.GetQuery(pathQueryParam)
	var post_data_request postDataRequest

	if !query_param_exists {
		ctx.String(http.StatusBadRequest, "Empty path")

		return
	}
	
	if err:= ctx.BindJSON(&post_data_request); err != nil {
		return
	}

	user_data, user_data_error := userservice.GetUserData(access_token)

	if user_data_error != nil {
		ctx.String(http.StatusBadRequest, user_data_error.Error())
		return
	}

	dataservice.SetData(user_data, path, post_data_request.Value)

	ctx.String(http.StatusOK, "OK")
}

func Setup() {
	fmt.Println("Setting routes")

  router := gin.Default()
	
	router.GET(endpoint_prefix, getData)
	router.POST(endpoint_prefix, postData)

	router.Run("localhost:8080")

	fmt.Println("Server started")
}

