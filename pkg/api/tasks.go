package api

import (
	"net/http"

	"github.com/Dmitry-Dyagilev/final-project/pkg/db"
)

type TaskResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	task, err := db.Tasks(50)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if task == nil {
		task = []*db.Task{}
	}
	writeJSON(w, http.StatusOK, TaskResp{
		Tasks: task,
	})
	return
}
