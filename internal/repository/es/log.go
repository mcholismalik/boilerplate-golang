package es

import (
	"context"

	"github.com/mcholismalik/boilerplate-golang/internal/driver/es"
	"github.com/mcholismalik/boilerplate-golang/internal/model/entity"
)

const (
	INDEX_LOG_ERROR = "log_error"
)

func InsertErrorLog(ctx context.Context, log *entity.LogErrorEntity) error {
	return es.Insert(ctx, INDEX_LOG_ERROR, log)
}
