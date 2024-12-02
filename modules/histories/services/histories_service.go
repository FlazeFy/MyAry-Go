package services

import (
	"context"
	"myary/modules/histories/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Interface & Struct
type HistoryService interface {
	Insert(history models.HistoryModel) (*mongo.InsertOneResult, error)
	GetAll() ([]models.HistoryModel, error)
	Delete(filter bson.M) (*mongo.DeleteResult, error)
}
type historyService struct {
	collection *mongo.Collection
}

func NewHistoryService(db *mongo.Database) HistoryService {
	return &historyService{
		collection: db.Collection("histories"),
	}
}

// Command Service
func (r *historyService) Insert(history models.HistoryModel) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.TODO(), history)
}
func (s *historyService) Delete(filter bson.M) (*mongo.DeleteResult, error) {
	return s.collection.DeleteOne(context.TODO(), filter)
}

// Query Service
func (r *historyService) GetAll() ([]models.HistoryModel, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}

	var histories []models.HistoryModel
	if err := cursor.All(context.TODO(), &histories); err != nil {
		return nil, err
	}

	return histories, nil
}
