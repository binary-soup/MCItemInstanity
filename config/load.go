package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/util"
)

const PATH = "config.json"

func Load() (*Config, error) {
	path, _ := os.Executable()
	path = filepath.Dir(path)

	file, err := os.Open(filepath.Join(path, PATH))
	if err != nil {
		return nil, util.ChainError(err, "error opening config file")
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, util.ChainError(err, "error decoding json")
	}

	cfg.Root = path
	return &cfg, nil
}
