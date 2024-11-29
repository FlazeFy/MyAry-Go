package repositories

import (
	"myary/modules/dictionaries/models"
	"myary/modules/dictionaries/services"
)

// Interface & Struct
type DictionaryService interface {
	CreateDictionary(dictionary models.DictionaryModel) error
	FetchDictionaries() ([]models.DictionaryModel, error)
}
type dictionaryService struct {
	repo services.DictionaryService
}

func NewDictionaryService(repo services.DictionaryService) DictionaryService {
	return &dictionaryService{repo: repo}
}

// Command Repo
func (s *dictionaryService) CreateDictionary(dictionary models.DictionaryModel) error {
	_, err := s.repo.Insert(dictionary)
	return err
}

// Query Repo
func (s *dictionaryService) FetchDictionaries() ([]models.DictionaryModel, error) {
	return s.repo.GetAll()
}
