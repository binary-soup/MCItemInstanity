package config

import (
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/util"
)

const PATH = "config.json"

func Load() (*Config, error) {
	path, _ := os.Executable()
	path = filepath.Dir(path)

	cfg, err := util.LoadJSON[Config]("config", filepath.Join(path, PATH))
	if err != nil {
		return nil, err
	}

	cfg.Root = path
	return cfg, nil
}
