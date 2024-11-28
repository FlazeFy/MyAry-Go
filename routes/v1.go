package routes

import (
	"myary/modules/diaries/http_handlers"
	"myary/modules/diaries/repositories"
	"myary/modules/diaries/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(r *gin.Engine, db *mongo.Database) {
	service := services.NewDiaryService(db)
	repo := repositories.NewDiaryService(service)
	handler := http_handlers.NewDiaryHandler(repo)

	r.POST("/api/v1/diary", handler.CreateDiary)
	r.GET("/api/v1/diary", handler.GetDiaries)
}
