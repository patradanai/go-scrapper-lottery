package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Street       string `bson:"street,omitempty"`
	AddressLine1 string `bson:"address_line1,omitempty"`
	AddressLine2 string `bson:"address_line2,omitempty"`
	City         string `bson:"city,omitempty"`
	Country      string `bson:"country,omitempty"`
	ZipCode      string `bson:"zip_code,omitempty"`
	Building     string `bson:"building,omitempty"`
}

type UserInfo struct {
	FirstName string  `bson:"first_name"`
	LastName  string  `bson:"last_name"`
	Phone     string  `bson:"phone"`
	Address   Address `bson:"address"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Roles     []Role             `bson:"roles"`
	UserInfo  UserInfo           `bson:"user_info"`
	Username  string             `bson:"username,omitempty"`
	Password  string             `bson:"password,,omitempty"`
	Active    bool               `bson:"active"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
