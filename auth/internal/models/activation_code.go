package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ActivationCode struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Code       string             `json:"code" bson:"code"`
	ExpireDate time.Time          `json:"expire_date" bson:"expire_date"`
}
