package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/arvph/ServerAndDB/internal/app/api"
)

var (
	configPath string = "configs/api.toml"
)

func init() {
	// Определяем флаг "-path" с помощью flag.StringVar()
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file .toml format")
}

func main() {
	// Анализирует командную строку, извлекает значение флага
	flag.Parse()
	log.Println("It works!")
	// Создание переменной config типа api.Config с начальными данными
	config := api.NewConfig()

	// читаем из .toml/.env, т.к. там может быть новая информация
	// записываем в config.Storage данные из файла
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Cannot find configs file. using default values:", err)
	}
	server := api.New(config)

	// api server start
	log.Fatal(server.Start())
}
