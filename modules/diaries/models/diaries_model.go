package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DiaryModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	DiaryTitle string             `bson:"diary_title" json:"diary_title"`
	DiaryDesc  string             `bson:"diary_desc" json:"diary_desc"`
	DiaryDate  primitive.DateTime `bson:"diary_date" json:"diary_date"`
	DiaryMood  int                `bson:"diary_mood" json:"diary_mood"`
	DiaryTired int                `bson:"diary_tired" json:"diary_tired"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  *time.Time         `bson:"updated_at"`
}

type StatsDiaryLifetimeModel struct {
	Total        int     `bson:"total" json:"total"`
	AverageMood  float32 `bson:"average_mood" json:"average_mood"`
	AverageTired float32 `bson:"average_tired" json:"average_tired"`
}
