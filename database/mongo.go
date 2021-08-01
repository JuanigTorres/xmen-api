package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/JuanigTorres/xmen-api/model/documents"
)

const (
	DB   = "xmen"
	USER = "xmen-api"
	HOST = "xmendb.5fwxe.mongodb.net"
)

var COLLECTIONS_CONFIG = []func(ctx context.Context){
	createDNAsCollection,
}

var client *mongo.Client

func connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// NewClient Creates a connection to the Mongodb Client
func NewClient() {
	var err error

	pass, exist := os.LookupEnv("MONGO_PASS")

	if !exist {
		log.Fatal("MONGO_PASS variable is empty")
	}

	url := fmt.Sprintf("mongodb+srv://%v:%v@%v/%v?retryWrites=true&w=majority", USER, pass, HOST, DB)

	log.Println("Connecting to ::", url)

	opts := options.Client().ApplyURI(url)
	client, err = mongo.NewClient(opts)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	connect()

	log.Println("Applying configurations ...")
	for _, config := range COLLECTIONS_CONFIG {
		config(ctx)
	}
}

func SaveDNA(document *documents.DNADocument) {
	log.Println("Saving document ::", document)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	insert, err := client.Database(DB).Collection("dnas").InsertOne(ctx, document.AsBson())
	if err != nil {
		log.Println("WARN ::", err)
	}
	log.Println("Saved as ::", insert)
}

func NumberOfDNAs(mutant bool) int64 {
	log.Printf("Searching number of dnas where [mutant = %v]\n", mutant)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count, err := client.Database(DB).Collection("dnas").CountDocuments(ctx, bson.D{{"mutant", mutant}})
	if err != nil {
		log.Println("WARN ::", err)
	}
	log.Println("Founded ::", count)
	return count
}

// createDNAsCollection Create and configure a collection from DNAs into the mongodb.
func createDNAsCollection(ctx context.Context) {
	name := "dnas"
	log.Println("Create collection with name ::", name)

	if e := client.Database(DB).CreateCollection(ctx, name); e != nil {
		log.Println(e)
	}

	model := mongo.IndexModel{
		Keys: bson.M{
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
