package server

import (
	"log"
	"noteproject/envs"
)

func InitServer() {
	// Инициализация внешних значений ENV
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		// Вывод сообщения об ошибке
		log.Fatal("Ошибка инициализации ENV: ", errEnvs)
	} else {
		log.Println("Инициализация ENV прошла успешно")
	}
	// Инициализация базы данных
}

func StartServer() {
	// Инициализация роутеров
	// Запуск сервера
	InitRotes()
}
