package firebase

import (
	"context"

	fs "cloud.google.com/go/firestore"
	"github.com/mcholismalik/boilerplate-golang/pkg/constant"
	"google.golang.org/api/iterator"
)

type (
	Firestore interface {
		FirestoreAddData(collection string, ID string, data interface{}) error
		FirestoreUpdateData(collection string, doc string, data interface{}) (err error)
	}

	firestore struct {
		client *fs.Client
	}
)

func NewFirestore(client *fs.Client) Firestore {
	return &firestore{
		client: client,
	}
}

func (r *firestore) FirestoreAddData(collection string, ID string, data interface{}) error {
	err := r.client.RunTransaction(context.Background(), func(ctx context.Context, tx *fs.Transaction) error {

		// check limit
		sanpShot, err := r.client.Collection(collection).Documents(ctx).GetAll()
		if err != nil {
			return err
		}

		// if len snapshot >= max data
		// before we add data, we need to delete first data
		if len(sanpShot) >= constant.MAX_DATA_FIRESTORE {
			query := r.client.Collection(collection).OrderBy("created", fs.Asc).Limit(1)
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

		err = tx.Create(r.client.Collection(collection).Doc(ID), data)
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *firestore) FirestoreUpdateData(collection string, doc string, data interface{}) (err error) {
	err = r.client.RunTransaction(context.Background(), func(ctx context.Context, tx *fs.Transaction) (err error) {
		totalOrders := r.client.Collection(collection)

		totalOrder := totalOrders.Doc(doc)
		_, err = totalOrder.Set(ctx, data)

		return
	})

	return
}
