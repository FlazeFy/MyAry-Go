package httphandlers

import (
	"myary/modules/histories/models"
	"myary/modules/histories/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Interface & Struct
type HistoryHandler struct {
	repo repositories.HistoryRepository
}

func NewHistoryHandler(repo repositories.HistoryRepository) *HistoryHandler {
	return &HistoryHandler{repo: repo}
}

// Command Handler
func (h *HistoryHandler) CreateHistory(c *gin.Context) {
	var history models.HistoryModel
	if err := c.ShouldBindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.repo.CreateHistory(history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong, please call admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "History created"})
}
func (h *HistoryHandler) DeleteHistory(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := h.repo.DeleteHistory(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong, please call admin"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "History not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "History deleted successfully"})
	}
}

// Query Handler
func (h *HistoryHandler) GetHistories(c *gin.Context) {
	histories, err := h.repo.FetchHistories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong, please call admin"})
		return
	}

	if len(histories) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "History found", "data": histories})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "History not found", "data": nil})
	}
}
