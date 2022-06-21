package utils

import "github.com/google/uuid"

type UserInput struct {
	Name             string `json:"name" validate:"required"`
	Username         string `json:"username" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required,min=8" `
	Confirm_password string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UserSigninInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
type ErrorResponse struct {
	Msg string `json:"msg"`
}
type SuccessResponse struct {
	AccessToken  *string                `json:"access_token"`
	RefreshToken *string                `json:"refresh_token"`
	Data         map[string]interface{} `json:"data"`
}
type SessionData struct {
	ID uuid.UUID
	Ip string
}

type ServiceInput struct {
	Title          string   `json:"title" validate:"required"`
	Instrument     []string `json:"instruments"`
	Locationid     string   `json:"location" validate:"required"`
	CatID          string   `json:"category" validate:"required"`
	Cost           string   `json:"cost" validate:"required"`
	AdditionalCost string   `json:"additionalcost"`
}

type OrderInput struct {
	ServiceID string   `json:"serviceid" validate:"required"`
	Adsress   []string `json:"address" validate:"required"`
}
type OrderDetails struct {
	ServiceID  uuid.UUID
	ProviderID uuid.UUID
	TotalCost  float64
	Adress     []string
}
