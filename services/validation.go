package services

import (
	"errors"
	"log"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

var vd *validator.Validate
var trans ut.Translator

func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// Validate checking data
func Validate(data interface{}) (map[string]string, bool, error) {
	var nameOfStruct string
	if t := reflect.TypeOf(data); t.Kind() == reflect.Ptr {
		nameOfStruct = t.Elem().Name()
	} else {
		nameOfStruct = t.Name()
	}

	err := vd.Struct(data)
	if err != nil {
		errorMap := map[string]string{}
		errorType := ""
		for _, e := range err.(validator.ValidationErrors) {
			if errorType == "" {
				errorType = e.Tag()
				errorMap[LcFirst(e.Field())] = formatMessage(e, trans, nameOfStruct)
			} else {
				if e.Tag() == errorType {
					errorMap[LcFirst(e.Field())] = formatMessage(e, trans, nameOfStruct)
				}
			}
		}
		return errorMap, true, nil
	}
	return map[string]string{}, false, nil
}

// Validation for initial function validation
func Validation() error {
	vd = validator.New()
	_ = vd.RegisterValidation("phone_number", ValidatePhone)
	translator, err := setValidate(vd)
	trans = translator
	if err != nil {
		log.Println("[Error] Validation.Validation : ", err)
		return err
	}
	return nil
}

func setValidate(v *validator.Validate) (ut.Translator, error) {
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Println("[Error] Configuration.setValidate : translator not found")
		return nil, errors.New("translator not found")
	}

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Println("[Error] Configuration.setValidate : ", err)
		return nil, err
	}

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("validName"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return trans, nil
}

func formatMessage(err validator.FieldError, trans ut.Translator, nameOfStruct string) string {
	message := ""
	field := ""

	field = Message("validateMessage.field."+err.Field(), nil)
	if field == "" {
		field = err.Field()
	}

	switch err.Tag() {
	case "required":
		message = Message("validateMessage.required", nil)
	case "phone_number":
		message = Message("validateMessage.phoneNumber", nil)
	default:
		message = err.Translate(trans)
	}
	return LcFirst(message)
}

// ValidatePhone custom validation for phone number
func ValidatePhone(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^[0-9]{5,12}$`)
	return regex.MatchString(fl.Field().String())
}
