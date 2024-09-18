package goconfigloader

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Data map[string]string
}

// NewConfigLoader will initialize new config loader with optional default map value
func NewConfigLoader(defaultValues ...map[string]string) *Config {
	defaultMap := map[string]string{}
	if len(defaultValues) > 0 {
		defaultMap = defaultValues[0]
	}

	config := Config{
		Data: defaultMap,
	}

	// set default value that is not in OS level
	for key, value := range defaultMap {
		osVal := os.Getenv(key)
		if len(osVal) == 0 {
			os.Setenv(key, value)
		}
	}

	// check for .env file if exists
	overridenValues := loadEnv()
	if len(overridenValues) > 0 {
		for key, value := range overridenValues {
			config.Data[key] = value
			os.Setenv(key, value)
		}
	}

	return &config
}

func loadEnv() map[string]string {
	envPath := ".env"
	data, err := os.ReadFile(envPath)
	// will try to read .env file in current directory first, then traverse backwards until found
	for i := 0; i < 5; i++ {
		if err == nil {
			break
		}
		envPath = "../" + envPath
		data, err = os.ReadFile(envPath)
	}

	if err != nil {
		return nil
	}

	envMapData := map[string]string{}
	stringData := string(data)
	for _, line := range strings.Split(stringData, "\n") {
		line = strings.TrimSpace(line)

		// ignore empty lines and comments
		if len(line) == 0 || line[0] == '#' || line[0] == '/' || line[0] == ';' {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) >= 2 {
			key := parts[0]
			value := strings.Join(parts[1:], "=")
			if value[0] == '"' && value[len(value)-1] == '"' {
				value = value[1 : len(value)-1]
			}

			envMapData[key] = value
		}
	}

	return envMapData
}

// GetString will return string value from OS environment or default value
func (c *Config) GetString(key string, defaultValue ...string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		if value, ok := c.Data[key]; ok {
			return value
		}
	}

	if len(val) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return val
}

// GetInt will return int value from OS environment or default value
func (c *Config) GetInt(key string, defaultValue ...int) int {
	def := 0
	if len(defaultValue) > 0 {
		def = defaultValue[0]
	}

	val := c.GetString(key)
	if len(val) == 0 {
		return def
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return i
}

// GetFloat64 will return float64 value from OS environment or default value
func (c *Config) GetFloat64(key string, defaultValue ...float64) float64 {
	var def float64
	if len(defaultValue) > 0 {
		def = defaultValue[0]
	}

	val := c.GetString(key)
	if len(val) == 0 {
		return def
	}
	i, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return def
	}
	return i
}

// GetBool will return bool value from OS environment or default value
func (c *Config) GetBool(key string, defaultValue ...bool) bool {
	def := false
	if len(defaultValue) > 0 {
		def = defaultValue[0]
	}

	val := c.GetString(key)
	if len(val) == 0 {
		return def
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return def
	}
	return b
}
