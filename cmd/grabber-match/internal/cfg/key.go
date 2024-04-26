package cfg

const (
	ConfigKeyEnvironment  = "ENVIRONMENT"
	ConfigKeyContextPath  = "CONTEXT_PATH"
	ConfigNewRelicKey     = "NEW_RELIC_KEY"
	ConfigNewRelicAppName = "NEW_RELIC_APP_NAME"

	ConfigKeyDBMySQLUsername         = "DB_MYSQL_USERNAME"
	ConfigKeyDBMySQLPassword         = "DB_MYSQL_PASSWORD"
	ConfigKeyDBMySQLHost             = "DB_MYSQL_HOST"
	ConfigKeyDBMySQLPort             = "DB_MYSQL_PORT"
	ConfigKeyDBMySQLDatabase         = "DB_MYSQL_DATABASE"
	ConfigKeyDBMaxIdleConnections    = "DB_MYSQL_MAX_IDLE_CONNECTIONS"
	ConfigKeyDBMaxOpenConnections    = "DB_MYSQL_MAX_OPEN_CONNECTIONS"
	ConfigKeyDBConnectionMaxLifetime = "DB_MYSQL_CONNECTION_MAX_LIFETIME"
	ConfigKeyDBMySQLLogBug           = "DB_MYSQL_LOG_BUG"

	ConfigKeyDBFaQuizMySQLUsername = "DB_MYSQL_FA_QUIZ_USERNAME"
	ConfigKeyDBFaQuizMySQLPassword = "DB_MYSQL_FA_QUIZ_PASSWORD"
	ConfigKeyDBFaQuizMySQLDatabase = "DB_MYSQL_FA_QUIZ_DATABASE"

	ConfigKeyHttpAddress     = "HTTP_ADDR"
	ConfigKeyHttpPort        = "HTTP_PORT"
	ConfigApiDefaultPageSize = "API_DEFAULT_PAGE_SIZE"
	ConfigApiMinPageSize     = "API_MIN_PAGE_SIZE"
	ConfigApiMaxPageSize     = "API_MAX_PAGE_SIZE"

	JWTSecretKey           = "JWT_SECRET_KEY"
	JWTTokenExpireInMinute = "JWT_TOKEN_EXPIRE_IN_MINUTE"
	JWTPublicKeyFile       = "JWT_PUBLIC_KEY_FILE"
	JWTPrivateKeyFile      = "JWT_PRIVATE_KEY_FILE"

	ConfigFireStoreProjectID           = "CLOUD_FIRESTORE_PROJECT_ID"
	ConfigFireStoreCredentials         = "CLOUD_FIRESTORE_FILE_CREDENTIALS"
	ConfigFireStoreImpersonatedAccount = "CLOUD_FIRESTORE_IMPERSONATED_ACCOUNT"
	ConfigFireStoreRootCollectionSales = "CLOUD_FIRESTORE_ROOT_COLLECTION_SALES"

	ConfigKafkaBrokers        = "KAFKA_BROKERS"
	ConfigKafkaTopics         = "KAFKA_TOPICS"
	ConfigKafkaGroupID        = "KAFKA_GROUP_ID"
	ConfigKafkaEnableSSL      = "KAFKA_ENABLE_SSL"
	ConfigKafkaClientCertFile = "KAFKA_CLIENT_CERT_FILE"
	ConfigKafkaClientKeyFile  = "KAFKA_CLIENT_KEY_FILE"
	ConfigKafkaCACertFile     = "KAFKA_CA_CERT_FILE"

	KeyJwtExpiredTime = "JWT_EXPIRED_TIME"

	ConfigKeyFolderTmp              = "FOLDER_TMP"
	ConfigKeyGCSBucket              = "GCS_BUCKET"
	ConfigKeyGCSBucketCDN           = "GCS_BUCKET_CDN"
	ConfigKeyGCSBucketCDNRootFolder = "GCS_BUCKET_CDN_ROOT_FOLDER"
	ConfigKeyCDNUrl                 = "GCS_CDN_URL"
	URLGoogleStorage                = "URL_GOOGLE_STORAGE"

	ConfigPubSubProjectID              = "CLOUD_PUBSUB_PROJECT_ID"
	ConfigPubSubRegion                 = "CLOUD_PUBSUB_REGION"
	ConfigPubSubCredentials            = "CLOUD_PUBSUB_FILE_CREDENTIALS"
	ConfigPubSubMaxOutstandingMessages = "CLOUD_PUBSUB_MAX_OUTSTANDING_MESSAGE"

	//-----------------------------
	TempConfigKeyDBWordpressUsername = "WORDPRESSDB_USERNAME"
	TempConfigKeyDBWordpressPassword = "WORDPRESSDB_PASSWORD"
	TempConfigKeyDBWordpressHost     = "WORDPRESSDB_HOST"
	TempConfigKeyDBWordpressPort     = "WORDPRESSDB_PORT"
	TempConfigKeyDBWordpressDatabase = "WORDPRESSDB_DATABASE"

	TempConfigKeyDBFaClassUsername = "FACLASSDB_USERNAME"
	TempConfigKeyDBFaClassPassword = "FACLASSDB_PASSWORD"
	TempConfigKeyDBFaClassHost     = "FACLASSDB_HOST"
	TempConfigKeyDBFaClassPort     = "FACLASSDB_PORT"
	TempConfigKeyDBFaClassDatabase = "FACLASSDB_DATABASE"
)
