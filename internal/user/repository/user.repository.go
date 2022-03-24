package repository

import (
	"context"
	"lottery-web-scrapping/configs"
	"lottery-web-scrapping/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	CreateOne(user *models.User) error
	FindId(username string) (*models.User, error)
	FindUser(user string) (*models.User, error)
}

type UserRepository struct {
	*mongo.Client
}

func NewUserRepository(c *mongo.Client) IUserRepository {
	return &UserRepository{
		c,
	}
}

func (c *UserRepository) CreateOne(user *models.User) error {
	userCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := userCollection.InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}

func (c *UserRepository) FindId(id string) (*models.User, error) {
	userCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filters := bson.M{"_id": id}
	user := models.User{}
	if err := userCollection.FindOne(ctx, filters).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *UserRepository) FindUser(username string) (*models.User, error) {
	userCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filters := bson.M{"username": username}
	user := models.User{}
	if err := userCollection.FindOne(ctx, filters).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
