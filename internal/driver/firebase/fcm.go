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

func FcmSend(ctx context.Context, topic, title, message string) error {
	response, err := fcmClient.Send(ctx, &messaging.Message{
		Data: map[string]string{
			"title": title,
			"body":  message,
		},
		Webpush: &messaging.WebpushConfig{
			Notification: &messaging.WebpushNotification{
				Title: title,
				Body:  message,
			},
		},
		Topic: topic,
	})

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("success fcm send", response)

	return nil
}

func FcmSubscribeToTopic(ctx context.Context, topic string, tokens []string) error {
	response, err := fcmClient.SubscribeToTopic(ctx, tokens, topic)
	if err != nil {
		log.Println("error subscribe to topic", err)
		return err
	}

	log.Println("success fcm subscribe tot topic", response)

	return nil
}
