package repository

import (
	"time"
	"context"
	"os"
	"errors"
	"log"

	"github.com/alicevvikk/bitaksi/driver-location-service/domain"
	"github.com/alicevvikk/bitaksi/driver-location-service/utils"

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
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client, nil

}

func (mr *mongoRepository) DeleteDriverById(id string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	coll := mr.client.Database(mr.db).Collection("driver-locations")

        objId, err := primitive.ObjectIDFromHex(id)
        if err != nil {
                return 0, err
        }

	filter := bson.M{"_id": objId}
	count, err := coll.DeleteOne(
		ctx,
		filter,
		nil,
	)

	return count.DeletedCount, nil
}

func (mr *mongoRepository) CreateDriver(locations domain.Locations) (count int) {

	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	coll := mr.client.Database(mr.db).Collection("driver-locations")

	for _, location := range locations {
		if location.Id.IsZero() == true {
			filter := bson.D{{"location", location.Location}}
			_, err := coll.InsertOne(ctx, filter, nil)
			if err != nil {
				return count
			}

			log.Println("Creating..")
			if err == nil {
				count ++
			}

		} else {

			filter := bson.D{{"_id", location.Id}}
			update := bson.D{{"$set", bson.D{{"location", location.Location}}}}
			_, err := coll.UpdateOne(ctx, filter, update, nil)
			if err != nil {
				log.Println("UPDATE error: ", err)
			} else {
			count ++
		}	}
	}

	log.Println("Atleast", count, "documents have been changed or created..")
	return count
}

func (mr *mongoRepository) DriverById(id string) (domain.DriverLocation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	coll := mr.client.Database(mr.db).Collection("driver-locations")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.DriverLocation{}, err
	}

	var result domain.DriverLocation

	filter := bson.M{"_id": objID}
	err = coll.FindOne(
		ctx,
		filter,
		nil,
	).Decode(&result)

	if err != nil {
		return domain.DriverLocation{}, err
	}

	return result, nil
}


func (mr *mongoRepository) DriverByLocation(userLocation *domain.Location) (*domain.ResponseLocation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	collection := mr.client.Database(mr.db).Collection("driver-locations")

	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry":
					userLocation,
					"$maxDistance": 3000,
			},
		},
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	driverLocation := domain.DriverLocation{}
	if cursor.Next(ctx) {
		err := cursor.Decode(&driverLocation)
		if err != nil {
			return nil, err
		}
	} else {

		log.Println("HERE1: ", err)
		return nil, errors.New("No match.")
	}

	log.Println("HERE: ", err)
	distance := utils.CalculateDistance(
		userLocation.Coordinates,
		driverLocation.Location.Coordinates)

	model := &domain.ResponseLocation{
		DriverLocation:	driverLocation,
		Distance:	distance,
	}
	return model, nil
}

func (mr *mongoRepository) Drivers() ([]domain.DriverLocation, error) {
	var results []domain.DriverLocation

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
		var result domain.DriverLocation

		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
