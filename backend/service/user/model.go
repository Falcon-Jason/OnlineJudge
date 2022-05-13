package user

import (
	"OnlineJudge_Backend/dao"
	"OnlineJudge_Backend/util"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var users *mongo.Collection

type user struct {
	Id       string `json:"user_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

func register(ctx context.Context, u *user) error {
	_, err := users.InsertOne(ctx, u)

	// handle write exception
	if e, ok := err.(mongo.WriteException); ok && e.HasErrorCode(11000) {
		return errors.New("username has been used")
	}

	return err
}

type loginReply struct {
	Id         string `json:"user_id"`
	Username   string `json:"username"`
	IsTeacher  bool   `json:"is_teacher"`
	LoginToken string `json:"login_token"`
}

func login(ctx context.Context, u *user) (*loginReply, error) {
	result := user{}
	if u.Username == "" || u.Password == "" {
		return nil, util.ErrInvalidRequest
	}

	// check username and password
	err := users.FindOne(ctx, u).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("wrong username or password")
	} else if err != nil {
		return nil, err
	}

	// generate token
	token := primitive.NewObjectID()
	key := fmt.Sprintf("user_id:%s", token.Hex())
	value := result.Id

	// set token to redis
	cmd := dao.Redis.Set(ctx, key, value, 30*time.Minute)
	if err = cmd.Err(); err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(result.Id)
	if err != nil {
		return nil, err
	}
	auth, err := Authorization(ctx, &AuthRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &loginReply{
		Id:         result.Id,
		LoginToken: token.Hex(),
		IsTeacher:  auth.IsTeacher || auth.IsAdmin,
		Username:   result.Username,
	}, nil
}

func queryUserIdByToken(ctx context.Context, token string) (string, error) {
	key := fmt.Sprintf("user_id:%s", token)
	data, err := dao.Redis.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", errors.New("token is invalid or expired")
	} else if err != nil {
		return "", err
	}

	return data, nil
}

func modifyUserInfo(ctx context.Context, u *user) error {
	id, err := primitive.ObjectIDFromHex(u.Id)
	if err != nil {
		return util.ErrInvalidRequest
	}

	u.Id = ""

	data, err := users.UpdateOne(
		ctx,
		bson.D{{"_id", id}},
		bson.D{{"$set", u}})

	if data.ModifiedCount == 0 {
		return util.ErrNothingModified
	}

	return err
}

type AuthRequest struct {
	Id primitive.ObjectID `bson:"_id"`
}

type AuthReply struct {
	Id        string `bson:"_id"`
	IsTeacher bool   `bson:"is_teacher"`
	IsAdmin   bool   `bson:"is_admin"`
}

func Authorization(ctx context.Context, r *AuthRequest) (*AuthReply, error) {
	var reply AuthReply
	err := users.FindOne(ctx, r).Decode(&reply)
	if err != nil {
		return nil, err
	}

	return &reply, nil
}

func initModel() {
	users = dao.Mongo.Collection("user")
}
