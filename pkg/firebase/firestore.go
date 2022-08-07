package firebase

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/mcholismalik/boilerplate-golang/pkg/constant"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
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

func FirestoreAddData(collection string, ID string, data interface{}) error {
	err := firestoreClient.RunTransaction(context.Background(), func(ctx context.Context, tx *firestore.Transaction) error {

		// check limit
		sanpShot, err := firestoreClient.Collection(collection).Documents(ctx).GetAll()
		if err != nil {
			return err
		}

		// if len snapshot >= max data
		// before we add data, we need to delete first data
		if len(sanpShot) >= constant.MAX_DATA_FIRESTORE {
			query := firestoreClient.Collection(collection).OrderBy("created", firestore.Asc).Limit(1)
			docs := tx.Documents(query)
			for {
				doc, err := docs.Next()
				if err != nil {
					if err == iterator.Done {
						break
					}
					if err != nil {
						return err
					}
				}
				// delete data
				err = tx.Delete(doc.Ref)
				if err != nil {
					return err
				}
			}
		}

		err = tx.Create(firestoreClient.Collection(collection).Doc(ID), data)
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

func FirestoreUpdateData(collection string, doc string, data interface{}) (err error) {
	err = firestoreClient.RunTransaction(context.Background(), func(ctx context.Context, tx *firestore.Transaction) (err error) {
		totalOrders := firestoreClient.Collection(collection)

		totalOrder := totalOrders.Doc(doc)
		_, err = totalOrder.Set(ctx, data)

		return
	})

	return
}
