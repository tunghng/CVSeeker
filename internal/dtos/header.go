package dtos

// Headers definition.
const (
	HeaderAuthorization     = "Authorization"
	HeaderXRequestID        = "X-Request-ID"
	HeaderXFcmToken         = "X_FCM_TOKEN"
	GinContextBasicUsername = "basic_username"
	GinContextLogRequest    = "log_request"
)

// Context key to transfer to next layer
const (
	ContextResponse = "x-response"
	ID              = "id"
	UserId          = "user_id"
	AccountNo       = "account_no"
	FullName        = "full_name"
)

const (
	BearerAuth = "bearer"
	BasicAuth  = "basic"
)
