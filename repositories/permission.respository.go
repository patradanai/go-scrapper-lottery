package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"lottery-web-scrapping/models"
	"lottery-web-scrapping/utils"
	"time"
)

type IPermissionRepository interface {
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
	permissionCollection := r.Client.Database(utils.LoadEnv("MONGO_DB_NAME")).Collection("permissions")
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
