package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"time"

	"github.com/Dmitry-Dyagilev/final-project/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
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
	res, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"id": strconv.FormatInt(res, 10)})
}

// Проверка на корректность полученного значения task.Date
func checkDate(task *db.Task) error {
	now := time.Now()
	if len(task.Date) == 0 {
		task.Date = now.Format(format)
		return nil
	}

	t, err := time.Parse(format, task.Date)
	if err != nil {
		return err
	}

	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(format)

		} else {
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = next
		}
	}
	return nil
}
