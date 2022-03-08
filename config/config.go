package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/spf13/viper"

	"github.com/edlorenzo/users-api/integration/redis"
	"github.com/edlorenzo/users-api/utils"
)

var (
	configPath string
	cfg        = &Config{}
)

const (
	ConfigPath = "CONFIG_PATH"
	Yaml       = "yaml"
)

func init() {
	flag.StringVar(&configPath, "config", "", "API config path")
}

type Config struct {
	RedisConfig  *redis.Conf `mapstructure:"redis"`
	GithubConfig *utils.Conf `mapstructure:"githubAPI"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getPwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.getPwd")
			}
			configPath = fmt.Sprintf("%s/config/config.yaml", getPwd)
		}
	}

	viper.SetConfigType(Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, nil
}
