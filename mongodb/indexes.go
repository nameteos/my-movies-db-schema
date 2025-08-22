package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoIndexer struct {
	Collection       *mongo.Collection
	VectorDimensions int
}

func (r MongoIndexer) CreateIndexes(ctx context.Context) ([]string, error) {
	indexes, err := r.Collection.Indexes().CreateMany(ctx, movieIndexes())
	if err != nil {
		return nil, err
	}
	return indexes, nil
}

func movieIndexes() []mongo.IndexModel {
	// Text index for basic keyword search on key fields present in our Movie schema.
	textIdx := mongo.IndexModel{
		Keys: bson.D{
			{Key: "title", Value: "text"},
			{Key: "overview", Value: "text"},
			{Key: "genres.name", Value: "text"},
			{Key: "spoken_languages.name", Value: "text"},
		},
		Options: options.Index().SetName("movies_text_idx"),
	}

	// Useful sort/filter indexes.
	createdAtIdx := mongo.IndexModel{
		Keys:    bson.D{{Key: "created_at", Value: -1}},
		Options: options.Index().SetName("created_at_desc_idx"),
	}

	yearIdx := mongo.IndexModel{
		Keys:    bson.D{{Key: "release_date", Value: 1}},
		Options: options.Index().SetName("release_date_idx"),
	}

	return []mongo.IndexModel{textIdx, createdAtIdx, yearIdx}
}
