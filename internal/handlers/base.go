package handlers

import (
	"CVSeeker/internal/dtos"
	"CVSeeker/internal/errors"
	"CVSeeker/internal/ginLogger"
	"CVSeeker/internal/meta"
	"CVSeeker/internal/validators"
	"CVSeeker/pkg/db"
	pkgUtils "CVSeeker/pkg/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/vi"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/dig"
	"net/http"
	"reflect"
)

// BaseHandler is common handler for handlers or middlewares.
type BaseHandler interface {
	Validate(c *gin.Context, data interface{}) error
	RespondError(c *gin.Context, err error)
	HandleResponse(c *gin.Context, data interface{}, err error)
	BindData(c *gin.Context, data interface{}) bool
}

// baseHandlerParams contains all dependencies of BaseHandler.
type baseHandlerParams struct {
	dig.In
	DB          *db.DB `name:"talentAcquisitionDB"`
	ErrorParser errors.ErrorParser
}

// LogResponse log response
type LogResponse struct {
	Meta meta.Meta   `json:"meta"`
	Data interface{} `json:"data"`
}

// NewBaseHandler returns a new instance of BaseHandler.
func NewBaseHandler(params baseHandlerParams) BaseHandler {
	return &baseHandler{
		errorParser: params.ErrorParser,
		db:          params.DB,
	}
}

type baseHandler struct {
	db          *db.DB
	errorParser errors.ErrorParser
}

// RespondError RespondError
func (_this *baseHandler) RespondError(c *gin.Context, err error) {
	_this.HandleResponse(c, nil, err)
}

func (_this *baseHandler) HandleResponse(c *gin.Context, data interface{}, err error) {
	if err != nil {
		_this.processError(c, err, false)
		return
	}

	// add log response to context
	contextResponse, err := json.Marshal(data)
	if err == nil {
		c.Set(dtos.ContextResponse, contextResponse)
	}

	if data == nil || (reflect.ValueOf(data).Kind() == reflect.Ptr && reflect.ValueOf(data).IsNil()) {
		c.JSON(http.StatusOK, nil)
		return
	}

	c.JSON(http.StatusOK, data)
	c.Next()
}

// Validate struct
func (_this *baseHandler) Validate(c *gin.Context, dataRequest interface{}) error {
	english := en.New()
	uni := ut.New(english, vi.New())
	trans, _ := uni.GetTranslator("vi")
	validate := validators.NewValidatorV10(trans)
	return validate.ValidateStruct(dataRequest)
}

// Validate struct
func (_this *baseHandler) BindData(c *gin.Context, params interface{}) bool {
	if err := c.ShouldBindJSON(params); err != nil {
		ginLogger.Gin(c).Errorf("Failed to c.ShouldBindJSON error: %v", err)
		_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
		return false
	}
	if err := _this.Validate(c, params); err != nil {
		ginLogger.Gin(c).Errorf("Failed to Validate error: %v", err)
		_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
		return false
	}
	return true
}

func (_this *baseHandler) processError(c *gin.Context, err error, overrideHttpStatus bool) {
	statusCode, data := _this.errorParser.Parse(err)
	contextResponse, err := json.Marshal(data)
	if err == nil {
		c.Set(dtos.ContextResponse, contextResponse)
	}

	if overrideHttpStatus {
		statusCode = http.StatusOK
	}
	// longpv2 - fwd 3rd status code as it is
	switch statusCode {
	case 50001101, 50001102, 50001201, 50001301, 50001501:
		adapterStatusCode := c.GetInt(errors.ERR_3RD)
		if adapterStatusCode == 0 {
			adapterStatusCode = http.StatusInternalServerError
		}
		c.JSON(adapterStatusCode, data)
	}
	c.JSON(statusCode, data)
}

func (_this *baseHandler) cleanUnsafeField(data interface{}) string {
	return pkgUtils.Struct2JSON(data)
}
