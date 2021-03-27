package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	configFile = "config.toml"
	secretKey = "hnf27"
)

type Mysql struct {
	Ip       string `mapstructure:"ip"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string
}

func (m *Mysql) GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		m.Username,
		m.Password,
		m.Ip,
		m.Port,
		m.Database,
	)
}

type Redis struct {
	Ip   string `mapstructure:"ip"`
	Port int    `mapstructure:"port"`
}

func (r *Redis) GetDSN() string {
	return fmt.Sprintf("%s:%d", r.Ip, r.Port)
}

type Config struct {
	Project   string `mapstructure:"project"`
	SecretKey string
	Mysql     Mysql `mapstructure:"mysql"`
	Redis     Redis `mapstructure:"redis"`
}

func InitConfig() *Config {
	var c Config
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("[configure]read config failed: %v", err))
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Sprintf("[configure]parse config failed: %v", err))
	}
	fmt.Println("[configure]loaded config succeed.")
	c.SecretKey = secretKey
	return &c
}
