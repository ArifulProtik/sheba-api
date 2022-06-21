package controller

import (
	"log"
	"net/http"

	"github.com/ArifulProtik/sheba-api/ent"
	"github.com/ArifulProtik/sheba-api/pkg/services"
	"github.com/ArifulProtik/sheba-api/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ServiceController struct {
	Service *services.ServiceServices
}

func (s *ServiceController) CreateService(c echo.Context) error {
	if c.Request().Body != nil {
		user := c.Get("user").(*ent.User)
		var inpt utils.ServiceInput

		if err := c.Bind(&inpt); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
				Msg: "Missing Data",
			})
		}
		err := utils.NewValidator().Struct(inpt)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, utils.ToErrResponse(err))
		}
		ser, err := s.Service.CreateServices(inpt, user)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
				Msg: "Internal Server Error",
			})
		}
		return c.JSON(http.StatusCreated, ser)
	}
	return c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
		Msg: "JSON data missing",
	})
}

func (s *ServiceController) DeleteService(c echo.Context) error {
	id := c.Param("id")
	user := c.Get("user").(*ent.User)
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Msg: "service Not found",
		})
	}

	ser, err := s.Service.GetServiceByID(uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Msg: "service Not found",
		})
	}
	if ser.Edges.User.ID != user.ID {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Msg: "StatusUnauthorized",
		})
	}
	return c.JSON(http.StatusAccepted, echo.Map{
		"msg": "deleted services",
	})
}

func (s *ServiceController) GetServices(c echo.Context) error {
	servs, err := s.Service.AllService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Msg: "Internal",
		})
	}
	return c.JSON(http.StatusOK, servs)
}

func (s *ServiceController) MyService(c echo.Context) error {
	user := c.Get("user").(*ent.User)

	servs, err := s.Service.ServicesByUser(user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Msg: "No services Found",
		})
	}
	return c.JSON(http.StatusOK, servs)
}

func (s *ServiceController) CreateLocation(c echo.Context) error {
	if c.Request().Body != nil {
		type locinpt struct {
			Name string `json:"name"`
		}
		var l locinpt
		if err := c.Bind(&l); err != nil {
			return c.JSON(500, echo.Map{"msg": "Nope"})
		}
		loc, err := s.Service.CreateLoc(l.Name)
		if err != nil {
			return c.JSON(500, err.Error())
		}
		return c.JSON(http.StatusCreated, loc)
	}
	return c.JSON(500, echo.Map{"msg": "json missing"})
}

func (s *ServiceController) CreateCat(c echo.Context) error {
	if c.Request().Body != nil {
		type locinpt struct {
			Name string `json:"name"`
		}
		var l locinpt
		if err := c.Bind(&l); err != nil {
			return c.JSON(500, echo.Map{"msg": "Nope"})
		}
		loc, err := s.Service.CreateCat(l.Name)
		if err != nil {
			return c.JSON(500, err.Error())
		}
		return c.JSON(http.StatusCreated, loc)
	}
	return c.JSON(500, echo.Map{"msg": "json missing"})
}
