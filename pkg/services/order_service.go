package services

import (
	"context"

	"github.com/ArifulProtik/sheba-api/ent"
	"github.com/ArifulProtik/sheba-api/ent/order"
	"github.com/ArifulProtik/sheba-api/ent/user"
	"github.com/ArifulProtik/sheba-api/pkg/utils"
	"github.com/google/uuid"
)

type OrderService struct {
	Client *ent.OrderClient
}

func (o *OrderService) CreateOrder(details utils.OrderDetails, usr *ent.User) (*ent.Order, error) {
	order, err := o.Client.Create().
		SetUser(usr).
		SetServiceid(details.ServiceID).
		SetProviderid(details.ProviderID).
		SetTotalcost(details.TotalCost).
		SetAddress(details.Adress).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *OrderService) UpdateOrderAcc(id uuid.UUID) error {
	err := o.Client.UpdateOneID(id).
		SetIsAccepted(true).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderService) UpdateOrderDec(id uuid.UUID) error {
	err := o.Client.UpdateOneID(id).SetIsDeclined(true).Exec(context.Background())
	if err != nil {
		return nil
	}
	return nil
}

func (o *OrderService) OrdersOfProvider(id uuid.UUID) ([]*ent.Order, error) {
	orders, err := o.Client.Query().Where(order.ProvideridEQ(id)).
		WithUser().All(context.Background())
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *OrderService) OrderOfUser(id uuid.UUID) ([]*ent.Order, error) {
	orders, err := o.Client.Query().Where(order.HasUserWith(user.IDEQ(id))).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *OrderService) GetOrderByID(id uuid.UUID) (*ent.Order, error) {
	order, err := o.Client.Query().Where(order.IDEQ(id)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *OrderService) DeleteOrder(id uuid.UUID) error {
	err := o.Client.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
