package errors

/*
Example error code is 5000102:
- 500 is HTTP status code (400, 401, 403, 500, ...)
- 01 is module represents for each handler
  - 00 for common error for all handler
  - 01 for health check handler

- 02 is actual error code, just auto increment and start at 1
*/
const ERR_3RD string = "3rd-code"

var (
	// Errors of module common
	// Format: ErrCommon<ERROR_NAME> = xxx00yy
	ErrCommonInternalServer    = ErrorCode("50000001")
	ErrCommonInvalidRequest    = ErrorCode("40000001")
	ErrCommonBindRequestError  = ErrorCode("40000002")
	ErrCommonExpiredToken      = ErrorCode("40100006")
	ErrAuthorizedNotPermission = ErrorCode("40000108")

	ErrUserFaAlreadyExist    = ErrorCode("40000113")
	ErrUserAlreadyExist      = ErrorCode("40000103")
	ErrInvalidOldPassword    = ErrorCode("40000104")
	ErrInvalidOldNewPassword = ErrorCode("40000105")
	ErrUserInputWrongOtp     = ErrorCode("40000106")
	ErrCommonUnauthorized    = ErrorCode("40100103")
	ErrUserBlocked           = ErrorCode("40000114")

	ErrEmailNotExist        = ErrorCode("40400109")
	ErrPhoneNotExist        = ErrorCode("40400104")
	ErrGmailNotExist        = ErrorCode("40400105")
	ErrFacebookNotExist     = ErrorCode("40400106")
	ErrAppleNotExist        = ErrorCode("40400107")
	ErrEmailAlreadyVerified = ErrorCode("40400100")
	ErrAlreadySendOtp       = ErrorCode("40400101")

	ErrorFileUploadNotNull  = ErrorCode("40000201")
	ErrorFileUploadMaximum  = ErrorCode("40000202")
	ErrorMimeTypeNotSupport = ErrorCode("40000203")
	ErrorUploadImageDrive   = ErrorCode("40000204")

	ErrorPaymentService         = ErrorCode("40400400")
	ErrorPaymentAlreadyEnrolled = ErrorCode("40000400")

	ErrorNotificationListEmpty = ErrorCode("40000500")
	ErrorNotificationMarkRead  = ErrorCode("40000501")

	ErrorDiscussionErrorRead = ErrorCode("40000600")
	ErrorRootThreadNotExist  = ErrorCode("400000601")

	ErrorOpenCombat       = ErrorCode("40400301")
	ErrorLeaderboardEmpty = ErrorCode("40400302")

	ErrorVoucherNotFound     = ErrorCode("40000700")
	ErrorVoucherApplyFail    = ErrorCode("40000701")
	ErrorVoucherApplyLock    = ErrorCode("40000702")
	ErrorVoucherInsufficient = ErrorCode("40000703")
)
