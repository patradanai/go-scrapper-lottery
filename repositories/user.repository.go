package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"lottery-web-scrapping/models"
	"lottery-web-scrapping/utils"
	"time"
)

type IUserRepository interface {
	FindById(username string) (*models.User, error)
	FindByUser(user string) (*models.User, error)
}

type UserRepository struct {
	BaseRepository
}

func NewUserRepository(c *mongo.Client) IUserRepository {
	return &UserRepository{
		BaseRepository{c},
	}
}

func (c *UserRepository) FindById(id string) (*models.User, error) {
	userCollection := c.Client.Database(utils.LoadEnv("MONGO_DB_NAME")).Collection("drawing_lottery")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filters := bson.M{"_id": id}
	user := models.User{}
	if err := userCollection.FindOne(ctx, filters).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *UserRepository) FindByUser(username string) (*models.User, error) {
	userCollection := c.Client.Database(utils.LoadEnv("MONGO_DB_NAME")).Collection("drawing_lottery")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filters := bson.M{"username": username}
	user := models.User{}
	if err := userCollection.FindOne(ctx, filters).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
