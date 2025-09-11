package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var (
	FirestoreClient *firestore.Client
)

const (
	BlogsCollection = "blogs"
	UsersCollection = "users"
)

// ✅ Init initializes only Firestore client (Storage removed)
func Init() error {
	// Load service account key from env variable
	credentialsPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credentialsPath == "" {
		// Fallback to hardcoded path if env var not set
		credentialsPath = "P:/blog-web/db/Firebase_Credentials.json"
	}

	// Initialize Firestore
	if err := InitFirestore(credentialsPath); err != nil {
		return fmt.Errorf("❌ Firestore initialization failed: %v", err)
	}

	log.Println("✅ Firestore initialized successfully")
	return nil
}

// ✅ InitFirestore creates the Firestore client
func InitFirestore(credentialsPath string) error {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "blog-web-d79ed", option.WithCredentialsFile(credentialsPath))
	if err != nil {
		return fmt.Errorf("failed to create Firestore client: %w", err)
	}

	FirestoreClient = client
	log.Println("✅ Firestore client initialized")
	return nil
}

// ✅ Close cleans up Firestore client
func Close() {
	if FirestoreClient != nil {
		err := FirestoreClient.Close()
		if err != nil {
			log.Printf("⚠️ Error closing Firestore: %v", err)
		} else {
			log.Println("✅ Firestore client closed")
		}
	}
}