package mongodb

import (
	"context"
	pb "github.com/blazee5/cloud-drive/microservices/auth/api/v1"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthStorage struct {
	db *mongo.Collection
}

func NewAuthStorage(db *mongo.Database) *AuthStorage {
	return &AuthStorage{db: db.Collection("users")}
}

func (s *AuthStorage) SignUp(ctx context.Context, input *pb.SignUpRequest) (string, error) {
	res, err := s.db.InsertOne(ctx, auth.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (s *AuthStorage) VerifyUser(ctx context.Context, input *pb.SignInRequest) (auth.User, error) {
	var user auth.User

	filter := bson.D{
		{"email", input.Email},
		{"password", input.Password},
	}

	err := s.db.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return auth.User{}, err
	}

	return user, nil
}
