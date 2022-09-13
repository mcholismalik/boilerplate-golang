package es

import (
	"os"

	el "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var Client *el.Client

func Init() {
	var err error
	Client, err = el.NewClient(
		el.SetURL(os.Getenv("ELASTIC_URL_1")),
		el.SetSniff(false),
		// el.SetBasicAuth(os.Getenv("ELASTIC_USERNAME"), os.Getenv("ELASTIC_PASSWORD")),
	)
	if err != nil {
		panic(err)
	}
	logrus.Info("Elasticsearch successfully connected")
}
