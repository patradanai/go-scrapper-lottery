package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Permissions []Permission       `bson:"permissions"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Active      string             `bson:"active,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
