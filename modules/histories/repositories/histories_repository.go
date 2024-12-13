package repositories

import (
	"myary/modules/histories/models"
	"myary/modules/histories/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface & Struct
type HistoryRepository interface {
	CreateHistory(history models.HistoryModel) error
	FetchHistories() ([]models.HistoryModel, error)
	DeleteHistory(id primitive.ObjectID) (*mongo.DeleteResult, error)
}
type historyService struct {
	service services.HistoryService
}

func NewHistoryService(service services.HistoryService) HistoryRepository {
	return &historyService{service: service}
}

// Command Repo
func (s *historyService) CreateHistory(history models.HistoryModel) error {
	_, err := s.service.Insert(history)
	return err
}
func (s *historyService) DeleteHistory(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	return s.service.Delete(filter)
}

// Query Repo
func (s *historyService) FetchHistories() ([]models.HistoryModel, error) {
	return s.service.GetAll()
}
