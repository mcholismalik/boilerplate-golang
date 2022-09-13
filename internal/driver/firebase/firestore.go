package firebase

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/mcholismalik/boilerplate-golang/pkg/constant"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var firestoreClient *firestore.Client
var firestoreOnce sync.Once

func FirestoreInit() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./firebase-admin-sdk.json")
	config := &firebase.Config{ProjectID: os.Getenv(constant.FIRESTORE_PROJECT_ID)}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	firestoreOnce.Do(func() {
		firestoreClient = client
	})
}

func GetFirestoreClient() *firestore.Client {
	return firestoreClient
}
