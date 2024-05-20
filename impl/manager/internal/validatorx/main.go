package validatorx

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var instance *validator.Validate
var once sync.Once

func New() *validator.Validate {
	once.Do(func() {
		instance = validator.New()
	})

	return instance
}
