package config

import (
	"fmt"
	"os"

	"go-rest/helper/exception"

	"github.com/joho/godotenv"
)

//Config ...
type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

//New config
func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	fmt.Println(err)
	exception.PanicIfNeeded(err)
	return &configImpl{}
}
