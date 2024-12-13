package repositories

import (
	"myary/modules/feedbacks/models"
	"myary/modules/feedbacks/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface & Struct
type FeedbackRepository interface {
	CreateFeedback(history models.FeedbackModel) error
	FetchFeedbacks() ([]models.FeedbackModel, error)
	DeleteFeedback(id primitive.ObjectID) (*mongo.DeleteResult, error)
}
type feedbackService struct {
	service services.FeedbackService
}

func NewFeedbackService(service services.FeedbackService) FeedbackRepository {
	return &feedbackService{service: service}
}

// Command Repo
func (s *feedbackService) CreateFeedback(history models.FeedbackModel) error {
	_, err := s.service.Insert(history)
	return err
}
func (s *feedbackService) DeleteFeedback(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	return s.service.Delete(filter)
}

// Query Repo
func (s *feedbackService) FetchFeedbacks() ([]models.FeedbackModel, error) {
	return s.service.GetAll()
}
