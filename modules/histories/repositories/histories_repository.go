package repositories

import (
	"myary/modules/histories/models"
	"myary/modules/histories/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface & Struct
type HistoryService interface {
	CreateHistory(history models.HistoryModel) error
	FetchHistories() ([]models.HistoryModel, error)
	DeleteHistory(id primitive.ObjectID) (*mongo.DeleteResult, error)
}
type historyService struct {
	repo services.HistoryService
}

func NewHistoryService(repo services.HistoryService) HistoryService {
	return &historyService{repo: repo}
}

// Command Repo
func (s *historyService) CreateHistory(history models.HistoryModel) error {
	_, err := s.repo.Insert(history)
	return err
}
func (s *historyService) DeleteHistory(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	return s.repo.Delete(filter)
}

// Query Repo
func (s *historyService) FetchHistories() ([]models.HistoryModel, error) {
	return s.repo.GetAll()
}
