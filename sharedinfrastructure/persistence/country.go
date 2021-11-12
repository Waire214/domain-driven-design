package persistence

import (
	"errors"
	"presentation/domain/entity"
	"presentation/domain/repository"

	"github.com/jinzhu/gorm"
)

type CountryInfra struct {
	database *gorm.DB
}

func NewCountryInfra(database *gorm.DB) *CountryInfra {
	return &CountryInfra{database}
}

//CountryInfra implements the repository.CountryRepository interface
var _ repository.CountryRepository = &CountryInfra{}

func (r *CountryInfra) CreateCountry(c entity.CountryStruct) (interface{}, error) {
	if c.Name == "" {
		return entity.CountryStruct{}, errors.New("invalid input")
	}
	r.database.Create(c)
	return c, nil
}
