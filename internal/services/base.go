package services

import (
	"CVSeeker/internal/errors"
	"CVSeeker/internal/meta"
	"CVSeeker/internal/validators"
	"context"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/vi"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/dig"
)

type BaseService interface {
	Validate(c context.Context, data interface{}) error
	HandlerError(c context.Context, err error) (*meta.Error, error)
}

// baseHandlerParams contains all dependencies of BaseHandler.
type baseHandlerParams struct {
	dig.In
	ErrorParser errors.ErrorParser
}

// NewBaseHandler returns a new instance of BaseHandler.
func NewBaseService(params baseHandlerParams) BaseService {
	return &baseService{
		errorParser: params.ErrorParser,
	}
}

type baseService struct {
	errorParser errors.ErrorParser
}

func (b *baseService) Validate(c context.Context, data interface{}) error {
	english := en.New()
	uni := ut.New(english, vi.New())
	trans, _ := uni.GetTranslator("en")
	validate := validators.NewValidatorV10(trans)
	return validate.ValidateStruct(data)
}

func (b *baseService) HandlerError(c context.Context, err error) (*meta.Error, error) {
	_, data := b.errorParser.Parse(err)
	return &data, nil
}
