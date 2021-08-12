package application

import (
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/2019/PPLB4/back-end/domain/entity"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/2019/PPLB4/back-end/domain/repository"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/2019/PPLB4/back-end/domain/service"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/2019/PPLB4/back-end/domain/value"
)

// Tinduck represents Tinduck application to be called by interface layer
type Tinduck interface {
	GetDuck(id int64) (*entity.Duck, error)
	GetDucks() ([]*entity.Duck, error)
	AddDuck(name string) error
	GetDuckMatch(id int64, x, y int) (*entity.Duck, error)
}

// TinduckImpl is the implementation of Tinduck
type TinduckImpl struct{}

var tinduck Tinduck

// GetTinduck returns Tinduck application
func GetTinduck() Tinduck {
	return tinduck
}

// InitTinduck injects implementation for Tinduck application
func InitTinduck(t Tinduck) {
	tinduck = t
}

// GetDuck returns duck with the given id
func (t *TinduckImpl) GetDuck(id int64) (*entity.Duck, error) {
	return repository.GetDuckRepository().Get(id)
}

// GetDucks returns ducks stored in repository
func (t *TinduckImpl) GetDucks() ([]*entity.Duck, error) {
	return repository.GetDuckRepository().GetAll()
}

// AddDuck inserts new duck with the given name to repository
func (t *TinduckImpl) AddDuck(name string) error {
	return repository.GetDuckRepository().Save(&entity.Duck{
		Name: name,
	})
}

// GetDuckMatch gets a matching duck from MatchingService given a duck's id and x, y location
func (t *TinduckImpl) GetDuckMatch(id int64, x, y int) (*entity.Duck, error) {
	duck, err := repository.GetDuckRepository().Get(id)

	if err != nil {
		return nil, err
	}

	return service.GetMatchingService().FindMatch(duck, value.Location{
		X: x,
		Y: y,
	})
}