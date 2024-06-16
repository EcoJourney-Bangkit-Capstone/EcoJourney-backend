package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Connect() (c *firebase.App, err error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("service-account.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer client.Close()

	return app, err
}
