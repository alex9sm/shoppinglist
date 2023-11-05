package main 

import (
	"fmt"
	"os"
	"github.com/charmbracelet/bubbletea"
)

type model struct {
	choices []string //items on the list
	cursor int // which item our cursor is pointing at
	selected map[int]struct{} // which items are selected
}

//list model
func listModel() model {
	return model{
		choices: []string{"buy chicken", "buy milk", "buy eggs"},
		selected: make(map[int]struct{}),
	}
}

//init bubbletea
func (m model) Init() tea.Cmd {
	return nil
}

//update function (takes input)
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "up", "w":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "s":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

//cursor
func (m model) View() string {
	s := "what should we buy at the store\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress q to quit. \n"
	return s
}

func main() {
	p := tea.NewProgram(listModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}