package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (UserEntity) CollectionName() string {
	return "User"
}

type UserEntity struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty" json:"id" sql:"ignore"`
	IdUser    *int64              `bson:"id_user" json:"id_user" sql:"primary_key"`
	Inserted  *time.Time          `bson:"inserted" json:"inserted"`
	UInserted *int64              `bson:"u_inserted" json:"u_inserted"`
	Updated   *time.Time          `bson:"updated" json:"updated"`
	UUpdated  *int64              `bson:"u_updated" json:"u_updated"`
	Name      *string             `bson:"name" json:"name"`
	Email     *string             `bson:"email" json:"email"`
	Password  *string             `bson:"password" json:"password"`
	Verified  *bool               `bson:"verified" json:"verified"`
}
