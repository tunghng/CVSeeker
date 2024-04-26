package cfg

const (
	KeyDBRedisAddr          = "DB_REDIS_ADDR"
	KeyDBRedisPassword      = "DB_REDIS_PASSWORD"
	KeyDBRedisMasterName    = "DB_REDIS_MASTER_NAME"
	KeyDBRedisSentinelAddrs = "DB_REDIS_SENTINEL_ADDRS"
	KeyDBRedisDB            = "DB_REDIS_DB"

	ElasticsearchUrl           = "ELK_URL"
	ElasticsearchKey           = "ELK_KEY"
	ElasticsearchUserName      = "ELK_USERNAME"
	ElasticsearchPassword      = "ELK_PASSWORD"
	ElasticsearchGtFilterIndex = "ELK_GT_FILTER_INDEX"
	ElasticsearchSaFilterIndex = "ELK_SA_FILTER_INDEX"

	ElasticsearchSaleIndex = "ELK_SALE_INDEX"
	RouteSourceFromDB      = "ROUTE_SOURCE_FROM_DB"

	ElasticsearchGtSubDFilterIndex = "ELK_GT_SUBD_FILTER_INDEX"

	ElasticsearchPromotionIndex = "ELK_PROMOTION_INDEX"

	KeyRestyEnableLog = "RESTY_ENABLE_LOG"

	KeyJiraHost             = "JIRA_HOST"
	KeyJiraUsername         = "JIRA_IMPERSONATE_EMAIL"
	KeyJiraUserID           = "JIRA_IMPERSONATE_ID"
	KeyJiraPassword         = "JIRA_TOKEN"
	KeyJiraTicketProjectKey = "JIRA_TICKET_PROJECT_KEY"
	KeyJiraTicketEnv        = "JIRA_TICKET_ENV"
	KeyAttachmentFolder     = "JIRA_GCS_ATTACHMENT_FOLDER"

	WpUrl      = "WP_URL"
	WpUsername = "WP_USERNAME"
	WpPassword = "WP_PASSWORD"

	GptApiKey = "GPT_API_KEY"

	ArangoDBUrl      = "ARANGO_HOST"
	ArangoDBPort     = "ARANGO_PORT"
	ArangoDBUsername = "ARANGO_USERNAME"
	ArangoDBPassword = "ARANGO_PASSWORD"
	ArangoDBName     = "ARANGO_DB"
)
