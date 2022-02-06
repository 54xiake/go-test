package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type User struct {
	Name string
	Age  int
	City string
}

var client *mongo.Client

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin@127.0.0.1:27017")
	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	//create()
	//batchCreate()
	//modify()
	//getDetail()
	//getList()
	//remove()
	//createIndex()
	updateUnset()
}

func create()  {
	collection := client.Database("test").Collection("users")
	ash := User{"yg", 33, "BaoDing"}
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func batchCreate()  {
	users := []interface{}{}
	collection := client.Database("test").Collection("users")
	for i:=0; i<10000; i++ {
		users = append(users, User{"yg", 33, "BaoDing"})
	}
	insertManyResult, err := collection.InsertMany(context.TODO(), users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

}

func modify()  {
	collection := client.Database("test").Collection("users")
	filter := bson.D{{"name", "yg"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(updateResult)
}

func getDetail()  {
	collection := client.Database("test").Collection("users")
	filter := bson.D{{"age", 34}}
	var result User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
}

func getList()  {
	collection := client.Database("test").Collection("users")
	one, _:= collection.Find(context.TODO(), bson.M{"name": "yg"})

	defer func() {  // 关闭
		if err := one.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	c := []User{}
	_ = one.All(context.TODO(), &c)   // 当然也可以用   next
	for _, r := range c{
		log.Println(r)
	}
}

func remove()  {
	collection := client.Database("test").Collection("users")
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{"name", "yg"}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(deleteResult)
}

func batchRemove()  {
	collection := client.Database("test").Collection("users")
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{"name", "yg"}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(deleteResult)

}

func createIndex()  {
	mod := mongo.IndexModel{
		Keys: bson.M{
			"Some Int": -1, // index in descending order
		},
		// create UniqueIndex option
		Options: options.Index().SetUnique(true),
	}

	// Create an Index using the CreateOne() method
	collection := client.Database("test").Collection("users")
	ind, err := collection.Indexes().CreateOne(context.TODO(), mod)  // returns (string, error)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ind)
}

func updateUnset()  {
	collection := client.Database("test").Collection("users")
	filter := bson.D{{"name", "yg"}}

	update := bson.D{
		{"$unset", bson.D{
			{"age", 100},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(updateResult)
}

func close()  {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}