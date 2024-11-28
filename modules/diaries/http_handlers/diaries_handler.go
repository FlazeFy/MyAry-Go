package http_handlers

import (
	"net/http"

	"myary/modules/diaries/models"
	"myary/modules/diaries/repositories"

	"github.com/gin-gonic/gin"
)

// Interface & Struct
type DiaryHandler struct {
	service repositories.DiaryService
}

func NewDiaryHandler(service repositories.DiaryService) *DiaryHandler {
	return &DiaryHandler{service: service}
}

// Command Handler
func (h *DiaryHandler) CreateDiary(c *gin.Context) {
	var diary models.DiaryModel
	if err := c.ShouldBindJSON(&diary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateDiary(diary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create diary"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Diary created"})
}

// Query Handler
func (h *DiaryHandler) GetDiaries(c *gin.Context) {
	diaries, err := h.service.FetchDiaries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch diary"})
		return
	}

	c.JSON(http.StatusOK, diaries)
}
