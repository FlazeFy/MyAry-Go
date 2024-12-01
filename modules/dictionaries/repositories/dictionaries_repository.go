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
	service          services.DictionaryService
	firestoreService services.DictionaryService
}

func NewDictionaryService(service services.DictionaryService) DictionaryService {
	return &dictionaryService{service: service}
}

// Command Repo
func (s *dictionaryService) CreateDictionary(dictionary models.DictionaryModel) error {
	// MongoDB
	_, _, err := s.service.Insert(dictionary)
	if err != nil {
		return err
	}

	// Firestore
	_, _, err = s.firestoreService.Insert(dictionary)
	if err != nil {
		return err
	}

	return nil
}

// Query Repo
func (s *dictionaryService) FetchDictionaries() ([]models.DictionaryModel, error) {
	return s.service.GetAll()
}
