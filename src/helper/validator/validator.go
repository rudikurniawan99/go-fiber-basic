package validator

import "github.com/go-playground/validator/v10"

func Validate[T any](dto *T) error {
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
