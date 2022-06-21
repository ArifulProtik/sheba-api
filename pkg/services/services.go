package services

import "github.com/ArifulProtik/sheba-api/ent"

type Service struct {
	Auth *AuthService
	User *UserService
	Serv *ServiceServices
}

func New(dbclient *ent.Client) *Service {
	return &Service{
		Auth: &AuthService{
			Client: dbclient.Auth,
		},
		User: &UserService{
			Client: dbclient.User,
		},
		Serv: &ServiceServices{
			Client:    dbclient.Service,
			LocClient: dbclient.Location,
			CatClient: dbclient.Category,
		},
	}
}
