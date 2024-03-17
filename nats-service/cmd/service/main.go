package main

import (
	"WB_Tech_level_0/internal/app"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	configPath = "configs/app.yaml"
)

func main() {
	config := app.NewConfig()            //Инициализация конфига
	file, err := os.ReadFile(configPath) //Открытие конфиг файла

	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(file, &config); err != nil { //Считывание конфига
		log.Fatal(err)
	}

	s := app.New(config)              //Инициализация сервиса
	if err := s.Start(); err != nil { //Запуск сервиса
		log.Fatal(err)
	}
}
