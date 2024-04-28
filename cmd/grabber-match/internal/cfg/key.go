package cfg

const (
	ConfigKeyEnvironment  = "ENVIRONMENT"
	ConfigKeyContextPath  = "CONTEXT_PATH"
	ConfigNewRelicKey     = "NEW_RELIC_KEY"
	ConfigNewRelicAppName = "NEW_RELIC_APP_NAME"

	ConfigKeyDBMySQLUsername = "DB_MYSQL_USERNAME"
	ConfigKeyDBMySQLPassword = "DB_MYSQL_PASSWORD"
	ConfigKeyDBMySQLHost     = "DB_MYSQL_HOST"
	ConfigKeyDBMySQLPort     = "DB_MYSQL_PORT"
	ConfigKeyDBMySQLDatabase = "DB_MYSQL_DATABASE"
	ConfigKeyDBMySQLLogBug   = "DB_MYSQL_LOG_BUG"

	ConfigKeyHttpAddress     = "HTTP_ADDR"
	ConfigKeyHttpPort        = "HTTP_PORT"
	ConfigApiDefaultPageSize = "API_DEFAULT_PAGE_SIZE"
	ConfigApiMinPageSize     = "API_MIN_PAGE_SIZE"
	ConfigApiMaxPageSize     = "API_MAX_PAGE_SIZE"

	JWTSecretKey           = "JWT_SECRET_KEY"
	JWTTokenExpireInMinute = "JWT_TOKEN_EXPIRE_IN_MINUTE"
	JWTPublicKeyFile       = "JWT_PUBLIC_KEY_FILE"
	JWTPrivateKeyFile      = "JWT_PRIVATE_KEY_FILE"

	KeyJwtExpiredTime = "JWT_EXPIRED_TIME"

	ConfigKeyFolderTmp              = "FOLDER_TMP"
	ConfigKeyGCSBucket              = "GCS_BUCKET"
	ConfigKeyGCSBucketCDN           = "GCS_BUCKET_CDN"
	ConfigKeyGCSBucketCDNRootFolder = "GCS_BUCKET_CDN_ROOT_FOLDER"
	ConfigKeyCDNUrl                 = "GCS_CDN_URL"
	URLGoogleStorage                = "URL_GOOGLE_STORAGE"
)
