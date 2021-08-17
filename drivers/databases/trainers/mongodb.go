package trainers

import (
	"context"
	"fmt"
	"pokemontrainer/business/trainers"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"

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

	opts := options.Find()
	// yesterdayDate := time.Now().AddDate(0, 0, -1)
	opts.SetSort(bson.D{
		primitive.E{
			Key:   "date",
			Value: -1,
		},
	})

	fmt.Println("trainerID", trainerID)
	//get latest login data
	// sortCursor, findErr := repo.Conn.Collection("trainer_login").Find(ctx, bson.D{
	// 	primitive.E{
	// 		Key: "date",
	// 		Value: bson.D{
	// 			primitive.E{
	// 				Key:   "$lt",
	// 				Value: time.Now(),
	// 			},
	// 		},
	// 	},
	// 	primitive.E{
	// 		Key:   "trainer_id",
	// 		Value: trainerID,
	// 	},
	// }, opts)

	// if findErr != nil {
	// 	return trainers.Domain{}, findErr
	// }
	// var logList []TrainerLogin
	// if findErr = sortCursor.All(ctx, &logList); findErr != nil {
	// 	return trainers.Domain{}, findErr
	// }

	// fmt.Println("loginData", logList)
	// fmt.Println("yesterdayDate", yesterdayDate)
	// fmt.Println("range of login", logList[0].Date.Day()-yesterdayDate.Day())

	var insertErr error

	// jika sudah beda hari, maka akan dilog
	// if len(logList) == 0 || logList[0].Date.Day()-yesterdayDate.Day() < 1 {

	fmt.Println("insert data")
	_, insertErr = repo.Conn.Collection("trainer_login").InsertOne(ctx, bson.D{
		{
			Key:   "trainer_id",
			Value: trainerID,
		},
		{
			Key:   "date",
			Value: time.Now(),
		},
	})
	// } else {
	// 	fmt.Println("not insert data")
	// }

	if insertErr != nil {
		return trainers.Domain{}, insertErr
	}

	return trainers.Domain{ID: trainerID}, nil
}
