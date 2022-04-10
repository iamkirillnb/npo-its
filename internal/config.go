package internal

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	PgDb      DbConfig     `yaml:"pg_db"`
	ServerCfg ServerConfig `yaml:"server"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SslMode  string `yaml:"ssl_mode"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (s *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

func (d *DbConfig) DbUrlConnection() string {
	dbUrl := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s",
		d.User, d.Database, d.Password, d.Host, d.Port, d.SslMode)
	return dbUrl
}

var instance *Config
var once sync.Once

func GetConfig(path string) *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig(path, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return instance
}
