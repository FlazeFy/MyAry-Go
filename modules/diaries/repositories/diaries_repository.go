package repositories

import (
	"fmt"
	"myary/modules/diaries/models"
	"myary/modules/diaries/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface & Struct
type DiaryService interface {
	CreateDiary(diary models.DiaryModel) error
	FetchDiaries() ([]models.DiaryModel, error)
	UpdateDiary(diary models.DiaryModel, updates map[string]interface{}) (*mongo.UpdateResult, error)
	DeleteDiary(id primitive.ObjectID) (*mongo.DeleteResult, error)
	FetchDiaryStatsLifetime() (models.StatsDiaryLifetimeModel, error)
	FetchDiaryById(id string) (*models.DiaryModel, error)
}
type diaryService struct {
	repo services.DiaryService
}

func NewDiaryService(repo services.DiaryService) DiaryService {
	return &diaryService{repo: repo}
}

// Command Repo
func (s *diaryService) CreateDiary(diary models.DiaryModel) error {
	_, err := s.repo.Insert(diary)
	return err
}
func (s *diaryService) UpdateDiary(diary models.DiaryModel, updates map[string]interface{}) (*mongo.UpdateResult, error) {
	// Define the allowed fields
	allowedFields := []string{"diary_title", "diary_mood", "diary_tired", "diary_date"}
	allowedUpdates := bson.M{}

	for _, field := range allowedFields {
		if value, exists := updates[field]; exists {
			allowedUpdates[field] = value
		}
	}

	filter := bson.M{"_id": diary.ID}
	update := bson.M{
		"$set": allowedUpdates,
	}

	return s.repo.Update(filter, update)
}
func (s *diaryService) DeleteDiary(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	return s.repo.Delete(filter)
}

// Query Repo
func (s *diaryService) FetchDiaries() ([]models.DiaryModel, error) {
	diary, err := s.repo.GetAll()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return diary, nil
}
func (s *diaryService) FetchDiaryStatsLifetime() (models.StatsDiaryLifetimeModel, error) {
	return s.repo.GetStatsDiaryLifetime()
}
func (s *diaryService) FetchDiaryById(id string) (*models.DiaryModel, error) {
	diary, err := s.repo.GetOneById(id)
	if err != nil {
		return nil, err
	}
	return &diary, nil
}
