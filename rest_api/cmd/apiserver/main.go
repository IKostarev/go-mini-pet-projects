package main

import (
	"flag"
	"log"
	"rest_api/internal/api/apiserver"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "путь до файла конфига")
}

func main() {
	flag.Parse()



	cfg := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		log.Fatal("Ошибка при чтении файла конфига: ", err)
	}

	s := apiserver.New(cfg)
	if err := s.Start(); err != nil {
		log.Fatal("Ошибка при старте сервера: ", err)
	}
}