package config

import "path/filepath"

type Config struct {
	Root              string
	MinecraftDataPath string `json:"minecraft_data_path"`
}

func (cfg *Config) JoinRoot(path string) string {
	return filepath.Join(cfg.Root, path)
}
