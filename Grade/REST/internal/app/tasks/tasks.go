package tasks

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

var Tasks = []Task{
	{
		Id:     1,
		Text:   "task 1",
		Status: StatusInProgress,
	},
	{
		Id:     2,
		Text:   "task 2",
		Status: StatusCompleted,
	},
}
