package services

import (
	"context"
	"strconv"

	"github.com/ArifulProtik/sheba-api/ent"
	"github.com/ArifulProtik/sheba-api/ent/category"
	"github.com/ArifulProtik/sheba-api/ent/location"
	"github.com/ArifulProtik/sheba-api/ent/service"
	"github.com/ArifulProtik/sheba-api/ent/user"
	"github.com/ArifulProtik/sheba-api/pkg/utils"
	"github.com/google/uuid"
)

type ServiceServices struct {
	Client    *ent.ServiceClient
	LocClient *ent.LocationClient
	CatClient *ent.CategoryClient
}

func (s *ServiceServices) GetLocAndCat(locid int, catid int) (*ent.Location, *ent.Category, error) {
	loc, err := s.LocClient.Query().Where(location.IDEQ(locid)).First(context.Background())
	if err != nil {
		return nil, nil, err
	}
	cat, err := s.CatClient.Query().Where(category.IDEQ(catid)).First(context.Background())
	if err != nil {
		return nil, nil, err
	}
	return loc, cat, nil
}

func (s *ServiceServices) CreateLoc(loc string) (*ent.Location, error) {
	newloc, err := s.LocClient.Create().SetLocation(loc).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return newloc, nil
}
func (s *ServiceServices) CreateCat(cat string) (*ent.Category, error) {
	newcat, err := s.CatClient.Create().SetName(cat).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return newcat, nil
}

func (s *ServiceServices) CreateServices(ser utils.ServiceInput, usr *ent.User) (*ent.Service, error) {
	locid, err := strconv.Atoi(ser.Locationid)
	if err != nil {
		return nil, err
	}
	catid, err := strconv.Atoi(ser.CatID)
	if err != nil {
		return nil, err
	}
	cost, err := strconv.ParseFloat(ser.Cost, 64)
	if err != nil {
		return nil, err
	}
	var a_cost float64
	if ser.AdditionalCost == "" {
		a_cost = 0
	} else {
		aa_cost, err := strconv.ParseFloat(ser.AdditionalCost, 64)
		if err != nil {
			return nil, err
		}
		a_cost = aa_cost
	}
	if ser.Instrument == nil {
		ser.Instrument = []string{}
	}

	newservice, err := s.Client.Create().
		SetUser(usr).
		SetTitle(ser.Title).
		SetCategoryid(catid).
		SetLocationid(locid).
		SetCost(cost).
		SetInstrument(ser.Instrument).
		SetAdditionalCost(a_cost).
		Save(context.Background())

	if err != nil {
		return nil, err
	}
	return newservice, nil
}

func (s *ServiceServices) DeleteService(id uuid.UUID) error {
	err := s.Client.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceServices) GetServiceByID(id uuid.UUID) (*ent.Service, error) {
	ser, err := s.Client.Query().WithUser().
		Where(service.IDEQ(id)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return ser, nil
}

func (s *ServiceServices) AllService() ([]*ent.Service, error) {
	sers, err := s.Client.Query().Order(ent.Desc(service.FieldCreatedAt)).
		WithUser().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return sers, nil
}
func (s *ServiceServices) ServicesByUser(id uuid.UUID) ([]*ent.Service, error) {

	sers, err := s.Client.Query().Where(service.HasUserWith(user.IDEQ(id))).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return sers, nil
}
