package scaner

import "time"

type Event struct {
	title     string
	err       error
	CreatedAt time.Time
	result    string
}

func NewEvent(title, answer string, err error) Event {
	return Event{
		title:     title,
		err:       err,
		result:    answer,
		CreatedAt: time.Now(),
	}
}
