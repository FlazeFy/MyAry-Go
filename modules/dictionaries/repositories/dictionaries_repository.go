package repositories

import (
	"myary/modules/dictionaries/models"
	"myary/modules/dictionaries/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface & Struct
type DictionaryService interface {
	CreateDictionary(dictionary models.DictionaryModel) error
	FetchDictionaries() ([]models.DictionaryModel, error)
	DeleteDictionary(id primitive.ObjectID) (*mongo.DeleteResult, error)
	FetchTotalDictionaryUsed() ([]models.StatsDictionary, error)
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
func (s *dictionaryService) DeleteDictionary(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	return s.service.Delete(filter)
}

// Query Repo
func (s *dictionaryService) FetchDictionaries() ([]models.DictionaryModel, error) {
	dictionaries, err := s.service.GetAll()
	if err != nil {
		return nil, err
	}
	return dictionaries, nil
}
func (s *dictionaryService) FetchTotalDictionaryUsed() ([]models.StatsDictionary, error) {
	dictionaries, err := s.service.GetTotalDictionaryUsed()
	if err != nil {
		return nil, err
	}
	return dictionaries, nil
}
