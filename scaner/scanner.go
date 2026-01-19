package scaner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todolist/list"
)

type Scanner struct {
	storageList *list.StorageTask
	events      []Event
}

func NewScanner(list *list.StorageTask) Scanner {

	return Scanner{
		storageList: list,
	}
}

func (s *Scanner) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPromt()

		ok := scanner.Scan()
		if !ok {
			return
		}

		inputString := scanner.Text()
		fields := strings.Fields(inputString)

		err, result := s.proccess(inputString)

		s.EventSave(fields, s.Str(fields), result, err) // fields[0] - команда, fields[1] - описание команды

		if result != "" {
			if result == "exit" {
				PrintExtit()
				return
			}

		}
	}

}
func (s *Scanner) proccess(inputString string) (error, string) {
	fields := strings.Fields(inputString)

	if len(fields) == 0 {
		return fmt.Errorf("Пустая строка"), ""
	}

	cmd := fields[0]
	strings.Fields(cmd)

	// exit
	if cmd == "exit" {
		return nil, cmd
	}

	// add
	if cmd == "add" {
		err := s.Add(fields)
		if err != nil {
			return err, ""
		}
		return nil, "Успешно была добавлена задача"

	}

	// list
	if cmd == "list" {
		err := s.ListInfo(fields)
		if err != nil {
			PrintErr(err)
			return err, ""
		}
		return nil, "Успешно была выполннена задача list"
	}

	// del
	if cmd == "del" {
		err := s.DeleteTask(fields)
		if err != nil {
			PrintErr(err)
			return err, ""
		}
		return nil, "Успешно была выполнена задача del"
	}

	// done
	if cmd == "done" {
		err := s.Done(fields)
		if err != nil {
			PrintErr(err)
			return err, ""
		}
		return nil, "Успешно была выполнена задача done"
	}

	//events
	if cmd == "events" {
		err := s.Event(fields)
		if err != nil {
			return err, ""
		}
		return nil, "Успешно была выполнена задача events"
	}
	//help
	if cmd == "help" {
		err := s.Help(fields)
		if err != nil {
			return err, ""
		}
		return nil, "Успешно была выполнена задача help"
	}

	PrintNN()
	return fmt.Errorf("некорректная команда"), ""
}

//

func (s *Scanner) Str(fields []string) string {
	if len(fields) == 0 {
		return ""
	}
	str := ""
	for i := 0; i < len(fields); i++ {
		str += fields[i]
		if i < len(fields)-1 {
			str += " "
		}
	}
	return str

}

// добавляем задачу
func (s *Scanner) Add(fields []string) error {
	if len(fields) < 3 {
		return fmt.Errorf("Не достаточно входных данных")
	}
	title := fields[1]

	description := ""

	for i := 2; i < len(fields); i++ {

		description += fields[i]

		if i != len(fields)-1 {
			description += " "
		}
	}

	task := list.NewTask(title, description)
	s.storageList.AddTask(task)
	PrintAdd(title)

	return nil
}

// list
func (s Scanner) ListInfo(fields []string) error {
	if len(fields) != 1 {
		return fmt.Errorf("Не достаточно данных")
	}
	tasks := s.storageList.ListTasks()
	PrintListTasks(tasks)
	return nil
}

// del
func (s *Scanner) DeleteTask(fields []string) error {

	if len(fields) != 2 {
		return fmt.Errorf("Некорректное количество данных")
	}

	description := fields[1]
	err := s.storageList.Delete(description)
	if err != nil {
		return fmt.Errorf("Не существует такой задачи")
	}
	PrintDelete()

	return nil
}

// done
func (s *Scanner) Done(fields []string) error {
	if len(fields) != 2 {
		return fmt.Errorf("Некорректное количество данных")
	}
	title := fields[1]
	s.storageList.Donetask(title)
	PrintIsDone()

	return nil
}

// eventSave

func (s *Scanner) EventSave(fields []string, title, answer string, err error) (error, []Event) {

	event := NewEvent(title, answer, err)
	s.events = append(s.events, event)

	return nil, s.events
}

// events
func (s *Scanner) Event(fields []string) error {
	if len(fields) != 1 {
		return fmt.Errorf("Неккоректное количество данных")
	}
	PrintEvent(s.events)
	return nil
}

// help
func (s *Scanner) Help(fields []string) error {
	if len(fields) != 1 {
		return fmt.Errorf("Неккоректное количество данных")
	}
	PrintHelp()
	return nil
}
