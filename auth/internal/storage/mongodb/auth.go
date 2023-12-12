package mongodb

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/auth"
	"github.com/blazee5/cloud-drive/auth/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AuthStorage struct {
	db *mongo.Database
}

func NewAuthStorage(db *mongo.Database) *AuthStorage {
	return &AuthStorage{db: db}
}

func (s *AuthStorage) SignUp(ctx context.Context, input *pb.SignUpRequest, code string) (string, error) {
	res, err := s.db.Collection("users").InsertOne(ctx, models.User{
		Username:    input.Username,
		Email:       input.Email,
		Password:    input.Password,
		IsActivated: false,
	})

	if err != nil {
		return "", err
	}

	userID := res.InsertedID.(primitive.ObjectID)

	if err != nil {
		return "", err
	}

	_, err = s.db.Collection("activation_codes").InsertOne(ctx, models.ActivationCode{
		Code:       code,
		UserID:     userID,
		ExpireDate: time.Now().Add(time.Hour * 2),
	})

	if err != nil {
		return "", err
	}

	return userID.Hex(), nil
}

func (s *AuthStorage) VerifyUser(ctx context.Context, input *pb.SignInRequest) (models.User, error) {
	var user models.User

	filter := bson.D{
		{"email", input.Email},
		{"password", input.Password},
	}

	err := s.db.Collection("users").FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *AuthStorage) GetActivationCode(ctx context.Context, code string) (models.ActivationCode, error) {
	var activationCode models.ActivationCode

	err := s.db.Collection("activation_codes").FindOne(ctx, bson.D{
		{"code", code},
	}).Decode(&activationCode)

	if err != nil {
		return models.ActivationCode{}, err
	}

	return activationCode, nil
}

func (s *AuthStorage) ActivateUser(ctx context.Context, userID string) error {
	objectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return err
	}

	err = s.db.Collection("users").FindOneAndUpdate(ctx, bson.D{{"_id", objectID}}, bson.D{
		{"$set", bson.D{
			{"is_activated", true},
		}},
	}).Err()

	if err != nil {
		return err
	}

	return nil
}

func (s *AuthStorage) DeleteActivationCode(ctx context.Context, ID string) error {
	objectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}

	_, err = s.db.Collection("activation_codes").DeleteOne(ctx, bson.D{{"_id", objectID}})

	if err != nil {
		return err
	}

	return nil
}
