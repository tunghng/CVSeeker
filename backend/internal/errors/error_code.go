package errors

/*
Example error code is 5000102:
- 500 is HTTP status code (400, 401, 403, 500, ...)
- 01 is module represents for each handler
  - 00 for common error for all handler
  - 01 for health check handler

- 02 is actual error code, just auto increment and start at 1
*/

var (
	// Errors of module common
	// Format: ErrCommon<ERROR_NAME> = xxx00yy
	ErrCommonInternalServer    = ErrorCode("50000001")
	ErrCommonInvalidRequest    = ErrorCode("40000001")
	ErrCommonBindRequestError  = ErrorCode("40000002")
	ErrCommonExpiredToken      = ErrorCode("40100006")
	ErrAuthorizedNotPermission = ErrorCode("40000108")
)
