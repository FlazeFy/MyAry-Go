package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeedbackModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FeedbackBody string             `bson:"feedback_body" json:"feedback_body"`
	FeedbackRate int                `bson:"feedback_rate" json:"feedback_rate"`
	CreatedAt    time.Time          `bson:"created_at"`
}
