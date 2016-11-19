package gocfg

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	data map[string]interface{}
}

func (c *Config) GetAllData() map[string]interface{} {
	return c.data
}

func (c *Config) GetString(key string) string {
	data := c.data[key]
	if str, ok := data.(string); ok {
		return str
	}
	return ""
}

func (c *Config) GetInt(key string) (int, error) {
	return strconv.Atoi(c.GetString(key))
}

func (c *Config) GetStringList(key string) []string {
	return strings.Split(c.GetString(key), ",")
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func openFile(path string) (*os.File, error) {
	if !pathExists(path) {
		return nil, fmt.Errorf("file not found in :%s", path)
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
