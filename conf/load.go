package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// LoadConfigFromToml 从Toml格式的配置文件加载配置
func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()

	// 读取Toml格式的配置
	_, err := toml.DecodeFile(filePath, config)
	fmt.Println("config:", filePath)
	if err != nil {
		return fmt.Errorf("load config from %s failed: %s", filePath, err)
	}

	return nil
}

// LoadConfigFromEnv 从环境变量加载配置
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	err := env.Parse(config)
	if err != nil {
		return err
	}

	return nil
}

// loadGloabal 加载 db 全局实例
func loadGloabal() (err error) {
	db, err = config.MySQL.getDBConn()
	if err != nil {
		return
	}
	return
}
