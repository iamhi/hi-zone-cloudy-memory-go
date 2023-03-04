package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iamhi/cloudy-memory-go/src/core/dataservice"
	"github.com/iamhi/cloudy-memory-go/src/core/userservice"
)

type Routes struct {}

func getData(ctx *gin.Context) {
	access_token := getToken(ctx)
	path, query_param_exists := ctx.GetQuery(pathQueryParam)

	if !query_param_exists || path == "" {
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

func getPaths(ctx *gin.Context) {
	access_token := getToken(ctx)

	user_data, user_data_error := userservice.GetUserData(access_token)

	if user_data_error != nil {
		ctx.String(http.StatusBadRequest, user_data_error.Error())
		return
	}

	user_paths := dataservice.GetPaths(user_data)

  ctx.IndentedJSON(http.StatusOK, user_paths)
}

func deleteData(ctx *gin.Context) {
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

	stored_user_data := dataservice.DeletePath(user_data, path)

  ctx.IndentedJSON(http.StatusOK, stored_user_data)
}


func Setup() {
	fmt.Println("Setting routes")

  router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders: []string{ "Content-Type", "Accept", "User-Agent" },
		AllowMethods: []string{ "POST", "GET", "DELETE" },
		AllowOrigins: []string{ "http://localhost:3000", "https://api.ibeenhi.com", "https://hi-zone.ibeenhi.com" },
		AllowWildcard: true,
	}))
	
	router.GET(endpoint_prefix, getData)
	router.POST(endpoint_prefix, postData)
	router.DELETE(endpoint_prefix, deleteData)
	router.GET(endpoint_prefix + "/paths", getPaths)

	router.Run("localhost:8080")

	fmt.Println("Server started")
}

