package repositories

import (
	"myary/modules/diaries/models"
	"myary/modules/diaries/services"
)

// Interface & Struct
type DiaryService interface {
	CreateDiary(diary models.DiaryModel) error
	FetchDiaries() ([]models.DiaryModel, error)
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

// Query Repo
func (s *diaryService) FetchDiaries() ([]models.DiaryModel, error) {
	return s.repo.GetAll()
}
