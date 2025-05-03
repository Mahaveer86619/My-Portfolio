package resume

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

// ConnectDB establishes a connection to MongoDB.
func ConnectDB(ctx context.Context) (*Mongo, error) {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://root:example@localhost:27017/"
		log.Println("MONGO_URI environment variable not set, using default URI:", mongoURI)
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	log.Println("Connected to MongoDB!")

	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		dbName = "portfolio"
		log.Println("MONGO_DB_NAME environment variable not set, using default database name:", dbName)
	}

	collectionName := os.Getenv("MONGO_COLLECTION_NAME")
	if collectionName == "" {
		collectionName = "resumes"
		log.Println("MONGO_COLLECTION_NAME environment variable not set, using default collection name:", collectionName)
	}

	mongo := &Mongo{
		Client:     client,
		Database:   client.Database(dbName),
		Collection: client.Database(dbName).Collection(collectionName),
	}
	return mongo, nil
}

// DisconnectDB closes the MongoDB connection.
func DisconnectDB(ctx context.Context, mongo *Mongo) {
	if mongo == nil {
		log.Println("No MongoDB connection to close.")
		return
	}

	if err := mongo.Client.Disconnect(ctx); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
	log.Println("Disconnected from MongoDB!")
}

// GetResume retrieves resume data from MongoDB.
func (m *Mongo) GetResume(ctx context.Context) (*Resume, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var resume Resume
	if err := m.Collection.FindOne(ctx, map[string]interface{}{}).Decode(&resume); err != nil {
		return nil, fmt.Errorf("failed to find resume: %w", err)
	}

	return &resume, nil
}