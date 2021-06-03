package configure

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

type Config struct {
	*viper.Viper
}

var config *Config

func New() *Config {
	if config == nil {
		config = &Config{
			viper.New(),
		}
		config.SetConfigName("main")
		config.AddConfigPath("./config")
		config.AddConfigPath(".")
		err := config.ReadInConfig() // Find and read the config file
		if err != nil {              // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	return config
}

// 加载多个配置文件,后面的会覆盖前面的
func (c *Config) Load(otherConfigFile ...string) {
	if len(otherConfigFile) > 0 {
		for _, configFile := range otherConfigFile {
			path, name := filepath.Split(configFile)
			if name != "" {
				c.SetConfigName(name)
			}
			if path != "" {
				c.AddConfigPath(path)
			}
			if err := c.MergeInConfig(); err != nil {
				panic(err)
			}
		}
	}
}
