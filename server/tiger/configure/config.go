package configure

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

type Config struct {
	*viper.Viper
}

// 加载配置，允许合并多个配置文件
func (c *Config) LoadConfig(configFile string, otherConfigFile ...string) {
	path, name := filepath.Split(configFile)
	if name != "" {
		c.SetConfigName(name)
	}
	if path != "" {
		c.AddConfigPath(path)
	}
	err := c.ReadInConfig()
	if err != nil {
		panic(err)
	}
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
	fmt.Println(c.AllSettings())
}
