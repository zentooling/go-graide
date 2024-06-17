package config

import (
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/zentooling/graide/internal/logger"
)

var log = logger.New("config")
var Cfg *ConfigStruct = nil

var mutex = sync.Mutex{}

type ConfigStruct struct {
	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"4000"`
		Host string `yaml:"host" env:"SERVER_HOST" env-default:"0.0.0.0"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user" env:"DB_USERNAME"`
		Password string `yaml:"pass" env:"DB_PASSWORD"`
	} `yaml:"database"`
}

func Config() *ConfigStruct {
	return Cfg
}

func New(fname string) *ConfigStruct {
	mutex.Lock()
	defer mutex.Unlock()
	if Cfg == nil {
		log.Println("initializaing config global")
		Cfg = &ConfigStruct{}
	}
	err := cleanenv.ReadConfig(fname, Cfg)
	cwd, _ := os.Getwd()
	if err != nil {
		log.Fatalf("unable to process cfg file %s from dir %s, %+v", fname, cwd, err)
	}
	fullFname := cwd + fname
	help, _ := cleanenv.GetDescription(Cfg, &fullFname)
	log.Println(help)

	return Cfg
}
