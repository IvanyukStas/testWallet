package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct{
	Env string	`yaml:"env" env-default:"dev"`
				
	HTTPServer	`yaml:"server"`
	Database
}

type HTTPServer struct{
	Host string `yaml:"host" env-default:"0.0.0.0:8080"`
	Port string `yaml:"port"`
	Timeout `yaml:"timeout"`
}

type Timeout struct{
	Server time.Duration`yaml: "server" env-default:"3s"`
	Idle time.Duration	`yaml:"idle" env-default: "60s"`
}



func MustLoad() *Config{
	const cfgPath = "configs/main.yaml"

	if _, err := os.Stat(cfgPath); err != nil {
		log.Fatal("error load config file: %s", err)
	}

	var cfg Config
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		log.Fatal("error reading config file: %s", err)
		
	}
	return &cfg
}