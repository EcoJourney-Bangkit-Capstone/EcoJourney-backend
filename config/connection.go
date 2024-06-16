package config

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
	"google.golang.org/api/option"
)

var (
	FirebaseApp     *firebase.App
	AuthClient      *auth.Client
	FirestoreClient *firestore.Client
	StorageClient   *storage.Client
)

func InitFirebase() error {
	ctx := context.Background()
	sa := option.WithCredentialsFile("service-account.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return fmt.Errorf("error initializing firebase app: %v", err)
	}

	FirebaseApp = app

	AuthClient, err = app.Auth(ctx)
	if err != nil {
		return fmt.Errorf("error initializing auth client: %v", err)
	}

	FirestoreClient, err = app.Firestore(ctx)
	if err != nil {
		return fmt.Errorf("error initializing firestore client: %v", err)
	}

	StorageClient, err = app.Storage(ctx)
	if err != nil {
		return fmt.Errorf("error initializing storage client: %v", err)
	}

	return nil
}
