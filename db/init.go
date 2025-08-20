package db

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/prachin77/blog-web/utils"
	"google.golang.org/api/option"
)

var (
	FirestoreClient *firestore.Client
	AuthClient      *auth.Client
	ctx             = context.Background()
)

const (
	UsersCollection = "users"
	BlogsCollection = "blogs"
)

func Init() error {
	fmt.Println("🔥 Initializing Firebase services...")

	projectID, err := utils.GetFirebaseProjectID("P:/blog-web/db/Firebase_Credentials.json")
	if err != nil {
		return fmt.Errorf("failed to get Firebase project ID: %v", err)
	}

	opt := option.WithCredentialsFile("P:/blog-web/db/Firebase_Credentials.json")
	config := &firebase.Config{ProjectID: projectID}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return fmt.Errorf("error initializing Firebase app: %v", err)
	}

	// Initialize Firestore
	FirestoreClient, err = app.Firestore(ctx)
	if err != nil {
		return fmt.Errorf("error getting Firestore client: %v", err)
	}
	fmt.Println("✅ Firestore initialized successfully")

	// Initialize Auth
	AuthClient, err = app.Auth(ctx)
	if err != nil {
		return fmt.Errorf("error getting Auth client: %v", err)
	}
	fmt.Println("✅ Auth client initialized successfully")

	return nil
}

func GetFirestore() *firestore.Client {
	return FirestoreClient
}

func GetAuth() *auth.Client {
	return AuthClient
}

func Close() {
	if FirestoreClient != nil {
		FirestoreClient.Close()
		fmt.Println("🔒 Firestore connection closed")
	}
}
