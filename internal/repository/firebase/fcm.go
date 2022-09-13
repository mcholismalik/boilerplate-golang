package firebase

import (
	"context"
	"log"

	"firebase.google.com/go/messaging"
)

type (
	Fcm interface {
		FcmSend(ctx context.Context, topic, title, message string) error
		FcmSubscribeToTopic(ctx context.Context, topic string, tokens []string) error
	}

	fcm struct {
		client *messaging.Client
	}
)

func NewFcm(client *messaging.Client) Fcm {
	return &fcm{
		client: client,
	}
}

func (r *fcm) FcmSend(ctx context.Context, topic, title, message string) error {
	response, err := r.client.Send(ctx, &messaging.Message{
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

func (r *fcm) FcmSubscribeToTopic(ctx context.Context, topic string, tokens []string) error {
	response, err := r.client.SubscribeToTopic(ctx, tokens, topic)
	if err != nil {
		log.Println("error subscribe to topic", err)
		return err
	}

	log.Println("success fcm subscribe tot topic", response)

	return nil
}
