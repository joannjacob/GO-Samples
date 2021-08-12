package repository

import (
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/2019/PPLB4/back-end/domain/entity"
)

// DuckRepository represents a storage of all existing ducks
type DuckRepository interface {
	Get(ID int64) (*entity.Duck, error)
	GetAll() ([]*entity.Duck, error)
	Save(duck *entity.Duck) error
}

var duckRepository DuckRepository

// GetDuckRepository returns the DuckRepository
func GetDuckRepository() DuckRepository {
	return duckRepository
}

// InitDuckRepository injects DuckRepository with its implementation
func InitDuckRepository(r DuckRepository) {
	duckRepository = r
}
