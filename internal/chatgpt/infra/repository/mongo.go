package repository

import (
	"context"
	"github.com/guil95/openai-lets-go/internal/chatgpt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"os"
)

const collectionChat = "chats"

type repository struct {
	db *mongo.Client
}

func NewCommandMongoRepository(db *mongo.Client) chatgpt.Repository {
	return repository{db}
}

func (r repository) GetChatByEntityID(ctx context.Context, entityID string) (*chatgpt.ChatGPT, error) {
	db := r.db.Database(os.Getenv("DB_DATABASE"))

	var chat chatgpt.ChatGPT

	err := db.Collection(collectionChat).FindOne(
		ctx,
		bson.D{{"entity_id", entityID}},
	).Decode(&chat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &chat, nil
}

func (r repository) SaveChat(ctx context.Context, gpt chatgpt.ChatGPT) error {
	db := r.db.Database(os.Getenv("DB_DATABASE"))

	_, err := db.Collection(collectionChat).InsertOne(ctx, gpt)
	if err != nil {
		zap.S().Error(err)
		return err
	}

	return nil
}

func (r repository) UpdateChat(ctx context.Context, gpt chatgpt.ChatGPT) error {
	db := r.db.Database(os.Getenv("DB_DATABASE"))

	updateFields := bson.M{"$set": bson.M{
		"text": gpt.Text,
	}}

	_, err := db.Collection(collectionChat).UpdateOne(ctx, bson.D{{"entity_id", gpt.EntityID}}, updateFields)
	if err != nil {
		zap.S().Error(err)
		return err
	}

	return nil
}
