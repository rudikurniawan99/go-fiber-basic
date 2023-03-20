package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/rudikurniawan99/go-fiber-basic/src/dtos"
)

func Validate[T any](dto *T) error {
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}

func CustomValidation[T any](dto *T, tag string, fn validator.Func) error {
	validate := validator.New()
	validate.RegisterValidation(tag, fn)
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}

func validatePassword(fl validator.FieldLevel) bool {
	pattern := `^[A-Za-z0-9#?!@\$%\^&\*\-]{8,}$`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	return re.MatchString(fl.Field().String())
}

func ValidateUserPassword(user *dtos.UserDTO) error {
	validate := validator.New()
	validate.RegisterValidation("password", validatePassword)

	if err := validate.Struct(user); err != nil {
		return err
	}
	return nil
}
