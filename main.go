package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type TodoList struct {
	tasks    []string
	cursor   int // Which item our cursor is pointing at
	selected map[int]struct{}
}

// Returns the initial list of tasks
func initialModel() *TodoList {
	return &TodoList{
		tasks:    []string{"Task one", "Task two", "Task three"},
		selected: make(map[int]struct{}),
	}
}

func (model *TodoList) Init() tea.Cmd {
	return nil
}

func (model *TodoList) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.KeyMsg:
		switch message.String() {
		case "q":
			return model, tea.Quit
		case "up", "k":
			if model.cursor > 0 {
				model.cursor--
			}
		case "down", "j":
			if model.cursor < len(model.tasks)-1 {
				model.cursor++
			}
		case "enter", " ":
			_, ok := model.selected[model.cursor]

			if ok {
				delete(model.selected, model.cursor)
			} else {
				model.selected[model.cursor] = struct{}{}
			}
		}
	}

	return model, nil
}

func (model *TodoList) View() string {
	presentation := "What should we complete today\n\n"

	for i, task := range model.tasks {
		// Is the cursor pointing at task?
		cursor := " "

		if model.cursor == i {
			cursor = ">"
		}

		checkedTask := " "
		if _, ok := model.selected[i]; ok {
			checkedTask = "x"
		}

		presentation += fmt.Sprintf("%s [%s] %s\n", cursor, checkedTask, task)
	}

	presentation += "\nPress q to quit.\n"

	return presentation
}

func main() {

	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Printf("There was an error: %v", err)
		os.Exit(1)
	}
}
