package indexes

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	Collection       *mongo.Collection
	VectorDimensions int
}

func (r MongoRepo) CreateIndexes(ctx context.Context) ([]string, error) {
	models := vectorSearchIndexes(r.VectorDimensions)
	vectorIndexes, err := r.Collection.SearchIndexes().CreateMany(ctx, models)
	if err != nil {
		return nil, err
	}
	indexes, err := r.Collection.Indexes().CreateMany(ctx, movieIndexes())
	if err != nil {
		return nil, err
	}

	return append(vectorIndexes, indexes...), nil
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

// VectorSearchIndexes returns Atlas Search index models that enable
// nearest neighbour (k-NN) vector search on embeddings.
//
// This uses a wildcard mapping so any embedding stored under
// embeddings.<name>.embedding will be searchable. Pass the vector
// dimensionality of your embedding model (e.g., 1536 for many OpenAI text models).
//
// Note: These are Atlas Search indexes and must be created via the
// SearchIndexes() helper on a collection (not with createIndexes).
func vectorSearchIndexes(dimensions int) []mongo.SearchIndexModel {
	definition := bson.D{
		{Key: "mappings", Value: bson.D{
			{Key: "dynamic", Value: false},
			{Key: "fields", Value: bson.D{
				{Key: "embeddings", Value: bson.D{
					{Key: "fields", Value: bson.D{
						{Key: "*", Value: bson.D{
							{Key: "fields", Value: bson.D{
								{Key: "embedding", Value: bson.D{
									// As of Atlas Search 2024+, "vector" with numDimensions is preferred.
									{Key: "type", Value: "vector"},
									{Key: "numDimensions", Value: dimensions},
									{Key: "similarity", Value: "cosine"},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}

	return []mongo.SearchIndexModel{
		{
			Definition: definition,
		},
	}
}
