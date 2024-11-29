package routes

import (
	diaryHandlers "myary/modules/diaries/http_handlers"
	diaryRepositories "myary/modules/diaries/repositories"
	diaryServices "myary/modules/diaries/services"

	historyHandlers "myary/modules/histories/http_handlers"
	historyRepositories "myary/modules/histories/repositories"
	historyServices "myary/modules/histories/services"

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
}
