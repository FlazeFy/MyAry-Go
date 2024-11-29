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
	}

	// History Module
	historyService := historyServices.NewHistoryService(db)
	historyRepo := historyRepositories.NewHistoryService(historyService)
	historyHandler := historyHandlers.NewHistoryHandler(historyRepo)
	historyGroup := r.Group("/api/v1/history")
	{
		historyGroup.POST("", historyHandler.CreateHistory)
		historyGroup.GET("", historyHandler.GetHistories)
	}

	// Dictionary Module
	dictionaryService := dictionaryServices.NewDictionaryService(db)
	dictionaryRepo := dictionaryRepositories.NewDictionaryService(dictionaryService)
	dictionaryHandler := dictionaryHandlers.NewDictionaryHandler(dictionaryRepo)
	dictionaryGroup := r.Group("/api/v1/dictionary")
	{
		dictionaryGroup.POST("", dictionaryHandler.CreateDictionary)
		dictionaryGroup.GET("", dictionaryHandler.GetDictionaries)
	}
}
