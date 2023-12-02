package config

import (
	"fmt"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Server ServerConfig
	Redis  RedisConfig
	Mysql  MysqlConfig
}

func TestConfigLoad(t *testing.T) {
	c := &Config{}
	conf.LoadConfig("./sectran-api.yaml", c)
	fmt.Println(c.Server.Address)
}
