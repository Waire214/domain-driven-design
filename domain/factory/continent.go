package factory

import "presentation/domain/entity"

type ContinentFactory struct {
	Name    string `json:"name"`
	Country entity.CountryStruct `json:"country"`
}