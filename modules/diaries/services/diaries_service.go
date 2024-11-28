package services

import (
	"context"

	"myary/modules/diaries/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Interface & Struct
type DiaryService interface {
	Insert(diary models.DiaryModel) (*mongo.InsertOneResult, error)
	GetAll() ([]models.DiaryModel, error)
}
type diaryService struct {
	collection *mongo.Collection
}

func NewDiaryService(db *mongo.Database) DiaryService {
	return &diaryService{
		collection: db.Collection("diaries"),
	}
}

// Command Service
func (r *diaryService) Insert(diary models.DiaryModel) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.TODO(), diary)
}

// Query Service
func (r *diaryService) GetAll() ([]models.DiaryModel, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}

	var diaries []models.DiaryModel
	if err := cursor.All(context.TODO(), &diaries); err != nil {
		return nil, err
	}

	return diaries, nil
}
