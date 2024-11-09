package rest


import (
	"encoding/json"
	"net/http"
	"strconv"
	"REST/internal/app/tasks"
)

func GetTasks(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks.Tasks); err != nil {
		http.Error(w, "JSON Encoder error", http.StatusInternalServerError)
		return
	}
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var task tasks.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "JSON Decoder error", http.StatusInternalServerError)
		return
	}
	task.Id = len(tasks.Tasks) + 1
	tasks.Tasks = append(tasks.Tasks, task)
	w.WriteHeader(http.StatusOK)
}

func ChangeTask(w http.ResponseWriter, r *http.Request) {
	var task tasks.Task

	id, err := strconv.Atoi(r.URL.Path[len("/tasks/"):])
	if err != nil {
		http.Error(w, "id should be int", http.StatusBadRequest)
		return
	}
	
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "JSON Decoder error", http.StatusInternalServerError)
		return
	}
	task.Id = id
	
	found := false
	for i, anotherTask := range tasks.Tasks {
		if anotherTask.Id == id {
			tasks.Tasks[i] = task
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "id not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Path[len("/tasks/"):])
	if err != nil {
		http.Error(w, "id should be int", http.StatusBadRequest)
		return
	}
		
	found := false
	for i, anotherTask := range tasks.Tasks {
		if anotherTask.Id == id {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "id not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}