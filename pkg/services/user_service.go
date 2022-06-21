package services

import (
	"context"

	"github.com/ArifulProtik/sheba-api/ent"
	"github.com/ArifulProtik/sheba-api/ent/user"
	"github.com/ArifulProtik/sheba-api/pkg/utils"
	"github.com/google/uuid"
)

type UserService struct {
	Client *ent.UserClient
}

func (a *UserService) SaveUser(usr utils.UserInput) (*ent.User, error) {
	newpass, _ := utils.HashBeforeSave(usr.Password)
	newusr, err := a.Client.Create().SetName(usr.Name).SetEmail(usr.Email).SetUsername(
		usr.Username,
	).SetPassword(string(newpass)).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return newusr, nil

}

func (a *UserService) FindUserByEmail(email string) (*ent.User, error) {
	usr, err := a.Client.Query().Where(user.EmailEQ(email)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return usr, nil
}
func (a *UserService) FindUserByID(id uuid.UUID) (*ent.User, error) {
	usr, err := a.Client.Query().Where(user.IDEQ(id)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return usr, nil
}
