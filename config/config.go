package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Server struct {
		Mode string
		Port string
	}

	DB map[string]string

	Log struct {
		Level   string
		Fpath   string
		Msize   int
		Mage    int
		Mbackup int
	}
}

func GetConfig(fpath string) (*Config, error) {
	cfg := new(Config)

	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := toml.NewDecoder(file).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// References
// Class material: lecture 12
