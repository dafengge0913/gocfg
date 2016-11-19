package gocfg

import (
	"encoding/json"
	"io/ioutil"
)

func ParseJson(path string) (*Config, error) {
	file, err := openFile(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data := make(map[string]interface{}, 0)
	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return nil, err
	}

	cfg := &Config{
		data: data,
	}

	return cfg, nil
}
