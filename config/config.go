package config

import "path/filepath"

type Config struct {
	Root          string
	MinecraftData string `json:"minecraft_data"`
	StaticData    string `json:"static_data"`
	Datapack      string `json:"datapack"`
}

func (cfg Config) JoinRoot(path ...string) string {
	return filepath.Join(cfg.Root, filepath.Join(path...))
}

func (cfg Config) StaticDataPath() string {
	return cfg.JoinRoot(cfg.StaticData)
}

func (cfg Config) JoinStaticData(path ...string) string {
	return cfg.JoinRoot(cfg.StaticData, filepath.Join(path...))
}

func (cfg Config) DatapackPath() string {
	return cfg.JoinRoot(cfg.Datapack)
}

func (cfg Config) JoinDatapack(path ...string) string {
	return cfg.JoinRoot(cfg.Datapack, filepath.Join(path...))
}
