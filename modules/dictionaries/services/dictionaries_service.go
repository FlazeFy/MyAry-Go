package services

import (
	"context"
	"errors"
	"myary/modules/dictionaries/models"

	"cloud.google.com/go/firestore"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Interface & Struct
type DictionaryService interface {
	Insert(dictionary models.DictionaryModel) (interface{}, string, error)
	GetAll() ([]models.DictionaryModel, error)
	Delete(filter bson.M) (*mongo.DeleteResult, error)
}
type dictionaryService struct {
	collection *mongo.Collection
}
type firestoreDictionaryService struct {
	collection *firestore.CollectionRef
}

func NewDictionaryService(db *mongo.Database) DictionaryService {
	return &dictionaryService{
		collection: db.Collection("dictionaries"),
	}
}
func NewFirestoreDictionaryService(firestoreClient *firestore.Client) DictionaryService {
	return &firestoreDictionaryService{
		collection: firestoreClient.Collection("dictionaries"),
	}
}

// Command Service
func (r *dictionaryService) Insert(dictionary models.DictionaryModel) (interface{}, string, error) {
	result, err := r.collection.InsertOne(context.TODO(), dictionary)
	if err != nil {
		return nil, "", err
	}
	return result, result.InsertedID.(string), nil
}
func (s *firestoreDictionaryService) Insert(dictionary models.DictionaryModel) (interface{}, string, error) {
	docRef, writeResult, err := s.collection.Add(context.TODO(), dictionary)
	if err != nil {
		return nil, "", err
	}
	return writeResult, docRef.ID, nil
}
func (s *dictionaryService) Delete(filter bson.M) (*mongo.DeleteResult, error) {
	return s.collection.DeleteOne(context.TODO(), filter)
}
func (s *firestoreDictionaryService) Delete(filter bson.M) (*mongo.DeleteResult, error) {
	documentID, ok := filter["_id"].(string)
	if !ok {
		return nil, errors.New("invalid filter: _id must be a string")
	}

	docRef := s.collection.Doc(documentID)

	_, err := docRef.Delete(context.TODO())
	if err != nil {
		return nil, err
	}

	deleteResult := &mongo.DeleteResult{
		DeletedCount: 1,
	}

	return deleteResult, nil
}

// Query Service
func (r *dictionaryService) GetAll() ([]models.DictionaryModel, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}

	var dictionaries []models.DictionaryModel
	if err := cursor.All(context.TODO(), &dictionaries); err != nil {
		return nil, err
	}

	return dictionaries, nil
}

func (s *firestoreDictionaryService) GetAll() ([]models.DictionaryModel, error) {
	iter := s.collection.Documents(context.TODO())
	var dictionaries []models.DictionaryModel

	for {
		doc, err := iter.Next()
		if err != nil {
			if err.Error() == "iterator: no more documents" {
				break
			}
			return nil, err
		}
		var dictionary models.DictionaryModel
		if err := doc.DataTo(&dictionary); err != nil {
			return nil, err
		}
		dictionaries = append(dictionaries, dictionary)
	}

	return dictionaries, nil
}
