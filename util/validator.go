package util

import (
	"fmt"
	"fortress/model"
	"github.com/go-playground/validator"
	"reflect"
	"strings"
)

var v *validator.Validate

func GetValidator() *validator.Validate {
	if v != nil {
		return v
	}

	v = validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	v.RegisterStructValidation(nil, model.User{})
	return v
}

func GetErrors(e error) []model.ValidationError {
	var errors []model.ValidationError
	for _, err := range e.(validator.ValidationErrors) {
		errors = append(errors, model.ValidationError{
			Namespace:       err.Namespace(),
			Field:           err.Field(),
			StructNamespace: err.StructNamespace(),
			StructField:     err.StructField(),
			Tag:             err.Tag(),
			ActualTag:       err.ActualTag(),
			Kind:            fmt.Sprintf("%v", err.Kind()),
			Type:            fmt.Sprintf("%v", err.Type()),
			Value:           fmt.Sprintf("%v", err.Value()),
			Param:           err.Param(),
		})
	}

	return errors
}
