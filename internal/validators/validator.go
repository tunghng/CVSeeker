package validators

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
)

const (
	excludeSpecialCharacters = "^[^~!@#$%^&*()+=\\/*+?,`<>%|:;\"'{}$]*$"
	phoneRegex               = `^(03|05|06|07|08|09|01[2|6|8|9])+([0-9]{8})$`
)

func NotAllowSpecial(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	data := fl.Field().String()
	ok, err := regexp.MatchString(excludeSpecialCharacters, data)
	if err != nil {
		return false
	}
	return ok
}

func IsPhone(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	phone := fl.Field().String()
	if len(phone) < 8 || len(phone) > 15 {
		return false
	}
	return regexp.MustCompile(phoneRegex).MatchString(phone)
}
