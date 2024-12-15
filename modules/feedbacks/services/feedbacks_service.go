package services

import (
	"context"
	"myary/modules/feedbacks/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Interface & Struct
type FeedbackService interface {
	Insert(feedback models.FeedbackModel) (*mongo.InsertOneResult, error)
	GetAll() ([]models.FeedbackModel, error)
	Delete(filter bson.M) (*mongo.DeleteResult, error)
	GetFeedBackStats() ([]models.StatsFeedback, error)
}

type feedbackService struct {
	collection *mongo.Collection
}

func NewFeedbackService(db *mongo.Database) FeedbackService {
	return &feedbackService{
		collection: db.Collection("feedbacks"),
	}
}

// Command Service
func (r *feedbackService) Insert(feedback models.FeedbackModel) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.TODO(), feedback)
}
func (s *feedbackService) Delete(filter bson.M) (*mongo.DeleteResult, error) {
	return s.collection.DeleteOne(context.TODO(), filter)
}

// Query Service
func (r *feedbackService) GetAll() ([]models.FeedbackModel, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}

	var feedbacks []models.FeedbackModel
	if err := cursor.All(context.TODO(), &feedbacks); err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func (r *feedbackService) GetFeedBackStats() ([]models.StatsFeedback, error) {
	pipeline := mongo.Pipeline{
		{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$feedback_rate"},
				{Key: "total", Value: bson.D{{Key: "$sum", Value: 1}}},
			}},
		},
		{
			{Key: "$project", Value: bson.D{
				{Key: "context", Value: "$_id"},
				{Key: "total", Value: 1},
				{Key: "_id", Value: 0},
			}},
		},
	}

	cursor, err := r.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	var results []models.StatsFeedback
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
