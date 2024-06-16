package config

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

// func Connect() (c *firebase.App, err error) {
// 	ctx := context.Background()

// 	client, err := app.Firestore(ctx)
// 	if err != nil {
// 		log.Fatalln(err)
// 		return nil, err
// 	}
// 	defer client.Close()

// 	return app, err
// }

func InitFirebaseApp(ctx context.Context) (*firebase.App, error) {
	conf := &firebase.Config{
		StorageBucket: os.Getenv("GOOGLE_BUCKET_NAME") + ".appspot.com",
	}
	opt := option.WithCredentialsFile("service-account.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return app, err
}

func GetStorageClient(app *firebase.App, ctx context.Context) (client *storage.Client, err error) {

	client, err = app.Storage(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetFirestoreClient(app *firebase.App, ctx context.Context) (client *firestore.Client, err error) {

	client, err = app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetAuthClient(app *firebase.App, ctx context.Context) (client *auth.Client, err error) {

	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return auth, nil
}
