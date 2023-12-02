package config

type ServerConfig struct {
	Address string
	Port    int
}

type MysqlConfig struct {
	Address  string
	Port     int
	Db       string
	User     string
	Password string
}

type RedisConfig struct {
	Address  string
	Port     int
	Password string
}
