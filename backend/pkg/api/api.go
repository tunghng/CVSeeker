package api

const (
	XForwardUserOpsHeader = "X-Forward-User"
	OsTypeHeader          = "Os-Type"
	OsVersionHeader       = "Os-Version"
)

type Config struct {
	DefaultPageSize int64
	MinPageSize     int64
	MaxPageSize     int64
}
