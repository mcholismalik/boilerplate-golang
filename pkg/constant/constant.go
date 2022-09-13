package constant

// General
const (
	APP      = "APP"
	APP_NAME = "github.com/mcholismalik/boilerplate-golang"
	PORT     = "PORT"
	ENV      = "ENV"
	VERSION  = "VERSION"
	HOST     = "HOST"
	SCHEME   = "SCHEME"
	JWT_KEY  = "JWT_KEY"

	CRON_ENABLED = "CRON_ENABLED"

	ELASTIC_URL_1        = "ELASTIC_URL_1"
	SENTRY_DSN           = "SENTRY_DSN"
	FIRESTORE_PROJECT_ID = "FIRESTORE_PROJECT_ID"
)

// Db
const (
	DB_DEFAULT_CREATED_BY = "system"
	DB_HOST               = "DB_HOST"
	DB_USER               = "DB_USER"
	DB_PASS               = "DB_PASS"
	DB_PORT               = "DB_PORT"
	DB_NAME               = "DB_NAME"
	DB_SSLMODE            = "DB_SSLMODE"
	DB_TZ                 = "DB_TZ"
	DB_GOBOIL_CLEAN       = "goboil_clean_db"

	MIGRATION_ENABLED = "MIGRATION_ENABLED"
	SEEDER_ENABLED    = "SEEDER_ENABLED"
)

const (
	LENGTH_CODE        = 20
	MAX_DATA_FIRESTORE = 50
)

type (
	contextKey string
	reqIDKey   string
)

const (
	CONTEXT_KEY contextKey = "context_key"
	REQ_ID_KEY  reqIDKey   = "req_id_key"
)
