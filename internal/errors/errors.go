package errors

import (
	"fmt"
	"go.uber.org/dig"
	"grabber-match/internal/meta"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

// Errors const definition.
const (
	DefaultErrorStatus  = http.StatusInternalServerError
	DefaultErrorCode    = 5000001
	DefaultErrorMessage = "Hệ thống xảy ra lỗi khi xử lý yêu cầu"
)

type ErrorParserConfig struct {
	PathConfigError string
}

// DefaultError constant definition.
var DefaultError = meta.Error{Meta: meta.New(DefaultErrorCode, nil, DefaultErrorMessage)}

type errorParserParams struct {
	dig.In
	Config *ErrorParserConfig
}

func NewErrorParser(params errorParserParams) ErrorParser {
	t, err := toml.LoadFile(filepath.Clean(params.Config.PathConfigError))

	if err != nil {
		log.Fatal(Wrap(err, "loading errors.toml file"))
	}

	return &errorParser{tree: t, config: params.Config}
}

// ErrorParser parses business errors to response.
type ErrorParser interface {
	Parse(err error) (int, meta.Error)
}

type errorParser struct {
	tree   *toml.Tree
	config *ErrorParserConfig
}

func (_this *errorParser) Parse(err error) (int, meta.Error) {
	var cusErr CustomError
	cusErr, ok := err.(CustomError)
	if !ok {
		return http.StatusInternalServerError, DefaultError
	}
	var (
		errCode     = cusErr.Code
		modulesTree = _this.tree.Get("modules").(*toml.Tree)
	)
	status, err := strconv.Atoi(errCode[0:3])
	if err != nil {
		return DefaultErrorStatus, DefaultError
	}
	errCodeNum, err := strconv.ParseInt(errCode, 10, 64)
	if err != nil {
		return DefaultErrorStatus, DefaultError
	}
	// for unwrapped error
	if len(errCode) != 8 {
		if cusErr.Params != nil {
			return status, meta.Error{Meta: meta.New(int(errCodeNum), cusErr.Params[0], err.Error())}
		}
		return status, meta.Error{Meta: meta.New(int(errCodeNum), nil, err.Error())}
	}
	errModStr := errCode[3:6]
	if !modulesTree.Has(errModStr) {
		return DefaultErrorStatus, DefaultError
	}
	modKeyStr := modulesTree.Get(errModStr).(string)

	if !_this.tree.Has(modKeyStr) {
		return DefaultErrorStatus, DefaultError
	}
	errorModuleTree := _this.tree.Get(modKeyStr).(*toml.Tree)

	if err != nil {
		return DefaultErrorStatus, DefaultError
	}
	errMsg := errorModuleTree.Get(errCode).(string)

	if cusErr.Params != nil {
		return status, meta.Error{Meta: meta.New(int(errCodeNum), cusErr.Params[0], fmt.Sprintf(errMsg))}
	}
	return status, meta.Error{Meta: meta.New(int(errCodeNum), nil, fmt.Sprintf(errMsg, cusErr.Params...))}
}

// Wrap wraps a normal error.
func Wrap(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

// New returns new error.
func New(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

// NewAliceCustom returns new error.
func NewAliceCustom(code int, args ...interface{}) error {
	return CustomError{
		Code:   fmt.Sprintf("%d", code),
		Params: args,
	}
}

// NewCusErr returns new CustomError as error.
func NewCusErr(err error, args ...interface{}) error {
	code := strconv.Itoa(DefaultErrorCode)
	if errCode, ok := err.(ErrorCode); ok {
		code = errCode.Error()
	}
	return CustomError{
		Code:   code,
		Params: args,
	}
}

// ErrorCode is error type code.
type ErrorCode string

func (_this ErrorCode) Error() string {
	return string(_this)
}

// CustomError is merchant integration custom error.
type CustomError struct {
	Code   string
	Params []interface{}
}

func (_this CustomError) Error() string {
	return fmt.Sprintf("Code: %v, Params: %v", _this.Code, _this.Params)
}
