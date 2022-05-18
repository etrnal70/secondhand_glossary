package main

import (
	"secondhand_glossary/internal/config"
	"secondhand_glossary/internal/handler/rest"
	"secondhand_glossary/internal/middleware/logger"
	"secondhand_glossary/platform"

	"github.com/labstack/echo/v4"
)

// @title           Secondhand Glossary API
// @version         1.0
// @description     API for Secondhand Glossary service
// @termsOfService  http://swagger.io/terms/

// @contact.name   Mochammad Hanif R
// @contact.url    http://www.github.com/etrnal70
// @contact.email  hanifrmn@pm.me

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /
func main() {
	conf := config.InitializeConfig()

	e := echo.New()
  e.Use(logger.LoggerMiddleware)
  e.HTTPErrorHandler = logger.CustomHTTPErrorHandler
	// g := e.Group("/")

	platformConn := platform.InitPlatform(conf)
	defer platformConn.Close()

	rest.RegisterProtectedGroupAPI(e, conf, platformConn)
	rest.RegisterAdminGroupAPI(e, conf, platformConn)
	rest.RegisterPublicGroupAPI(e, conf, platformConn)

	e.Logger.Fatal(e.Start(":8081"))
}
