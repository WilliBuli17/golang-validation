package golang_validation

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

type User struct {
	Username        string            `validate:"required"`
	Email           string            `validate:"required,email"`
	Password        string            `validate:"required,min=6,max=8"`
	ConfirmPassword string            `validate:"required,min=6,max=8,eqfield=Password"`
	UserAddresses   []Address         `validate:"required,dive"`
	Hobbies         []string          `validate:"required,dive,required,min=1"`
	Schools         map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
	Wallet          map[string]int    `validate:"required,dive,keys,required,min=2,endkeys,required,gt=10"`
}

type Address struct {
	City    string `validate:"required"`
	Country string `validate:"required"`
}

type School struct {
	Name string `validate:"required"`
}

func TestUserStructValidation(t *testing.T) {
	validation := validator.New()

	user := User{
		Username:        "Willi",
		Email:           "Willi@email.com",
		Password:        "Rahasia",
		ConfirmPassword: "Rahasia",
		UserAddresses: []Address{
			{
				City:    "Antah Berantah",
				Country: "Fatamorgana",
			},
			{
				City:    "Antah Berantah 2",
				Country: "Fatamorgana 2",
			},
		},
		Hobbies: []string{
			"Hobi 1", "Hobi 2", "Hobi 3",
		},
		Schools: map[string]School{
			"TK": {
				Name: "TK",
			},
			"SD": {
				Name: "SD",
			},
		},
		Wallet: map[string]int{
			"Wallet1": 1_000_000,
			"Wallet2": 1_000_000,
		},
	}

	err := validation.Struct(user)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldsError := range validationErrors {
			t.Error("ERR :", fieldsError.Field(), "on tag", fieldsError.Tag(), "with error", fieldsError.Error())
		}
	}
}
