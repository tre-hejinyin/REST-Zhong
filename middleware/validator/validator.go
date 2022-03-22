package validator

import (
	"regexp"
	"unicode"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"server/middleware/logger"
)

type Regexp = string

const (
	PASSWORD Regexp = `^[A-Za-z0-9]{6,16}$`
)

func Register() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("password", checkPassword)
		if err != nil {
			logger.Error("validator password", zap.Error(err))
		}
	}
}

// checkPassword
func checkPassword(fl validator.FieldLevel) bool {
	v, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	return verifyPassword(v)
}

func verifyPassword(pw string) bool {
	if !regexp.MustCompile(PASSWORD).MatchString(pw) {
		return false
	}
	var num, letter bool
	for _, r := range pw {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsLetter(r):
			letter = true
		}
	}
	return num && letter
}
