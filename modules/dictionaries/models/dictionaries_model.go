package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DictionaryModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	DictionaryType string             `bson:"dictionary_type" json:"dictionary_type"`
	DictionaryName string             `bson:"dictionary_name" json:"dictionary_name"`
}
