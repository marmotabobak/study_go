package rest


import (
	"encoding/json"
	"net/http"
	"strconv"
)

type TaskStatus string

const (
	StatusCompleted  TaskStatus = "completed"
	StatusInProgress TaskStatus = "in progress"
	StatusUnknown    TaskStatus = "unknown"
)

type Task struct {
	Id     int        `json:"id"`
	Text   string     `json:"text"`
	Status TaskStatus `json:"status"`
}

var tasks []Task

func GetTasks(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "JSON Encoder error", http.StatusInternalServerError)
		return
	}
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "JSON Decoder error", http.StatusInternalServerError)
		return
	}
	task.Id = len(tasks) + 1
	tasks = append(tasks, task)
	w.WriteHeader(http.StatusOK)
}

func ChangeTask(w http.ResponseWriter, r *http.Request) {
	var task Task

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
	for i, anotherTask := range tasks {
		if anotherTask.Id == id {
			tasks[i] = task
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
	for i, anotherTask := range tasks {
		if anotherTask.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
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