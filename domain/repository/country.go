package repository
import (
	"presentation/domain/entity"
)

type CountryRepository interface {
	CreateCountry (entity.CountryStruct) (interface{}, error)
}
