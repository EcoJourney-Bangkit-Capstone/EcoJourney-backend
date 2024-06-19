package config

import (
	"context"
	"fmt"
	"mime/multipart"
	"io"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"cloud.google.com/go/storage"
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

	StorageClient, err = storage.NewClient(ctx, sa)
	if err != nil {
		return fmt.Errorf("error initializing storage client: %v", err)
	}

	return nil
}

func UploadImageToGCS(file multipart.File, fileName string) (string, error) {
	ctx := context.Background()
	bucketName := os.Getenv("GOOGLE_BUCKET_NAME")

	fmt.Printf("Uploading image with fileName: %s\n", fileName)
	fmt.Printf("Bucket Name: %s\n", bucketName)

	wc := StorageClient.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("error closing writer: %v", err)
	}

	imageURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, fileName)
	return imageURL, nil
}
