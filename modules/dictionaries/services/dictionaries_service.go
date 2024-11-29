package services

import (
	"context"
	"myary/modules/dictionaries/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Interface & Struct
type DictionaryService interface {
	Insert(dictionary models.DictionaryModel) (*mongo.InsertOneResult, error)
	GetAll() ([]models.DictionaryModel, error)
}
type dictionaryService struct {
	collection *mongo.Collection
}

func NewDictionaryService(db *mongo.Database) DictionaryService {
	return &dictionaryService{
		collection: db.Collection("dictionaries"),
	}
}

// Command Service
func (r *dictionaryService) Insert(dictionary models.DictionaryModel) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.TODO(), dictionary)
}

// Query Service
func (r *dictionaryService) GetAll() ([]models.DictionaryModel, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}

	var dictionaries []models.DictionaryModel
	if err := cursor.All(context.TODO(), &dictionaries); err != nil {
		return nil, err
	}

	return dictionaries, nil
}
