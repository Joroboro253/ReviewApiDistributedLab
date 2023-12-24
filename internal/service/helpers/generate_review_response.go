package helpers

import (
	"encoding/json"
	"net/http"
	"review_api/internal/data"
)

func GenerateReviewResponse(w http.ResponseWriter, review data.Review) {
	// Устанавливаем заголовок Content-Type для ответа
	w.Header().Set("Content-Type", "application/json")

	// Создаем структуру для ответа
	response := struct {
		Data data.Review `json:"data"`
	}{
		Data: review,
	}

	// Кодируем структуру в JSON и отправляем в ответе
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// В случае ошибки отправляем статус 500 Internal Server Error
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Устанавливаем HTTP-статус 200 OK
	w.WriteHeader(http.StatusOK)
}
