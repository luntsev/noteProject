package server

import (
	"github.com/gin-gonic/gin"
	"noteproject/handlers"
)

func InitRotes() {
	// Инициализация  роута (по умолчанию)
	router := gin.Default()
	// Создание заметки
	router.POST("/note/", handlers.CreateNoteHandler)
	// Редактирование заметки
	router.PUT("/note/:id", handlers.UpdateNoteHandler)
	// Удаление заметки
	router.DELETE("/note/:id", handlers.DeleteNoteHandler)
	// Получение заметки
	router.GET("/note/:id", handlers.GetNoteHandler)
	// Получение списка всех заметок
	router.GET("/notes", handlers.GetNotesHandler)

	// Запуск сервера на порту 9100
	router.Run(":9100")
}
