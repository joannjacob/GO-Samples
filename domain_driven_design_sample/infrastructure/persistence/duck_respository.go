package persistence

import (
	"errors"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/2019/PPLB4/back-end/domain/entity"
)

// DuckRepositoryImpl is the implementation of DuckRepository
type DuckRepositoryImpl struct {
	DB map[int64]*entity.Duck
}

// Get returns a duck from the database with the id
func (r *DuckRepositoryImpl) Get(id int64) (*entity.Duck, error) {
	if r.DB == nil {
		return nil, errors.New("database error")
	}

	if r.DB[id] == nil {
		return nil, errors.New("duck not found")
	}

	return r.DB[id], nil
}

// GetAll return all ducks stored in database
func (r *DuckRepositoryImpl) GetAll() ([]*entity.Duck, error) {
	if r.DB == nil {
		return nil, errors.New("database error")
	}

	ducks := []*entity.Duck{}
	for _, duck := range r.DB {
		ducks = append(ducks, duck)
	}

	return ducks, nil
}

// Save insert a duck to database
func (r *DuckRepositoryImpl) Save(duck *entity.Duck) error {
	if duck == nil {
		return errors.New("nil duck")
	}
	if r.DB == nil {
		return errors.New("database error")
	}

	duck.ID = int64(len(r.DB) + 1)
	r.DB[duck.ID] = duck
	return nil
}
