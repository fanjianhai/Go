package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Users2 struct {
	Name   string `form:"name" json:"name" validate:"required"`
	Age    uint8  `form:"age" json:"age" validate:"required,gt=18"`
	Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	Code   string `form:"code" json:"code" validate:"required,len=6"`
}

func main() {
	users := &Users2{
		Name:   "admin",
		Age:    12,
		Passwd: "123",
		Code:   "123456",
	}
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(users)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(trans)) //Age必须大于18
			return
		}
	}

	return
}
