package db

import (
    "context"
    "fmt"

    "github.com/prachin77/blog-web/model"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/api/iterator"
)

// CheckEmailExists checks if a user with the given email already exists
func CheckEmailExists(email string) (bool, error) {
    if FirestoreClient == nil {
        return false, fmt.Errorf("firestore client not initialized")
    }

    ctx := context.Background()

    // Query users collection for the email
    iter := FirestoreClient.Collection(UsersCollection).
        Where("email", "==", email).
        Limit(1).
        Documents(ctx)

    defer iter.Stop()

    // Check if any document exists
    _, err := iter.Next()
    if err != nil {
        // Check if it's a "not found" error (which means email doesn't exist)
        if err == iterator.Done {
            // No document found - email doesn't exist
            return false, nil
        }
        // Some other error occurred
        return false, fmt.Errorf("error checking email existence: %v", err)
    }

    // Document found - email exists
    return true, nil
}

// CreateUser creates a new user in the Firestore database and returns the Firebase-generated ID
func CreateUser(user *model.User) (string, error) {
    if FirestoreClient == nil {
        return "", fmt.Errorf("firestore client not initialized")
    }

    ctx := context.Background()

    // Generate a new document reference with auto-generated ID
    docRef := FirestoreClient.Collection(UsersCollection).NewDoc()
    
    // Add the user to Firestore
    _, err := docRef.Set(ctx, user)
    if err != nil {
        return "", fmt.Errorf("error creating user: %v", err)
    }

    // Return the Firebase-generated document ID
    fmt.Printf("✅ User created successfully with ID: %s\n", docRef.ID)
    return docRef.ID, nil
}

// GetUserByEmail retrieves a user by email along with their ID
func GetUserByEmail(email string) (*model.User, string, error) {
    if FirestoreClient == nil {
        return nil, "", fmt.Errorf("firestore client not initialized")
    }

    ctx := context.Background()

    iter := FirestoreClient.Collection(UsersCollection).
        Where("email", "==", email).
        Limit(1).
        Documents(ctx)

    defer iter.Stop()

    doc, err := iter.Next()
    if err != nil {
        if err == iterator.Done {
            return nil, "", fmt.Errorf("user not found")
        }
        return nil, "", fmt.Errorf("error getting user: %v", err)
    }

    var user model.User
    err = doc.DataTo(&user)
    if err != nil {
        return nil, "", fmt.Errorf("error parsing user data: %v", err)
    }

    return &user, doc.Ref.ID, nil
}

// GetUserByID retrieves a user by their Firebase document ID
func GetUserByID(userID string) (*model.User, error) {
    if FirestoreClient == nil {
        return nil, fmt.Errorf("firestore client not initialized")
    }

    ctx := context.Background()

    doc, err := FirestoreClient.Collection(UsersCollection).Doc(userID).Get(ctx)
    if err != nil {
        if status.Code(err) == codes.NotFound {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("error getting user: %v", err)
    }

    var user model.User
    err = doc.DataTo(&user)
    if err != nil {
        return nil, fmt.Errorf("error parsing user data: %v", err)
    }

    return &user, nil
}