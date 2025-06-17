package belajargolangvalidation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	user := "acumalaka"

	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationTwoVariable(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "rahasia"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()
	user := "123"

	err := validate.Var(user, "required,number")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	user := "12345"

	err := validate.Var(user, "required,number,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStructValidation(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "acumalaka@yahoo.com",
		Password: "acumalaka",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "acu",
		Password: "acu",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestCrossFieldValidation(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()
	registerRequest := RegisterUser{
		Username:        "acumalaka@yahoo.com",
		Password:        "acumalaka",
		ConfirmPassword: "acumalaka",
	}

	err := validate.Struct(registerRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Address: Address{
			City:    "",
			Country: "",
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCollectionValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{City: "", Country: ""},
			{City: "", Country: ""},
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicCollectionValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"required,dive,required,min=3"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{City: "", Country: ""},
			{City: "", Country: ""},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"X",
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMapValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,required"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{City: "", Country: ""},
			{City: "", Country: ""},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"X",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SDN Bahlil",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "",
			},
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicMapValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,required"`
		Wallet    map[string]int    `validate:"dive,keys,required,endkeys,required,gt=1000"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{City: "", Country: ""},
			{City: "", Country: ""},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"X",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SDN Bahlil",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "",
			},
		},
		Wallet: map[string]int{
			"BCA":     1000000,
			"Mandiri": 0,
			"":        1001,
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestAlias(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar,min=5"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:     "",
		Name:   "",
		Owner:  "",
		Slogan: "",
	}

	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}

	return true
}

func TestCustomValidationFunction(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	request := LoginRequest{
		Username: "ACUMALAKA",
		Password: "",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestCustomValidationParameter(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	request := Login{
		Phone: "091230912740931",
		Pin:   "123456",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrRule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	request := Login{
		Username: "123",
		Password: "123",
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}
