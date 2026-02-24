package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	step      int
	folder    string
	column    string
	skills    string
	mode      string
	threshold string
	input     textinput.Model
}

func initialModel() model {
	ti := textinput.New()
	ti.Focus()
	return model{
		input: ti,
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			value := m.input.Value()
			if value == "" {
				// Don't advance if input is empty
				return m, nil
			}

			switch m.step {
			case 0:
				m.folder = value
				m.step++
				m.input.SetValue("")
				return m, nil
			case 1:
				m.column = value
				m.step++
				m.input.SetValue("")
				return m, nil
			case 2:
				m.skills = value
				m.step++
				m.input.SetValue("")
				return m, nil
			case 3:
				m.mode = value
				if value == "semantic" {
					m.step++
					m.input.SetValue("")
					return m, nil
				} else {
					// Skip threshold for non-semantic modes
					return m, tea.Quit
				}
			case 4:
				m.threshold = value
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	title := lipgloss.NewStyle().Bold(true).Render("SkillSearch\n\n")

	switch m.step {
	case 0:
		return title + "Enter Folder Path:\n" + m.input.View()
	case 1:
		return title + "Enter Column Header:\n" + m.input.View()
	case 2:
		return title + "Enter Skills (comma separated):\n" + m.input.View()
	case 3:
		return title + "Mode (exact/semantic):\n" + m.input.View()
	case 4:
		return title + "Threshold (if semantic):\n" + m.input.View()
	}
	return "Done"
}

func RunTUI() model {
	p := tea.NewProgram(initialModel())
	m, _ := p.Run()
	return m.(model)
}
