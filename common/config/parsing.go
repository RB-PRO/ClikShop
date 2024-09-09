package config

import "github.com/jinzhu/configor"

func ParseConfig(fileNameConfig string) (config Config, err error) {
	err = configor.Load(&config, fileNameConfig)
	return config, err
}
