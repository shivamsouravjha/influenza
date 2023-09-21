package mongo

import (
	"context"
	"fmt"
	"sync"

	"github.com/shivamsouravjha/influenza/config"
	"github.com/shivamsouravjha/influenza/structure"
	requestStruct "github.com/shivamsouravjha/influenza/structure/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	once   sync.Once
)

func InitDb() error {
	// Use sync.Once to ensure the initialization code runs only once.
	once.Do(func() {
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		mongoURI := "mongodb://" + config.Get().MongoUsername + ":" + config.Get().MongoPassword + "@127.0.0.1:27017"

		opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
		// Create a new client and connect to the server
		var err error
		client, err = mongo.Connect(context.TODO(), opts)
		if err != nil {
			panic(err)
		}

		// Send a ping to confirm a successful connection
		pingCmd := bson.M{"ping": 1}
		if err := client.Database("admin").RunCommand(context.TODO(), pingCmd).Err(); err != nil {
			panic(err)
		}

		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

		if err := createDatabaseAndIndex(); err != nil {
			fmt.Println(err.Error())
		}
	})

	return nil
}

func Client() *mongo.Client {
	return client
}

func createDatabaseAndIndex() error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"linkedin": 1},
		Options: options.Index().SetUnique(true),
	}

	collection := client.Database("influenza").Collection("inlfuenza")
	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}

func CreateItem(item requestStruct.FeedbackData, collection string) error {
	_, err := client.Database("influenza").Collection(collection).InsertOne(context.TODO(), item)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func Find(filter bson.M, collection string) (*structure.Influencer, error) {
	var existingUser structure.Influencer
	dbCollection := client.Database("influenza").Collection(collection)
	err := dbCollection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No document found for the given filter
			fmt.Printf("No document found for filter: %+v\n", filter)
		} else {
			// Other error occurred
			fmt.Printf("Error while querying MongoDB: %v\n", err)
		}
		fmt.Println(filter)
		fmt.Println(err.Error())
		return nil, err
	}
	return &existingUser, nil
}

func CreateInfluenza(item structure.Influencer, collection string) (*mongo.InsertOneResult, error) {
	res, err := client.Database("influenza").Collection(collection).InsertOne(context.TODO(), item)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return res, nil
}

func GroupFind(filter bson.M, collection string) ([]*structure.Influencer, error) {
	var existingUsers []*structure.Influencer
	dbCollection := client.Database("influenza").Collection(collection)
	// projection := bson.M{
	// 	"name":     1,
	// 	"linkedin": 1,
	// }
	fmt.Println("dfdsfdsfsdfdsfds")
	options := options.Find()
	cursor, err := dbCollection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		fmt.Println(filter, "here")
		fmt.Println(err.Error())
		return nil, err
	}
	// fmt.Println(cursor)
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user structure.Influencer
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println("Error decoding document:", err)
			continue
		}
		fmt.Println("even here")
		existingUsers = append(existingUsers, &user)
	}

	if err := cursor.Err(); err != nil {
		fmt.Println("Error iterating over cursor:", err)
		return nil, err
	}

	return existingUsers, nil
}

func Delete(filter bson.M, collection string) (*structure.Influencer, error) {
	var existingUser structure.Influencer
	dbCollection := client.Database("influenza").Collection(collection)
	_, err := dbCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No document found for the given filter
			fmt.Printf("No document found for filter: %+v\n", filter)
		} else {
			// Other error occurred
			fmt.Printf("Error while querying MongoDB: %v\n", err)
		}
		fmt.Println(filter)
		fmt.Println(err.Error())
		return nil, err
	}
	return &existingUser, nil
}
