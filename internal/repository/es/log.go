package es

import (
	"context"

	"github.com/mcholismalik/boilerplate-golang/internal/model/entity"
	el "github.com/olivere/elastic/v7"
)

const (
	INDEX_LOG_ERROR = "log_error"
)

type (
	Log interface {
		Base
		InsertErrorLog(ctx context.Context, log *entity.LogErrorEntity) error
	}

	log struct {
		Base
	}
)

func NewLog(client *el.Client) Log {
	base := NewBase(client)
	return &log{
		base,
	}
}

func (r *log) InsertErrorLog(ctx context.Context, log *entity.LogErrorEntity) error {
	return r.Insert(ctx, INDEX_LOG_ERROR, log)
}
