package repository

import (
	"context"
	"lottery-web-scrapping/configs"
	"lottery-web-scrapping/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRoleRepository interface {
	FindByName(name string) (*models.Role, error)
}

type RoleRepository struct {
	*mongo.Client
}

func NewRoleRepository(c *mongo.Client) IRoleRepository {
	return &RoleRepository{
		c,
	}
}

func (r *RoleRepository) FindByName(name string) (*models.Role, error) {
	roleCollection := r.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("roles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	role := models.Role{}
	filters := bson.M{"name": name}

	if err := roleCollection.FindOne(ctx, filters).Decode(&role); err != nil {
		return nil, err
	}

	return &role, nil
}
