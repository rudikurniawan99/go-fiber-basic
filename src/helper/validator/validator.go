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
	// pattern := `^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$ %^&*-]).{8,}$`
	pattern := `kurniawan`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	return re.MatchString(fl.Field().String())
}

func ValidateUserPassword() error {
	validate := validator.New()
	user := &dtos.UserDTO{}
	validate.RegisterValidation("password", validatePassword)

	if err := validate.Struct(user); err != nil {
		return err
	}
	return nil
}
