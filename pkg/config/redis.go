package config

import (
	"fmt"
)

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

func (c *Redis) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}
