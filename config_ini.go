package gocfg

import (
	"bufio"
	"io"
	"strings"
)

func ParseIni(path string) (*Config, error) {
	file, err := openFile(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data := make(map[string]interface{}, 0)
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		str := strings.TrimSpace(string(line))
		if len(str) < 1 {
			continue
		}
		if str[0] == '[' || str[0] == ';' || !strings.ContainsRune(str, '=') {
			continue
		}
		kv := strings.Split(str, "=")
		if len(kv) != 2 {
			continue
		}
		data[kv[0]] = kv[1]
	}

	cfg := &Config{
		data: data,
	}

	return cfg, nil
}
