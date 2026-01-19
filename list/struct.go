package list

import "time"

type Task struct {
	Title       string
	Description string
	IsDone      bool

	CreatedAt time.Time
	DoneAt    *time.Time
}

func NewTask(title, description string) Task {
	return Task{
		Title:       title,
		Description: description,

		IsDone:    false,
		CreatedAt: time.Now(),
		DoneAt:    nil,
	}
}

func (t *Task) Done() {
	t.IsDone = true
	now := time.Now()
	t.DoneAt = &now
}
