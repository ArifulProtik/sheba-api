package router

import (
	"github.com/ArifulProtik/sheba-api/controller"
	"github.com/ArifulProtik/sheba-api/pkg/auth"
	"github.com/ArifulProtik/sheba-api/pkg/services"
	"github.com/labstack/echo/v4"
)

func InitRouter(r *echo.Group, s *services.Service, auth *auth.Token, key string) {
	handler := controller.New(s, auth)

	r.POST("/signup", handler.Auth.Signup)
	r.POST("/signin", handler.Auth.Signin)
	r.GET("/refresh", handler.Auth.Refresh)

	// services
	r.GET("/services", handler.Serv.GetServices)

	r.Use(handler.Auth.IsAuth)

	r.GET("/myservices", handler.Serv.MyService)
	r.POST("/service/create", handler.Serv.CreateService)
	r.POST("/location/create", handler.Serv.CreateLocation)
	r.POST("/category/create", handler.Serv.CreateCat)
	r.DELETE("/service/:id", handler.Serv.DeleteService)
	r.GET("/logout", handler.Auth.Logout)
}
