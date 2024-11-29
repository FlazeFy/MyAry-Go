package http_handlers

import (
	"net/http"

	"myary/modules/diaries/models"
	"myary/modules/diaries/repositories"

	"github.com/gin-gonic/gin"
)

// Interface & Struct
type DiaryHandler struct {
	repo repositories.DiaryService
}

func NewDiaryHandler(repo repositories.DiaryService) *DiaryHandler {
	return &DiaryHandler{repo: repo}
}

// Command Handler
func (h *DiaryHandler) CreateDiary(c *gin.Context) {
	var diary models.DiaryModel
	if err := c.ShouldBindJSON(&diary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateDiary(diary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong, please call admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Diary created"})
}

// Query Handler
func (h *DiaryHandler) GetDiaries(c *gin.Context) {
	diaries, err := h.repo.FetchDiaries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong, please call admin"})
		return
	}

	if len(diaries) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Diary found", "data": diaries})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Diary not found", "data": nil})
	}
}
