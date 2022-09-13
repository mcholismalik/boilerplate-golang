package sentry

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func Init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		TracesSampleRate: 1.0,
		ServerName:       fmt.Sprintf("%s v%s", os.Getenv("APP"), os.Getenv("VERSION")),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sentry.Flush(2 * time.Second)
}
