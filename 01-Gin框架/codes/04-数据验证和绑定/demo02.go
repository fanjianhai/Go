package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Users struct {
	Name   string `form:"name" json:"name" validate:"required,CustomValidationErrors"` //包含自定义函数
	Age    uint8  `form:"age" json:"age" validate:"required,gt=18"`
	Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	Code   string `form:"code" json:"code" validate:"required,len=6"`
}

func main() {

	users := &Users{
		Name:   "admin1",
		Age:    12,
		Passwd: "123",
		Code:   "123456",
	}
	validate := validator.New()
	//注册自定义函数
	_ = validate.RegisterValidation("CustomValidationErrors", CustomValidationErrors)
	err := validate.Struct(users)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err) //Key: 'Users.Name' Error:Field validation for 'Name' failed on the 'CustomValidationErrors' tag
			return
		}
	}
	return
}

func CustomValidationErrors(fl validator.FieldLevel) bool {
	return fl.Field().String() != "admin"
}
