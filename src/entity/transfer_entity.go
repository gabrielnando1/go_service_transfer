package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (TransferEntity) CollectionName() string {
	return "Transfer"
}

type TransferEntity struct {
	ID             *primitive.ObjectID `bson:"_id,omitempty" json:"id" sql:"ignore"`
	ExternalId     *string             `bson:"external_id" json:"external_id"`
	InternalId     *string             `bson:"internal_id" json:"internal_id"`
	Inserted       *time.Time          `bson:"inserted" json:"inserted"`
	Updated        *time.Time          `bson:"updated" json:"updated"`
	AccountFrom    *string             `bson:"account_from" json:"account_from"`
	AccountTarget  *string             `bson:"account_target" json:"account_target"`
	Amount         *int64              `bson:"amount" json:"amount"`
	ExpirationDate *time.Time          `bson:"expiration_date" json:"expiration_date"`
	ExpectedOn     *time.Time          `bson:"expected_on" json:"expected_on"`
	Status         *string             `bson:"status" json:"status"`
}
