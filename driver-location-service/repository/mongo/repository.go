package repository

import (
	"time"
	"context"
	"os"
	"errors"
	"fmt"

	"github.com/alicevvikk/bitaksi/driver-location-service/domain"
	"github.com/alicevvikk/bitaksi/driver-location-service/logger"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/joho/godotenv"
)

type mongoRepository struct {
	client *mongo.Client
	db	string
	timeout	time.Duration
}

func NewMongoRepository(mongoDbName string, mongoTimeout int) (domain.DriverLocationRepository, error){
	newClient, err := newMongoClient()
	if err != nil {
		return nil, err
	}
	repo := &mongoRepository{
		client:	 newClient,
		db:	 mongoDbName,
		timeout: time.Duration(mongoTimeout) * time.Second,
	}

	return repo, nil
}

func newMongoClient() (*mongo.Client, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("No .env file found: repo.newMongoClient")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		str := "You must set your MONGOB_URI ennivornmental variable. : repo.newMongoClient"
		return nil, errors.New(str)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("mongo.newMongoClient %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("mongo.newMongoClient %w", err)
	}
	return client, nil

}
func (mr *mongoRepository)createLocationIndex() {
	logger.Info("Starting to create location index. from: repo.mongo.createLocationIndex")
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	model := mongo.IndexModel{
		Keys: bson.D{
			{"location", "2dsphere"},
		},
	}

	coll := mr.client.Database(mr.db).Collection("driver-locations")

	var indexView mongo.IndexView
	indexView = coll.Indexes()
	_, err := indexView.CreateOne(
		ctx,
		model,
		nil,
	)

	if err != nil {
		logger.Fatal("Can't create the index. from: repo.mongo.createLocationIndex")
	}
	logger.Info("Index created successfully from: repo.mongo.createLocationInex")
}
func (mr *mongoRepository) ImportInitialData() {
	logger.Info("Starting to import initial data.. from: repo.mongo.ImportInitialData")
	loadCoordinates()
	mr.createLocationIndex()

	coll := mr.client.Database(mr.db).Collection("driver-locations")

	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	_, err := coll.InsertMany(ctx, locations)
	if err != nil {
		logger.Fatal("error. from repo.mongo.ImportInitialData")
	}
	logger.Info("Initial data imporeted successfully.. from: repo.mongo.ImportInitialData")
}

func (mr *mongoRepository) DeleteDriverById(id primitive.ObjectID) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	coll := mr.client.Database(mr.db).Collection("driver-locations")

        filter := bson.M{"_id": id}
	count, err := coll.DeleteOne(
		ctx,
		filter,
		nil,
	)
	if err != nil {
		return 0, err
	}

	return count.DeletedCount, nil
}

func (mr *mongoRepository) CreateDriver(location *domain.DriverLocation) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	coll := mr.client.Database(mr.db).Collection("driver-locations")

	filter := bson.D{{"location", location.Location}}
	_, err := coll.InsertOne(ctx, filter, nil)
	if err != nil {
		return 0, err
	}

	return 1, err
}

func (mr *mongoRepository) UpdateDriver(location *domain.DriverLocation) (int64, error){
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	coll := mr.client.Database(mr.db).Collection("driver-locations")

	filter := bson.D{{"_id", location.Id}}
	update := bson.D{{"$set", bson.D{{"location", location.Location}}}}
	res, err := coll.UpdateOne(ctx, filter, update, nil)

	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

func (mr *mongoRepository) DriverById(id primitive.ObjectID) (*domain.DriverLocation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	coll := mr.client.Database(mr.db).Collection("driver-locations")

	result := new(domain.DriverLocation)
	filter := bson.M{"_id": id}
	err := coll.FindOne(
		ctx,
		filter,
		nil,
	).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func (mr *mongoRepository) DriverByLocation(userLocation *domain.Location, r float64) (*domain.DriverLocation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	collection := mr.client.Database(mr.db).Collection("driver-locations")
	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry":
					userLocation,
					"$maxDistance": r,
			},
		},
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	driverLocation := new(domain.DriverLocation)
	if cursor.Next(ctx) {
		err := cursor.Decode(driverLocation)
		if err != nil {
			return nil, err
		}
	} else {

		return nil, errors.New("No match.")
	}

	return driverLocation, nil
}

func (mr *mongoRepository) Drivers() (domain.DriverLocations, error) {
	var results domain.DriverLocations

	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	collection := mr.client.Database(mr.db).Collection("driver-locations")
	filter := bson.M{}

	cursor, err := collection.Find(
		ctx,
		filter,
		nil,
	)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		result := new(domain.DriverLocation)

		err := cursor.Decode(result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}


