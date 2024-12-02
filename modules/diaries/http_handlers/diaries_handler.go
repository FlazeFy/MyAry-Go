package http_handlers

import (
	"net/http"

	"myary/modules/diaries/models"
	"myary/modules/diaries/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (h *DiaryHandler) UpdateDiary(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.repo.UpdateDiary(models.DiaryModel{ID: objectID}, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update diary"})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Diary not found or no changes applied"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Diary updated successfully"})
	}
}
func (h *DiaryHandler) DeleteDiary(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := h.repo.DeleteDiary(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete diary"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Diary not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Diary deleted successfully"})
	}
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
