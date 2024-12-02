package services

import (
	"context"
	"time"

	"myary/modules/diaries/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Interface & Struct
type DiaryService interface {
	Insert(diary models.DiaryModel) (*mongo.InsertOneResult, error)
	GetAll() ([]models.DiaryModel, error)
	Update(filter bson.M, update bson.M) (*mongo.UpdateResult, error)
	Delete(filter bson.M) (*mongo.DeleteResult, error)
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
	diary.CreatedAt = time.Now()
	diary.UpdatedAt = nil

	return r.collection.InsertOne(context.TODO(), diary)
}
func (s *diaryService) Update(filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	if _, ok := update["$set"]; !ok {
		update["$set"] = bson.M{}
	}

	setFields, ok := update["$set"].(bson.M)
	if !ok {
		setFields = bson.M{}
		for key, value := range update["$set"].(map[string]interface{}) {
			setFields[key] = value
		}
		update["$set"] = setFields
	}

	setFields["updated_at"] = time.Now()

	return s.collection.UpdateOne(context.TODO(), filter, update)
}
func (s *diaryService) Delete(filter bson.M) (*mongo.DeleteResult, error) {
	return s.collection.DeleteOne(context.TODO(), filter)
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
