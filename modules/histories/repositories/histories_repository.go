package repositories

import (
	"myary/modules/histories/models"
	"myary/modules/histories/services"
)

// Interface & Struct
type HistoryService interface {
	CreateHistory(history models.HistoryModel) error
	FetchHistories() ([]models.HistoryModel, error)
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

// Query Repo
func (s *historyService) FetchHistories() ([]models.HistoryModel, error) {
	return s.repo.GetAll()
}
