package gocfg

import (
	"fmt"
	"os"
	"reflect"
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
	data := c.data[key]
	if i, ok := data.(string); ok {
		return strconv.Atoi(i)
	}
	if i, ok := data.(int); ok {
		return i, nil
	}
	if i, ok := data.(float64); ok {
		return int(i), nil
	}
	return 0, fmt.Errorf("%v can not convert to int type, is %v", data, reflect.TypeOf(data))
}

func (c *Config) GetStringList(key string) []string {
	data := c.data[key]
	if str, ok := data.(string); ok {
		return strings.Split(str, ",")
	}

	if list, ok := data.([]interface{}); ok {
		strList := make([]string, 0)
		for _, obj := range list {
			if str, ok := obj.(string); ok {
				strList = append(strList, str)
			}
		}
		return strList
	}

	return nil
}

func (c *Config) GetBool(key string) bool {
	data := c.GetString(key)
	switch data {
	case "true", "True", "TRUE", "yes", "Yes", "YES", "ok", "Ok", "OK":
		return true
	default:
		return false
	}
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
