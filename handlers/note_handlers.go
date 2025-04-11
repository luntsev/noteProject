package handlers

import (
	"fmt"
	"net/http"
	"noteproject/database"
	"noteproject/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// Обработка запроса для создания заметки
func CreateNoteHandler(ctx *gin.Context) {
	// Создание новой заметки
	var note models.Note
	// Получаем данные из запроса
	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Получить уникальный id
	note.Id = uuid.New().String()
	// Тестовый ID автора
	note.AuthorID = 1

	// Получаем коллекцию "notes"
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", note.AuthorID))

	// Вставляем заметку в коллекцию
	_, errInsert := collection.InsertOne(ctx, note)
	if errInsert != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": errInsert.Error()})
	}
	// Если ошибок нет, то возвращаем заметку и статус 200
	ctx.JSON(http.StatusOK, gin.H{
		"note":    note,
		"message": "Заметка успешно создана"})
}

// Обработка запроса для получения заметки по ID
func GetNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetNoteHandler")
}

// Обработка запроса для получения всех заметок
func GetNotesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetNotesHandler")
}

// Обработка запроса для удаления заметки по ID
func DeleteNoteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", 1))
	filter := bson.M{"id": id}

	if result, err := collection.DeleteOne(ctx, filter); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if result.DeletedCount == 0 {
		ctx.JSON(http.StatusOK, "Заметка не найдена")
	} else {
		ctx.JSON(http.StatusOK, "Заметка успешно удалена")
	}
}

// Обработка запроса для редактирования заметки по ID
func UpdateNoteHandler(ctx *gin.Context) {
	authorId := 1

	ctx.JSON(http.StatusOK, "UpdateNoteHandler")
}
