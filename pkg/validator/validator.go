package validator

import (
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslator "github.com/go-playground/validator/v10/translations/en"
)

type Validate = validator.Validate

var (
	once sync.Once
	vld  *Validate
)

func New() *Validate {
	once.Do(func() {
		vld = validator.New(validator.WithRequiredStructEnabled())
		enTranslator.RegisterDefaultTranslations(vld, GetTranslator())
	})
	return vld
}

func GetTranslator() ut.Translator {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	return trans
}
