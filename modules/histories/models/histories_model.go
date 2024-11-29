package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HistoryModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	HistoryContext string             `bson:"history_context" json:"history_context"`
	HistoryType    string             `bson:"history_type" json:"history_type"`
	CreatedBy      string             `bson:"created_by" json:"created_by"`
}
