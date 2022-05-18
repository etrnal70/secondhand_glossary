package rest

import (
	// "os"
	"secondhand_glossary/internal/config"
	// "secondhand_glossary/internal/middleware/jwt"
	"secondhand_glossary/internal/repository"
	"secondhand_glossary/internal/service"
	"secondhand_glossary/platform"
  _ "secondhand_glossary/docs"

	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	// "github.com/labstack/echo/v4/middleware"
)

func RegisterProtectedGroupAPI(e *echo.Echo, conf config.Config, platform *platform.Connection) {
	// categoryRepo := repository.NewCategoryRepoDriver(platform.Persistence)
	deviceRepo := repository.NewDeviceRepoDriver(platform.Persistence)
	// traitRepo := repository.NewTraitRepoDriver(platform.Persistence)
	userRepo := repository.NewUserRepoDriver(platform.Persistence)
	// categoryService := service.NewCategoryService(categoryRepo)
	deviceService := service.NewDeviceService(deviceRepo)
	// traitService := service.NewTraitService(traitRepo)
	userService := service.NewUserService(userRepo)

	// categoryController := CategoryController{
	// 	conf: conf,
	// 	s:    categoryService,
	// }
	deviceController := DeviceController{
		conf: conf,
		s:    deviceService,
	}
	// traitController := TraitController{
	// 	conf: conf,
	// 	s:    traitService,
	// }
	userController := UserController{
		conf: conf,
		s:    userService,
	}

	// TODO Auth, Logging Middleware
	protectedApi := e.Group("/protected")
	// protectedApi.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	Claims:      &jwt.CustomClaims{},
	// 	TokenLookup: "query:user",
	// 	// SigningKey:              []byte(os.Getenv("JWT_SECRET")),
	// 	ErrorHandlerWithContext: jwt.JWTError,
	// }))

	// User service
	protectedApi.GET("/user/profile", userController.GetProfileDetailsController)
	protectedApi.POST("/user/profile", userController.UpdateProfileController)

	// Device service
	protectedApi.POST("/device/:device_id/review", deviceController.AddDeviceReviewController)
	protectedApi.PUT("/device/:device_id/review/:review_id", deviceController.EditDeviceReviewController)
	protectedApi.DELETE("/device/:device_id/review/:review_id", deviceController.DeleteDeviceReviewController)

	// Category service

	// Trait service
}

func RegisterAdminGroupAPI(e *echo.Echo, conf config.Config, platform *platform.Connection) {
	categoryRepo := repository.NewCategoryRepoDriver(platform.Persistence)
	deviceRepo := repository.NewDeviceRepoDriver(platform.Persistence)
	traitRepo := repository.NewTraitRepoDriver(platform.Persistence)
	userRepo := repository.NewUserRepoDriver(platform.Persistence)
	categoryService := service.NewCategoryService(categoryRepo)
	deviceService := service.NewDeviceService(deviceRepo)
	traitService := service.NewTraitService(traitRepo)
	userService := service.NewUserService(userRepo)

	categoryController := CategoryController{
		conf: conf,
		s:    categoryService,
	}
	deviceController := DeviceController{
		conf: conf,
		s:    deviceService,
	}
	traitController := TraitController{
		conf: conf,
		s:    traitService,
	}
	userController := UserController{
		conf: conf,
		s:    userService,
	}

	// TODO Auth, Logging Middleware
	adminApi := e.Group("/admin")

	// User service
	adminApi.GET("/user", userController.GetUsersController)
	adminApi.DELETE("/user/delete/:user_id", userController.DeleteUserController)

	// Device Service
	adminApi.POST("/device", deviceController.AddDeviceController)
	adminApi.PUT("/device", deviceController.EditDeviceController)
	adminApi.DELETE("/device/:device_id", deviceController.DeleteDeviceController)
	adminApi.POST("/device/:device_id/link", deviceController.AddDeviceLinkController)
	adminApi.PUT("/device/:device_id/link/:link_id", deviceController.EditDeviceLinkController)
	adminApi.DELETE("/device/:device_id/link/:link_id", deviceController.DeleteDeviceLinkController)
	adminApi.POST("/device/:device_id/trait/:trait_id", deviceController.AddDeviceTraitController)
	adminApi.DELETE("/device/:device_id/trait/:trait_id", deviceController.DeleteDeviceTraitController)

	// Category Service
	adminApi.POST("/category", categoryController.AddCategoryController)
	adminApi.DELETE("/category/:category_id", categoryController.DeleteCategoryController)

	// Trait Service
	adminApi.POST("/trait", traitController.AddTraitController)
	adminApi.PUT("/trait/:trait_id", traitController.EditTraitController)
	adminApi.DELETE("/trait/:trait_id", traitController.DeleteTraitController)
}

func RegisterPublicGroupAPI(e *echo.Echo, conf config.Config, platform *platform.Connection) {
	categoryRepo := repository.NewCategoryRepoDriver(platform.Persistence)
	deviceRepo := repository.NewDeviceRepoDriver(platform.Persistence)
	traitRepo := repository.NewTraitRepoDriver(platform.Persistence)
	userRepo := repository.NewUserRepoDriver(platform.Persistence)
	categoryService := service.NewCategoryService(categoryRepo)
	deviceService := service.NewDeviceService(deviceRepo)
	traitService := service.NewTraitService(traitRepo)
	userService := service.NewUserService(userRepo)

	categoryController := CategoryController{
		conf: conf,
		s:    categoryService,
	}
	deviceController := DeviceController{
		conf: conf,
		s:    deviceService,
	}
	traitController := TraitController{
		conf: conf,
		s:    traitService,
	}
	userController := UserController{
		conf: conf,
		s:    userService,
	}

	// TODO Logging Middleware
	publicApi := e.Group("/public")

	// Swagger
	publicApi.GET("/swagger/*", echoSwagger.WrapHandler)

	// User service
	publicApi.POST("/user/login", userController.LoginController)
	publicApi.POST("/user/register", userController.RegisterController)

	// Device service
	publicApi.GET("/device/:device_id", deviceController.GetDeviceController)
	publicApi.GET("/device", deviceController.GetDevicesController)
	publicApi.GET("/device/:device_id/review", deviceController.GetDeviceReviewsController)
	publicApi.GET("/device/:device_id/review/:review_id", deviceController.GetDeviceReviewController)
	publicApi.GET("/device/:device_id/score", deviceController.GetDeviceScoreController)
	publicApi.GET("/device/:device_id/link", deviceController.GetDeviceLinksController)
	publicApi.GET("/device/:device_id/link/:link_id", deviceController.GetDeviceLinkController)

	// Category service
	publicApi.GET("/category", categoryController.GetCategoriesController)
	publicApi.GET("/category/:category_id", categoryController.GetCategoryDevicesController)

	// Trait service
	publicApi.GET("/trait", traitController.GetTraitsController)
	publicApi.GET("/trait/:trait_id", traitController.GetTraitDevicesController)
}

// func RegisterCrawlerGroupAPI(g *echo.Group, conf config.Config, platform *platform.Connection) {}
