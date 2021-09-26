package init

import "github.com/fauzanmh/olp-user/docs"

func setupSwagger() {
	docs.SwaggerInfo.Title = "Online Learning Platform API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8100"
	docs.SwaggerInfo.BasePath = "/api"
}
