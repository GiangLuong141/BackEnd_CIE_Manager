package responses

import (
	"context"
	model "example/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepository interface {
	GetUserByEmail(email string) (*model.Admin, error)
}

type mongoUserRepository struct {
	client         *mongo.Client
	collectionName string
}

// GetUserByEmail implements AdminRepository.
func (m *mongoUserRepository) GetUserByEmail(email string) (*model.Admin, error) {
	// Get the collection
	collection := m.client.Database("test").Collection(m.collectionName)

	// Create a filter to find the user by email
	filter := bson.M{"email": email}

	// Fetch the user from the database
	var admin model.Admin
	err := collection.FindOne(context.Background(), filter).Decode(&admin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &admin, nil

}

func NewAdminRepository(client *mongo.Client, collectionName string) AdminRepository {
	return &mongoUserRepository{
		client:         client,
		collectionName: collectionName,
	}
}
