package validators

import (
	"integration-suspect-service/modules/suspect/entities"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators(v *validator.Validate) {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	v.RegisterValidation("dateformat", DateValidator)
	v.RegisterValidation("noSpecialChar", NoSpecialCharValidator)
	v.RegisterValidation("isNegative", IsNegativeValidator)
	v.RegisterValidation("isCitizenID", CitizenIDValidator)
	v.RegisterValidation("isJuristicID", JuristicIDValidator)
	v.RegisterValidation("isName", NameValidator)
	v.RegisterValidation("isCompanyName", CompanyNameValidator)
}

func DateValidator(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	if dateStr == "" {
		return false
	}

	date, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return false
	}

	return date.Format("02-01-2006") == dateStr
}

func NoSpecialCharValidator(fl validator.FieldLevel) bool {
	specialCharRegex := regexp.MustCompile(`[!@#$%^&*()_+{}\[\]:;<>,.?~\\/]`)

	return !specialCharRegex.MatchString(fl.Field().String())
}

func OnlyNumberValidator(fl validator.FieldLevel) bool {
	digitsRegex := regexp.MustCompile("^[0-9]+$")

	return digitsRegex.MatchString(fl.Field().String())
}

func IsNegativeValidator(fl validator.FieldLevel) bool {
	kind := fl.Field().Kind()

	switch kind {
	case reflect.Int:
		return fl.Field().Interface().(int) < 0
	case reflect.String:
		valStr := fl.Field().Interface().(string)

		parseInt, err := strconv.Atoi(valStr)
		if err != nil {
			return false
		}
		return parseInt > 0
	default:
		return false
	}
}

func CitizenIDValidator(fl validator.FieldLevel) bool {

	value := fl.Field().String()
	requiredField := fl.Param()
	parentValue := fl.Parent().FieldByName(requiredField).String()

	if parentValue == entities.EntityTP_PERSON {
		isOnlyDigits := OnlyNumberValidator(fl)

		if isOnlyDigits && len(value) == 13 {
			return true
		}

	}

	if value != "" {
		return false
	}

	return true
}

func JuristicIDValidator(fl validator.FieldLevel) bool {

	value := fl.Field().String()
	requiredField := fl.Param()
	parentValue := fl.Parent().FieldByName(requiredField).String()

	if parentValue == entities.EntityTP_ENTITY {
		isNoSpecialChar := NoSpecialCharValidator(fl)

		if isNoSpecialChar {
			return true
		}

	}

	if value != "" {
		return false
	}

	return true
}

func NameValidator(fl validator.FieldLevel) bool {

	value := fl.Field().String()
	requiredField := fl.Param()
	parentValue := fl.Parent().FieldByName(requiredField).String()

	if parentValue == entities.EntityTP_PERSON {
		if value != "" {
			return true
		}
	}

	if value != "" {
		return false
	}

	return true
}

func CompanyNameValidator(fl validator.FieldLevel) bool {

	value := fl.Field().String()
	requiredField := fl.Param()
	parentValue := fl.Parent().FieldByName(requiredField).String()

	if parentValue == entities.EntityTP_ENTITY {
		if value != "" {
			return true
		}
	}

	if value != "" {
		return false
	}

	return true
}
