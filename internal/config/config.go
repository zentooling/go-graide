package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/zentooling/graide/internal/logger"
)

var log = logger.New("config")
var cfg *configStruct = nil

var mutex = sync.Mutex{}

type configStruct struct {
	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"4000"`
		Host string `yaml:"host" env:"SERVER_HOST" env-default:"0.0.0.0"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host" env:"DB_HOST"`
		Port     string `yaml:"port" env:"DB_PORT"`
		Database string `yaml:"database" env:"DB_DATABASE"`
		Username string `yaml:"user" env:"DB_USERNAME"`
		Password string `yaml:"pass" env:"DB_PASSWORD"`
	} `yaml:"database"`
	Template struct {
		RootDir string `yaml:"root_dir" env:"TEMPLATE_ROOT"`
	} `yaml:"template"`
}

func Instance() *configStruct {
	if cfg == nil {
		log.Fatalln("config not initiailized. config.New() must be called before config.Instance()")
	}
	return cfg
}

func New(fname string) *configStruct {
	mutex.Lock()
	defer mutex.Unlock()
	if cfg == nil {
		log.Println("initializaing config global")
		cfg = &configStruct{}
	}
	err := cleanenv.ReadConfig(fname, cfg)
	cwd, _ := os.Getwd()
	if err != nil {
		log.Fatalf("unable to process cfg file %s from dir %s, %+v", fname, cwd, err)
	}
	fullFname := filepath.Join(cwd, fname)
	help, _ := cleanenv.GetDescription(cfg, &fullFname)
	log.Println(help)

	return cfg
}
