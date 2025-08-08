package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// EmbeddingObject represents an OpenAI embedding object
// Based on https://platform.openai.com/docs/api-reference/embeddings/object
type EmbeddingObject struct {
	Object    string    `json:"object" bson:"object"`       // Always "embedding"
	Embedding []float32 `json:"embedding" bson:"embedding"` // Vector representation
	Index     int       `json:"index" bson:"index"`         // Index of the embedding
	Model     string    `json:"model" bson:"model"`         // Model used to generate the embedding
	Usage     struct {
		PromptTokens int `json:"prompt_tokens" bson:"prompt_tokens"` // Number of tokens in the prompt
		TotalTokens  int `json:"total_tokens" bson:"total_tokens"`   // Total number of tokens used
	} `json:"usage,omitempty" bson:"usage,omitempty"` // Usage information
}

type Movie struct {
	ID               primitive.ObjectID         `bson:"_id,omitempty" json:"id"`
	Embeddings       map[string]EmbeddingObject `bson:"embeddings,omitempty" json:"embeddings"`
	Adult            bool                       `json:"adult"`
	Budget           int64                      `json:"budget"`
	Genres           []TMDBGenre                `json:"genres"`
	ExternalID       int                        `json:"external_id"`
	IMDbID           string                     `json:"imdb_id"`
	OriginalLanguage string                     `json:"original_language"`
	OriginalTitle    string                     `json:"original_title"`
	Overview         string                     `json:"overview"`
	PosterPath       string                     `json:"poster_path"`
	ReleaseDate      string                     `json:"release_date"`
	Revenue          int64                      `json:"revenue"`
	Runtime          int                        `json:"runtime"`
	SpokenLanguages  []TMDBLanguage             `json:"spoken_languages"`
	Status           string                     `json:"status"`
	Tagline          string                     `json:"tagline"`
	Title            string                     `json:"title"`
	CreatedAt        time.Time                  `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time                  `bson:"updated_at" json:"updated_at"`
}

// TMDBMovie represents the response from TMDB movie detail API
// https://developer.themoviedb.org/reference/movie-detai
type TMDBMovie struct {
	Adult            bool           `json:"adult"`
	Budget           int64          `json:"budget"`
	Genres           []TMDBGenre    `json:"genres"`
	ID               int            `json:"id"`
	IMDbID           string         `json:"imdb_id"`
	OriginalLanguage string         `json:"original_language"`
	OriginalTitle    string         `json:"original_title"`
	Overview         string         `json:"overview"`
	PosterPath       string         `json:"poster_path"`
	ReleaseDate      string         `json:"release_date"`
	Revenue          int64          `json:"revenue"`
	Runtime          int            `json:"runtime"`
	SpokenLanguages  []TMDBLanguage `json:"spoken_languages"`
	Status           string         `json:"status"`
	Tagline          string         `json:"tagline"`
	Title            string         `json:"title"`
}

type TMDBGenre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TMDBLanguage struct {
	ISO6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name"`
}
