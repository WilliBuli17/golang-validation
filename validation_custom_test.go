package golang_validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type Seller struct {
	Name     string `validate:"varchar,custom1,custom2CrossField=Owner"`
	Owner    string `validate:"varchar,min=5"`
	Pin      string `validate:"required,custom2WithParam=6"`
	Identify string `validate:"required,email|numeric"` // ini pake or rule (|) jadi dapat bernilai true jika cuma salah satu di penuhi
}

func MustValidValue(level validator.FieldLevel) bool {
	value, success := level.Field().Interface().(string)
	if success {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}

	return true
}

func MustValidValueWithParam(level validator.FieldLevel) bool {
	regexNumber := regexp.MustCompile("^[0-9]+$")

	length, errConv := strconv.Atoi(level.Param())
	if errConv != nil {
		panic(errConv.Error())
	}

	value, success := level.Field().Interface().(string)
	if success {
		if !regexNumber.MatchString(value) {
			return false
		}
	}

	return len(value) == length
}

func MustEqualsIgnoreCase(level validator.FieldLevel) bool {
	value, _, _, ok := level.GetStructFieldOK2() // yang _ pertama itu tipe data, dan _ kedua itu datanya bisa nil apa tidak
	if !ok {
		panic("field not ok")
	}

	firstValue := strings.ToUpper(level.Field().Interface().(string))
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestAlias(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	err1 := validate.RegisterValidation("custom1", MustValidValue)
	if err1 != nil {
		t.Error(err1.Error())
	}

	err2 := validate.RegisterValidation("custom2WithParam", MustValidValueWithParam)
	if err2 != nil {
		t.Error(err2.Error())
	}

	err3 := validate.RegisterValidation("custom2CrossField", MustEqualsIgnoreCase)
	if err3 != nil {
		t.Error(err3.Error())
	}

	seller := Seller{
		Name:     "WILLI",
		Owner:    "Willi",
		Pin:      "123456",
		Identify: "wwww@www.www",
	}

	errValidate := validate.Struct(seller)

	if errValidate != nil {
		validationErrors := errValidate.(validator.ValidationErrors)
		for _, fieldsError := range validationErrors {
			t.Error("ERR :", fieldsError.Field(), "on tag", fieldsError.Tag(), "with error", fieldsError.Error())
		}
	}
}
