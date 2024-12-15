package routes

import (
	diaryHandlers "myary/modules/diaries/http_handlers"
	diaryRepositories "myary/modules/diaries/repositories"
	diaryServices "myary/modules/diaries/services"

	historyHandlers "myary/modules/histories/http_handlers"
	historyRepositories "myary/modules/histories/repositories"
	historyServices "myary/modules/histories/services"

	dictionaryHandlers "myary/modules/dictionaries/http_handlers"
	dictionaryRepositories "myary/modules/dictionaries/repositories"
	dictionaryServices "myary/modules/dictionaries/services"

	feedbackHandlers "myary/modules/feedbacks/http_handlers"
	feedbackRepositories "myary/modules/feedbacks/repositories"
	feedbackServices "myary/modules/feedbacks/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(r *gin.Engine, db *mongo.Database) {
	// Diary Module
	diaryService := diaryServices.NewDiaryService(db)
	diaryRepo := diaryRepositories.NewDiaryService(diaryService)
	diaryHandler := diaryHandlers.NewDiaryHandler(diaryRepo)
	diaryGroup := r.Group("/api/v1/diary")
	{
		diaryGroup.POST("", diaryHandler.CreateDiary)
		diaryGroup.GET("", diaryHandler.GetDiaries)
		diaryGroup.GET("/:id", diaryHandler.GetDiaryById)
		diaryGroup.GET("/stats/lifetime", diaryHandler.GetDiaryStatsLifetime)
		diaryGroup.PUT("/:id", diaryHandler.UpdateDiary)
		diaryGroup.DELETE("/:id", diaryHandler.DeleteDiary)
	}

	// History Module
	historyService := historyServices.NewHistoryService(db)
	historyRepo := historyRepositories.NewHistoryService(historyService)
	historyHandler := historyHandlers.NewHistoryHandler(historyRepo)
	historyGroup := r.Group("/api/v1/history")
	{
		historyGroup.POST("", historyHandler.CreateHistory)
		historyGroup.GET("", historyHandler.GetHistories)
		historyGroup.DELETE("/:id", historyHandler.DeleteHistory)
	}

	// Dictionary Module
	dictionaryService := dictionaryServices.NewDictionaryService(db)
	dictionaryRepo := dictionaryRepositories.NewDictionaryService(dictionaryService)
	dictionaryHandler := dictionaryHandlers.NewDictionaryHandler(dictionaryRepo)
	dictionaryGroup := r.Group("/api/v1/dictionary")
	{
		dictionaryGroup.POST("", dictionaryHandler.CreateDictionary)
		dictionaryGroup.GET("", dictionaryHandler.GetDictionaries)
		dictionaryGroup.GET("/stats", dictionaryHandler.GetTotalDictionaryUsed)
		dictionaryGroup.DELETE("/:id", dictionaryHandler.DeleteDictionary)
	}

	// Feedback Module
	feedbackService := feedbackServices.NewFeedbackService(db)
	feedbackRepo := feedbackRepositories.NewFeedbackService(feedbackService)
	feedbackHandler := feedbackHandlers.NewFeedbackHandler(feedbackRepo)
	feedbackGroup := r.Group("/api/v1/feedback")
	{
		feedbackGroup.POST("", feedbackHandler.CreateFeedback)
		feedbackGroup.GET("", feedbackHandler.GetFeedBack)
		feedbackGroup.GET("/stats", feedbackHandler.GetFeedBackStats)
		feedbackGroup.DELETE("/:id", feedbackHandler.DeleteFeedback)
	}
}
