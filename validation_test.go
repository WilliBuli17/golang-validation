package golang_validation

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestValidation(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("Validate is nill")
	}
}

func TestValidationField(t *testing.T) {
	validate := validator.New()
	user := "Willi"

	err := validate.Var(user, "required") // -- required -- gunanya memastikan var user tidak kosong

	if err != nil {
		t.Error(err.Error())
	}
}

func TestValidationTwoVariables(t *testing.T) {
	validate := validator.New()

	password := "Rahasia"
	confirmPassword := "Rahasia"

	err := validate.VarWithValue(password, confirmPassword, "eqfield") // -- eqfield -- gunanya memastikan var password dan confirmPassword nilainya sama

	if err != nil {
		t.Error(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	value := "123"

	err := validate.Var(value, "required,numeric,min=2,max=5") // contoh Multiple Tag Validation dan Tag Parameter

	if err != nil {
		t.Error(err.Error())
	}
}
