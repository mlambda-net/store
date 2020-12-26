package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mlambda-net/store/pkg/infrastructure/api"
	"github.com/mlambda-net/store/pkg/ports/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.accessCode AccessCode
// @tokenUrl https://oauth.mitienda.co.cr/token
// @authorizationurl https://oauth.mitienda.co.cr/authorize
// @oauth2RedirectUrl http://localhost:8002/store/swagger
// @scope.admin Grants read and write access to administrative information
func main() {

	services := api.NewApi()

	docs.SwaggerInfo.Title = "store API"
	docs.SwaggerInfo.Description = "This is the api for the store service."
	docs.SwaggerInfo.Version = services.GetVersion()
	docs.SwaggerInfo.Host = services.GetHost()
	docs.SwaggerInfo.BasePath = services.Path()
	docs.SwaggerInfo.Schemes = []string{"http", "https"}


	go func() {
		r := gin.New()
    url := ginSwagger.URL("http://localhost:8002/store/swagger/doc.json")
    oauht := ginSwagger.OAuth2("http://localhost:8002/store/swagger/oauth2-redirect.html")
    deep := ginSwagger.DeepLinking(true)
    r.GET("/store/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url, oauht, deep))
    _ = r.Run( fmt.Sprintf(":%d", services.Docs()))
	}()
  services.Start()
}
