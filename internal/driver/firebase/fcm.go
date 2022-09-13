package firebase

import (
	"context"
	"log"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var fcmClient *messaging.Client
var fcmOnce sync.Once

func FcmInit() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./firebase-admin-sdk.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Messaging(context.TODO())
	if err != nil {
		log.Fatalf("messaging: %s", err)
	}

	fcmOnce.Do(func() {
		fcmClient = client
	})
}

func GetFcmClient() *messaging.Client {
	return fcmClient
}
