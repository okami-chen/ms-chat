package config

import (
	"fmt"
)

type Cache struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

func (c *Cache) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}
