package httphandlers

import (
	"myary/modules/histories/models"
	"myary/modules/histories/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Interface & Struct
type HistoryHandler struct {
	repo repositories.HistoryService
}

func NewHistoryHandler(repo repositories.HistoryService) *HistoryHandler {
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
