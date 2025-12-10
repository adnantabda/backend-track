package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"task_manager/Domain"
)

// UserRepositoryMongo implements the UserRepository interface for MongoDB
type UserRepositoryMongo struct {
	collection *mongo.Collection
}

// NewUserRepositoryMongo creates a new UserRepositoryMongo instance
func NewUserRepositoryMongo(db *mongo.Database) *UserRepositoryMongo {
	repo := &UserRepositoryMongo{
		collection: db.Collection("users"),
	}

	// Create unique index on email
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := repo.collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		panic(err)
	}

	return repo
}

// Create adds a new user to the database
func (r *UserRepositoryMongo) Create(user domain.User) (domain.User, error) {
	user.ID = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return domain.User{}, domain.ErrEmailAlreadyExists
		}
		return domain.User{}, err
	}

	return user, nil
}

// GetByEmail retrieves a user by their email
func (r *UserRepositoryMongo) GetByEmail(email string) (domain.User, error) {
	var user domain.User

	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return user, nil
}

// GetByID retrieves a user by their ID
func (r *UserRepositoryMongo) GetByID(id string) (domain.User, error) {
	var user domain.User

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = r.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return user, nil
}
