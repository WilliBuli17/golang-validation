package golang_validation

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"testing"
)

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(RegisterRequest)

	if strings.ToUpper(registerRequest.Username) == strings.ToUpper(registerRequest.Email) ||
		strings.ToUpper(registerRequest.Username) == strings.ToUpper(registerRequest.Phone) {
		//success
	} else {
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}
}

func TestValidateStructLevel(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	registerRequest := RegisterRequest{
		Username: "willi@exampleMail.com",
		Email:    "willi@exampleMail.com",
		Phone:    "1835982590250",
		Password: "2340205",
	}

	errValidate := validate.Struct(registerRequest)

	if errValidate != nil {
		validationErrors := errValidate.(validator.ValidationErrors)
		for _, fieldsError := range validationErrors {
			t.Error("ERR :", fieldsError.Field(), "on tag", fieldsError.Tag(), "with error", fieldsError.Error())
		}
	}
}
