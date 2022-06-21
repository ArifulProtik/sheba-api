package controller

import (
	"github.com/ArifulProtik/sheba-api/pkg/auth"
	"github.com/ArifulProtik/sheba-api/pkg/services"
)

type Controller struct {
	Auth  *AuthController
	Serv  *ServiceController
	Order *OrderController
}

func New(services *services.Service, AuthToken *auth.Token) *Controller {
	return &Controller{
		Auth: &AuthController{
			UserService: services.User,
			AuthService: services.Auth,
			AuthToken:   AuthToken,
		},
		Serv: &ServiceController{
			Service: services.Serv,
		},
		Order: &OrderController{
			Service: services.Order,
			Serv:    services.Serv,
		},
	}
}
