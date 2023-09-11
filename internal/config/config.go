package config

import "github.com/BurntSushi/toml"

type Config struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Token    string `toml:"token"`
}

func ParseTomlFile(path string) (Config, error) {
	config := Config{}
	_, err := toml.DecodeFile(path, &config)
	return config, err
}
