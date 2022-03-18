package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OauthClient struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserId       primitive.ObjectID `bson:"user_id"`
	Name         string             `bson:"name,omitempty"`
	ClientId     string             `bson:"client_id,omitempty"`
	ClientSecret string             `bson:"client_secret,omitempty"`
	ApiKey       string             `bson:"api_key,omitempty"`
	Revoke       bool               `bson:"revoke"`
	ExpiredAt    time.Time          `bson:"expired_at,omitempty"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}
