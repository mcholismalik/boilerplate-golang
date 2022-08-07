package log

import (
	"context"

	"github.com/mcholismalik/boilerplate-golang/pkg/elasticsearch"
)

const (
	INDEX_LOG_ERROR = "log_error"
)

func InsertErrorLog(ctx context.Context, log *LogError) error {
	return elasticsearch.Insert(ctx, INDEX_LOG_ERROR, log)
}
