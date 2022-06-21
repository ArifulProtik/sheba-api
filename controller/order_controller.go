package controller

import (
	"net/http"

	"github.com/ArifulProtik/sheba-api/ent"
	"github.com/ArifulProtik/sheba-api/pkg/services"
	"github.com/ArifulProtik/sheba-api/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	Service *services.OrderService
	Serv    *services.ServiceServices
}

func (o *OrderController) CreateOrder(c echo.Context) error {
	if c.Request().Body != nil {
		user := c.Get("user").(*ent.User)
		var requestbody utils.OrderInput
		if err := c.Bind(&requestbody); err != nil {
			return c.JSON(http.StatusUnavailableForLegalReasons, utils.ErrorResponse{
				Msg: "Missing Data",
			})
		}
		servviceuid, err := uuid.Parse(requestbody.ServiceID)
		if err != nil {
			return c.JSON(http.StatusNotFound, utils.ErrorResponse{
				Msg: "No Service Found By ID",
			})
		}
		serv, err := o.Serv.GetServiceByID(servviceuid)
		if err != nil {
			return c.JSON(http.StatusNotFound, utils.ErrorResponse{
				Msg: "No Service Found",
			})
		}
		totalcost := serv.Cost + serv.AdditionalCost
		var od utils.OrderDetails
		od = utils.OrderDetails{
			ServiceID:  serv.ID,
			ProviderID: serv.Edges.User.ID,
			TotalCost:  totalcost,
			Adress:     requestbody.Adsress,
		}
		neworder, err := o.Service.CreateOrder(od, user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
				Msg: err.Error(),
			})
		}
		return c.JSON(http.StatusCreated, neworder)

	}
	return c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
		Msg: "missing data",
	})
}

func (o *OrderController) ServiceRequests(c echo.Context) error {
	user := c.Get("user").(*ent.User)

	orders, err := o.Service.OrdersOfProvider(user.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"msg": "No Request Found",
		})
	}
	return c.JSON(http.StatusOK, orders)
}

func (o *OrderController) AcceptOrder(c echo.Context) error {
	user := c.Get("user").(*ent.User)
	oid := c.Param("id")
	ouid, err := uuid.Parse(oid)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Msg: "No Service Found By ID",
		})
	}
	order, err := o.Service.GetOrderByID(ouid)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Msg: "No Order Found By ID",
		})
	}
	serv, _ := o.Serv.GetServiceByID(order.Serviceid)
	if user.ID != serv.Edges.User.ID {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Msg: "StatusUnauthorized",
		})
	}
	err = o.Service.UpdateOrderAcc(order.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Msg: "Internal Error",
		})
	}
	return c.JSON(http.StatusAccepted, echo.Map{"msg": "accepted"})

}

func (o *OrderController) DeclineOrder(c echo.Context) error {
	user := c.Get("user").(*ent.User)
	oid := c.Param("id")
	ouid, err := uuid.Parse(oid)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Msg: "No Service Found By ID",
		})
	}
	order, err := o.Service.GetOrderByID(ouid)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Msg: "No Order Found By ID",
		})
	}
	serv, _ := o.Serv.GetServiceByID(order.Serviceid)
	if user.ID != serv.Edges.User.ID {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Msg: "StatusUnauthorized",
		})
	}
	err = o.Service.UpdateOrderDec(order.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Msg: "Internal Error",
		})
	}
	return c.JSON(http.StatusAccepted, echo.Map{"msg": "declined"})

}

func (o *OrderController) OrderHistory(c echo.Context) error {
	usr := c.Get("user").(*ent.User)
	orders, err := o.Service.OrderOfUser(usr.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Msg: "Internal Error",
		})
	}
	return c.JSON(http.StatusOK, orders)
}
