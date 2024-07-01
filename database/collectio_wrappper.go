package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type CollectionWrapper struct {
	collection *mongo.Collection
}

func NewCollectionWrapper(dbName, collectionName string) *CollectionWrapper {
	client := GetClient()
	collection := client.Database(dbName).Collection(collectionName)
	return &CollectionWrapper{collection: collection}
}

func (c *CollectionWrapper) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	result, err := c.collection.InsertOne(ctx, document)
	if err != nil {
		log.Printf("InsertOne failed: %v", err)
		return nil, err
	}
	return result, nil
}

func (c *CollectionWrapper) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	result := c.collection.FindOne(ctx, filter)
	return result
}

func (c *CollectionWrapper) UpdateOne(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error) {
	result, err := c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("UpdateOne failed: %v", err)
		return nil, err
	}
	return result, nil
}

func (c *CollectionWrapper) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	result, err := c.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("DeleteOne failed: %v", err)
		return nil, err
	}
	return result, nil
}
