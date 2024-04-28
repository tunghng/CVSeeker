package validators

import (
	"errors"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"log"
	"strings"
)

type ValidatorV10 struct {
	Validate *validator.Validate
	trans    ut.Translator
}

func NewValidatorV10(trans ut.Translator) ValidatorV10 {
	validate := validator.New()
	registerCustomerValidator(validate)

	// Translate
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)
	translateOverride(validate, trans)

	return ValidatorV10{
		Validate: validate,
		trans:    trans,
	}
}

func (_this *ValidatorV10) ValidateStruct(data interface{}) error {
	err := _this.Validate.Struct(data)
	if err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)
		errsTrans := errs.Translate(_this.trans)
		return errors.New(_this.printErrorMsg(errsTrans))
	}
	return nil
}

func (_this *ValidatorV10) printErrorMsg(errTrans validator.ValidationErrorsTranslations) string {
	var errMsgs []string
	for _, v := range errTrans {
		errMsgs = append(errMsgs, fmt.Sprintf("%s", v))
	}

	return strings.Join(errMsgs, ", ")
}

func registerCustomerValidator(validate *validator.Validate) {
	var err error
	err = validate.RegisterValidation("not_allow_special", NotAllowSpecial)
	if err != nil {
		log.Fatal(errors.New("validators LoadData for not_allow_special error: " + err.Error()))
	}
	err = validate.RegisterValidation("is_phone", IsPhone)
	if err != nil {
		log.Fatal(errors.New("validators LoadData for is_phone error: " + err.Error()))
	}
}

func translateOverride(validate *validator.Validate, trans ut.Translator) {
	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}
