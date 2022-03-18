package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Permission struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	MenuID    primitive.ObjectID `bson:"menu_id"`
	Name      string             `bson:"name,omitempty"`
	Action    string             `bson:"action,omitempty"`
	Active    bool               `bson:"active"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
