package application

import (
	"presentation/domain/entity"
	"presentation/domain/repository"
)

type CountryApp struct {
	theCountry repository.CountryRepository
}

var _CountryApplication = &CountryApp{}

type CountryApplication interface {
	CreateCountry (entity.CountryStruct) (interface{}, error)
}

func (c *CountryApp) CreateCountry(country entity.CountryStruct) (interface{}, error) {
	return c.theCountry.CreateCountry(country)
}