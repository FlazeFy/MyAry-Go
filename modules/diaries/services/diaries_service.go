package services

import (
	"context"
	"fmt"
	"time"

	"myary/modules/diaries/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Interface & Struct
type DiaryService interface {
	Insert(diary models.DiaryModel) (*mongo.InsertOneResult, error)
	GetAll() ([]models.DiaryModel, error)
	Update(filter bson.M, update bson.M) (*mongo.UpdateResult, error)
	Delete(filter bson.M) (*mongo.DeleteResult, error)
	GetStatsDiaryLifetime() (models.StatsDiaryLifetimeModel, error)
	GetOneById(id string) (models.DiaryModel, error)
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
func (r *diaryService) GetOneById(id string) (models.DiaryModel, error) {
	var diary models.DiaryModel

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return diary, fmt.Errorf("invalid ID format: %v", err)
	}

	filter := bson.M{"_id": objectID}
	err = r.collection.FindOne(context.TODO(), filter).Decode(&diary)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return diary, fmt.Errorf("diary not found")
		}
		return diary, err
	}

	return diary, nil
}
func (r *diaryService) GetStatsDiaryLifetime() (models.StatsDiaryLifetimeModel, error) {
	pipeline := mongo.Pipeline{
		{
			{
				Key: "$group", Value: bson.D{
					{Key: "_id", Value: nil},
					{Key: "total", Value: bson.D{{Key: "$sum", Value: 1}}},
					{Key: "average_mood", Value: bson.D{{Key: "$avg", Value: "$diary_mood"}}},
					{Key: "average_tired", Value: bson.D{{Key: "$avg", Value: "$diary_tired"}}},
				}},
		},
	}

	cursor, err := r.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return models.StatsDiaryLifetimeModel{}, err
	}

	var results []models.StatsDiaryLifetimeModel
	if err := cursor.All(context.TODO(), &results); err != nil {
		return models.StatsDiaryLifetimeModel{}, err
	}

	if len(results) == 0 {
		return models.StatsDiaryLifetimeModel{
			Total:        0,
			AverageMood:  0,
			AverageTired: 0,
		}, nil
	}

	return results[0], nil
}
