package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/JuanigTorres/xmen-finder/model/documents"
)

const DB  = "xmen"
const URI = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000"

var COLLECTIONS_CONFIG = []func(ctx context.Context) {
	createDNAsCollection,
}

var client *mongo.Client

func connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
	defer cancel()
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
	defer cancel()
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func NewClient() {
	var err error

	opts := options.Client().ApplyURI(URI)
	client, err = mongo.NewClient(opts)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
	defer cancel()

	connect()

	log.Println("Applying configurations ...")
	for _, config := range COLLECTIONS_CONFIG {
		config(ctx)
	}
}

func SaveDNA(document *documents.DNADocument) {
	log.Println("Saving document ::", document)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	insert, err := client.Database(DB).Collection("dnas").InsertOne(ctx, document.AsBson())
	if err != nil {
		log.Println("WARN ::", err)
	}
	fmt.Println(insert)
}

func Stats() {

}

func createDNAsCollection(ctx context.Context) {
	name := "dnas"
	log.Println("Create collection with name ::", name)

	if e := client.Database(DB).CreateCollection(ctx, name); e != nil {
		log.Println(e)
	}

	model := mongo.IndexModel{
		Keys: bson.M {
			"dna": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	constraint, err := client.Database(DB).Collection(name).Indexes().CreateOne(ctx, model)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Added unique constrain ::", constraint)
}