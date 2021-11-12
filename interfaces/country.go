package interfaces

import (
	"encoding/json"
	"errors"
	"net/http"
	"presentation/application"
	"presentation/domain/entity"
)

type CountryInterface struct {
	cu application.CountryApplication
}

func NewCountry(cu application.CountryApplication) CountryInterface {
	return CountryInterface{
		cu: cu,
	}
}

func (cu *CountryInterface) CreateCountry(w http.ResponseWriter, r *http.Request) {
	var country entity.CountryStruct

	json.NewDecoder(r.Body).Decode(&country)
	newCountry, err := cu.cu.CreateCountry(country)
	if err != nil {
		responseError := errors.New("unable to register country")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(responseError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCountry)
}
