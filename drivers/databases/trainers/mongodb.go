package trainers

import (
	"context"
	"pokemontrainer/business/trainers"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

//MongodbTrainerRepository ...
type MongodbTrainerRepository struct {
	Conn *mongo.Database
}

// NewMongodbTrainerRepository menghubungkan business ke database
func NewMongodbTrainerRepository(conn *mongo.Database) trainers.MongodbRepository {
	return &MongodbTrainerRepository{
		Conn: conn,
	}
}

// LoginLog add login data to collection trainer_login
func (repo *MongodbTrainerRepository) LoginLog(ctx context.Context, trainerID uint) (trainers.Domain, error) {

	_, err := repo.Conn.Collection("trainer_login").InsertOne(ctx, bson.D{
		{
			Key:   "trainer_id",
			Value: trainerID,
		},
		{
			Key:   "date",
			Value: time.Now(),
		},
	})

	if err != nil {
		return trainers.Domain{}, err
	}

	return trainers.Domain{ID: trainerID}, nil
}
