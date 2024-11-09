package rest

// import (
// 	"bytes"
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"REST/internal/app/tasks"
// )

// func TestGetTask(t *testing.T) {

// 	w := httptest.NewRecorder()

// 	GetTasks(w)

// 	res := w.Result()

// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("200 expected got: %d", res.StatusCode)
// 	}

// 	tasksJSON, _ := json.Marshal(tasks)
// 	buf, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
// 	defer res.Body.Close()	

// 	if string(tasksJSON) + "\n" != string(buf) {
// 		t.Errorf("Expected: %s, got: %s", tasksJSON, buf)
// 	}
// }

// func TestAddTask(t *testing.T) {

// 	task := Task{
// 		Text: "foo",
// 		Status: StatusUnknown,
// 	}
// 	taskJSON, _ := json.Marshal(task)

// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(taskJSON))

// 	AddTask(w, r)

// 	res := w.Result()

// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("200 expected got: %d", res.StatusCode)
// 	}

// 	GetTasks(w)

// 	res = w.Result()

// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("200 expected got: %d", res.StatusCode)
// 	}

// 	tasksJSON, _ := json.Marshal(tasks)
// 	buf, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
// 	defer res.Body.Close()	

// 	if string(tasksJSON) + "\n" != string(buf) {
// 		t.Errorf("Expected: %s, got: %s", tasksJSON, buf)
// 	}
// }