package es

import (
	"context"

	el "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type (
	Base interface {
		Insert(ctx context.Context, index string, log interface{}) error
		Update(ctx context.Context, index, ID string, update map[string]interface{}) error
		Search(ctx context.Context, index string, searchSource *el.SearchSource) (*el.SearchResult, error)
	}

	base struct {
		client *el.Client
	}
)

func NewBase(client *el.Client) Base {
	return &base{
		client: client,
	}
}

func (r *base) Insert(ctx context.Context, index string, log interface{}) error {
	if _, err := r.client.Index().Index(index).
		Type("_doc").
		BodyJson(log).
		Do(ctx); err != nil {

		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "cannot insert data",
			"Index":         index,
			"Data":          log,
		}).Error(err.Error())
		return err
	}

	return nil
}

func (r *base) Update(ctx context.Context, index, ID string, update map[string]interface{}) error {
	if _, err := r.client.Update().
		Index(index).
		Type("_doc").
		Id(ID).Doc(update).Do(ctx); err != nil {

		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "cannot update data",
			"ID":            ID,
			"Index":         index,
			"Data":          update,
		}).Error(err.Error())
		return err
	}

	return nil
}

func (r *base) Search(ctx context.Context, index string, searchSource *el.SearchSource) (*el.SearchResult, error) {
	results, err := r.client.Search().
		Index(index).
		SearchSource(searchSource).
		Do(ctx)

	if err != nil {

		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "cannot search data",
		}).Error(err.Error())

		return nil, err
	}

	return results, nil
}
