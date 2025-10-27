package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"gopkg.in/yaml.v3"
)

type SecretKeys struct {
	S_16 string `yaml:"_16"`
	S_24 string `yaml:"_24"`
	S_32 string `yaml:"_32"`
}

type Config struct {
	SecretKeys SecretKeys `yaml:"secret_keys"`
}

var (
	C    *Config
	once sync.Once
)

func load() (*Config, error) {
	var err error
	once.Do(func() {
		_, filename, _, _ := runtime.Caller(0)
		dir := filepath.Dir(filename)
		path := filepath.Join(dir, "config.yaml")

		file, e := os.ReadFile(path)
		if e != nil {
			err = fmt.Errorf("failed to read %s: %w", path, e)
			return
		}

		var c Config
		if e := yaml.Unmarshal(file, &c); e != nil {
			err = fmt.Errorf("failed to parse config.yaml: %w", e)
			return
		}

		C = &c
	})

	return C, err
}

func GetConfig() {
	_, err := load()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
}
