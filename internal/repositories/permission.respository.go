package repositories

import (
	"context"
	"lottery-web-scrapping/configs"
	"lottery-web-scrapping/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPermissionRepository interface {
	FindAllPermission() ([]models.Permission, error)
}

type PermissionRepository struct {
	BaseRepository
}

func NewPermissionRepository(c *mongo.Client) IPermissionRepository {
	return &PermissionRepository{
		BaseRepository{c},
	}
}

func (r *BaseRepository) FindAllPermission() ([]models.Permission, error) {
	permissionCollection := r.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("permissions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filters := bson.M{}

	cur, err := permissionCollection.Find(ctx, filters)
	if err != nil {
		return nil, err
	}

	permissions := make([]models.Permission, 0)

	for cur.Next(ctx) {
		permission := models.Permission{}

		if err := cur.Decode(&permission); err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	return permissions, nil
}

func (r *BaseRepository) FindOnePermission(filter string) (*models.Permission, error) {
	permissionCollection := r.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("permissions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	permission := models.Permission{}

	if err := permissionCollection.FindOne(ctx, filter).Decode(&permission); err != nil {
		return nil, err
	}

	return &permission, nil
}
