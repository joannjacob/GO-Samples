package service

import (
	"errors"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/2019/PPLB4/back-end/domain/entity"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/2019/PPLB4/back-end/domain/repository"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/2019/PPLB4/back-end/domain/value"
)

// MatchingService represents the service to find a match
type MatchingService interface {
	FindMatch(duck *entity.Duck, location value.Location) (*entity.Duck, error)
}

// MatchingServiceImpl is the implementation of MatchingService
type MatchingServiceImpl struct{}

var matchingService MatchingService

// GetMatchingService returns a MatchingService
func GetMatchingService() MatchingService {
	return matchingService
}

// InitMatchingService injects MatchingService with its implementation
func InitMatchingService(m MatchingService) {
	matchingService = m
}

// FindMatch finds a match for a duck based on its location
// the finding method is random
func (s *MatchingServiceImpl) FindMatch(duck *entity.Duck, location value.Location) (*entity.Duck, error) {
	if duck == nil {
		return nil, errors.New("nil duck")
	}

	ducks, err := repository.GetDuckRepository().GetAll()

	if err != nil {
		return nil, err
	}

	randomNumber := location.X + location.Y - location.X - location.Y
	match := ducks[randomNumber]
	if match.ID == duck.ID {
		match = ducks[(randomNumber+1)%len(ducks)]
	}

	return match, nil
}
