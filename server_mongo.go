package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServerMongo struct {
	database string
	client   *mongo.Client
}

func NewServerMongo(database string) *ServerMongo {
	return &ServerMongo{
		database: database,
	}
}

func (s *ServerMongo) Connect(connectionString string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return err
	}
	s.client = client
	return nil
}

func (s *ServerMongo) Disconnect() error {
	if s.client != nil {
		return s.client.Disconnect(context.TODO())
	}
	return nil
}

func (s *ServerMongo) GetCollection(collection string, query bson.M) ([]bson.M, error) {
	coll := s.client.Database(s.database).Collection(collection)
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	cursorErr := cursor.All(context.TODO(), &results)
	if cursorErr != nil {
		return nil, cursorErr
	}
	return results, nil
}
