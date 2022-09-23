package tool

import (
	"fmt"
	"github.com/spf13/viper"
)

//读取本地config目录下的yml配置文件
func UnmarshalConfig(config interface{}, configName string) error {
	viper.SetConfigName(configName) // 设置配置文件名为configName
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return fmt.Errorf("no such config file: %v\n", err)
		} else {
			// Config file was found but another error was produced
			return fmt.Errorf("read config error: %v\n", err)
		}
	}

	if err := viper.Unmarshal(config); err != nil {
		return err
	}
	return nil
}
