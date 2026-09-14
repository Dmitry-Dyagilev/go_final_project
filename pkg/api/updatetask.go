package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Dmitry-Dyagilev/final-project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Не указан идентификатор"})
		return
	}
	result, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Задача не найдена"})
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &task); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if task.Title == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "не указан заголовок задачи"})
		return
	}

	if err := checkDate(&task); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	err = db.UpdateTask(&task)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Задача не найдена"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{})
}
