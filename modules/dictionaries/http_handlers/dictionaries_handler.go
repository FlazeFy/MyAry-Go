package httphandlers

import (
	"myary/modules/dictionaries/models"
	"myary/modules/dictionaries/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Interface & Struct
type DictionaryHandler struct {
	repo repositories.DictionaryService
}

func NewDictionaryHandler(repo repositories.DictionaryService) *DictionaryHandler {
	return &DictionaryHandler{repo: repo}
}

// Command Handler
func (h *DictionaryHandler) CreateDictionary(c *gin.Context) {
	var dictionary models.DictionaryModel
	if err := c.ShouldBindJSON(&dictionary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.repo.CreateDictionary(dictionary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong, please call admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Dictionary created"})
}

// Query Handler
func (h *DictionaryHandler) GetDictionaries(c *gin.Context) {
	dictionaries, err := h.repo.FetchDictionaries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong, please call admin"})
		return
	}

	if len(dictionaries) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Dictionary found", "data": dictionaries})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Dictionary not found", "data": nil})

	}
}
