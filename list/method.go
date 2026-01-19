package list

import (
	"fmt"
)

type StorageTask struct {
	tasks map[string]Task
}

func NewStorageTask() *StorageTask {
	return &StorageTask{
		tasks: make(map[string]Task),
	}
}

// add task
func (s *StorageTask) AddTask(task Task) {
	if _, ok := s.tasks[task.Title]; ok {
		fmt.Println("уже есть существующая задача с таким Title")
		return
	}

	s.tasks[task.Title] = task
	fmt.Println("Успешно сохранили задачу")

}

// show all list tasks
func (s *StorageTask) ListTasks() map[string]Task {
	return s.tasks
}

// done task
func (s *StorageTask) Donetask(title string) error {

	task, ok := s.tasks[title]
	if !ok {
		return fmt.Errorf("Такой задачи не существует!!!")
	}

	task.Done()
	s.tasks[title] = task

	return nil

}

// delete tas
func (s *StorageTask) Delete(title string) error {

	_, ok := s.tasks[title]
	if !ok {
		return fmt.Errorf("Такой задачи не существует!!!")
	}

	delete(s.tasks, title)
	return nil
}
