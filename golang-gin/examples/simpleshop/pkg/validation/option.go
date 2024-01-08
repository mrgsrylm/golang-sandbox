package validation

import (
	"reflect"
	"strings"

	enLocales "github.com/go-playground/locales/en"
	universalT "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

// Option validation option
type Option interface {
	apply(*option)
}

// option implement
type option struct {
	validator *validator.Validate
	uni       *universalT.UniversalTranslator
	trans     *universalT.Translator
}

type optionFn func(*option)

func (optFn optionFn) apply(opt *option) {
	optFn(opt)
}

// WithValidator set validator
func WithValidator(v *validator.Validate) Option {
	return optionFn(func(opt *option) {
		opt.validator = v
	})
}

// WithUniversalTranslator set UniversalTranslator
func WithUniversalTranslator(uni *universalT.UniversalTranslator) Option {
	return optionFn(func(opt *option) {
		opt.uni = uni
	})
}

// WithTranslator set Translator
func WithTranslator(trans *universalT.Translator) Option {
	return optionFn(func(opt *option) {
		opt.trans = trans
	})
}

func getDefaultOption() *option {
	v := validator.New()

	translator := enLocales.New()
	uni := universalT.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		return nil
	}

	if err := enTranslations.RegisterDefaultTranslations(v, trans); err != nil {
		return nil
	}

	_ = v.RegisterTranslation("password", trans, func(ut universalT.Translator) error {
		return ut.Add("password", "{0} is not strong enough, password must be at least 6 characters", true)
	}, func(ut universalT.Translator, fe validator.FieldError) string {
		t, _ := ut.T("password", fe.Field())
		return t
	})

	_ = v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) >= 6
	})

	_ = v.RegisterTranslation("countryCode", trans, func(ut universalT.Translator) error {
		return ut.Add("countryCode", "{0} must be at least 2 characters and start with '+'", true)
	}, func(ut universalT.Translator, fe validator.FieldError) string {
		t, _ := ut.T("countryCode", fe.Field())
		return t
	})

	_ = v.RegisterValidation("countryCode", func(fl validator.FieldLevel) bool {
		codeLen := len(fl.Field().String())
		if codeLen == 0 {
			return true
		}
		if codeLen < 2 || !strings.HasPrefix(fl.Field().String(), "+") {
			return false
		}
		return true
	})

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		jsonTag := fld.Tag.Get("json")
		if jsonTag == "" {
			return fld.Name
		}

		name := strings.SplitN(jsonTag, ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	return &option{
		validator: v,
		uni:       uni,
		trans:     &trans,
	}
}

func getOption(opts ...Option) *option {
	opt := getDefaultOption()
	for _, o := range opts {
		o.apply(opt)
	}

	return opt
}
