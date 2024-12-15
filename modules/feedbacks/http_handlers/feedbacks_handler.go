package httphandlers

import (
	"myary/modules/feedbacks/models"
	"myary/modules/feedbacks/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Interface & Struct
type FeedbackHandler struct {
	repo repositories.FeedbackRepository
}

func NewFeedbackHandler(repo repositories.FeedbackRepository) *FeedbackHandler {
	return &FeedbackHandler{repo: repo}
}

// Command Handler
func (h *FeedbackHandler) CreateFeedback(c *gin.Context) {
	var feedback models.FeedbackModel
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.repo.CreateFeedback(feedback); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong, please call admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Feedback created"})
}
func (h *FeedbackHandler) DeleteFeedback(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := h.repo.DeleteFeedback(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong, please call admin"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Feedback not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Feedback deleted successfully"})
	}
}

// Query Handler
func (h *FeedbackHandler) GetFeedBack(c *gin.Context) {
	feedbacks, err := h.repo.FetchFeedbacks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong, please call admin"})
		return
	}

	if len(feedbacks) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Feedback found", "data": feedbacks})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Feedback not found", "data": nil})
	}
}
func (h *FeedbackHandler) GetFeedBackStats(c *gin.Context) {
	feedback, err := h.repo.FetchFeedBackStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong, please call admin"})
		return
	}

	if len(feedback) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Feedback not found", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feedback found", "data": feedback})
}
