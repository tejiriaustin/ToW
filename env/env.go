package env

import (
	"log"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	conf *sync.Map
}

func init() {
	_ = godotenv.Load(".env")
}

func NewEnvironment() Config {
	return Config{
		conf: &sync.Map{},
	}
}

func (e *Config) SetEnv(name string, value interface{}) *Config {
	e.conf.Store(name, value)
	return e
}

func (e *Config) GetFloat64(key string) float64 {
	value, ok := e.conf.Load(key)
	if !ok {
		log.Fatal("couldn't load config value: ", key)
		return 0
	}

	valueAsString := value.(string)
	valueAsFloat, err := strconv.ParseFloat(valueAsString, 8)

	if err != nil {
		log.Fatal("couldn't parse value as string: ", err.Error())
		return 0
	}
	return valueAsFloat
}

func (e *Config) GetInt64(key string) int64 {
	value, ok := e.conf.Load(key)
	if !ok {
		log.Fatal("couldn't load config value: ", key)
		return 0
	}

	valueAsInt := value.(int)
	return int64(valueAsInt)
}

func (e *Config) GetAsString(key string) string {
	value, ok := e.conf.Load(key)
	if !ok {
		log.Fatal("couldn't load config value: ", key)
		return ""
	}
	valueAsString, ok := value.(string)
	if !ok {
		return ""
	}

	return valueAsString
}

func (e *Config) GetAsBytes(key string) []byte {
	value, ok := e.conf.Load(key)
	if !ok {
		log.Fatal("couldn't load config value: ", key)
		return nil
	}

	valueAsString, ok := value.(string)
	if !ok {
		return nil
	}

	return []byte(valueAsString)
}
