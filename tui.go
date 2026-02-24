// Package main provides a colorful terminal user interface (TUI) for collecting
// search parameters from the user. Built with Bubble Tea framework for
// an interactive, step-by-step configuration experience.
package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// model represents the TUI state and user inputs during the configuration process.
// It follows the Bubble Tea pattern with Update/View methods.
type model struct {
	step      int             // Current input step (0-4)
	folder    string          // Excel files directory path
	column    string          // Column name to search within
	skills    string          // User's skill query (comma-separated)
	mode      string          // Search mode: "exact" or "exact + semantic"
	threshold string          // Similarity threshold for semantic search
	input     textinput.Model // Bubble Tea text input component
}

// initialModel creates a new TUI model with styled text input.
// Sets up the color scheme using a modern dark theme with bright accents.
func initialModel() model {
	ti := textinput.New()
	ti.Focus()
	// Style the input field with modern colors
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF79C6"))  // Pink prompt
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2"))    // Off-white text
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#50FA7B")) // Green cursor
	return model{
		input: ti,
	}
}

// Init satisfies the Bubble Tea Model interface. No initial commands needed.
func (m model) Init() tea.Cmd { return nil }

// Update handles user input and state transitions for the TUI.
// Processes keyboard events and advances through configuration steps.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle Ctrl+C for graceful exit
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		// Process Enter key - advance to next step or complete
		if msg.String() == "enter" {
			value := m.input.Value()
			// Prevent advancing with empty input
			if value == "" {
				return m, nil
			}

			// Handle each configuration step
			switch m.step {
			case 0: // Folder path input
				m.folder = value
				m.step++
				m.input.SetValue("")
				return m, nil
			case 1: // Column header input
				m.column = value
				m.step++
				m.input.SetValue("")
				return m, nil
			case 2: // Skills input
				m.skills = value
				m.step++
				m.input.SetValue("")
				return m, nil
			case 3: // Mode selection (1 or 2)
				switch value {
				case "1":
					m.mode = "exact"
					// Skip threshold for exact mode - exit TUI
					return m, tea.Quit
				case "2":
					m.mode = "exact + semantic"
					// Continue to threshold input
					m.step++
					m.input.SetValue("")
					return m, nil
				default:
					// Invalid input, don't advance
					return m, nil
				}
			case 4: // Threshold input (final step)
				m.threshold = value
				// Configuration complete - exit TUI
				return m, tea.Quit
			}
		}
	}

	// Update the text input component and pass through any commands
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

// View renders the current TUI state with colorful styling and emojis.
// Each step shows a different prompt with consistent color theming.
func (m model) View() string {
	// Define color styles for consistent theming
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF6B6B")). // Red title text
		Background(lipgloss.Color("#1A1A1A")). // Dark background
		Padding(0, 2).
		MarginBottom(1)

	promptStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#4ECDC4")) // Teal for section headers

	modeStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#45B7D1")). // Blue for mode header
		Bold(true)

	optionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#96CEB4")). // Green for options
		Bold(false)

	inputPromptStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFEAA7")). // Yellow for input prompts
		Bold(true)

	// Render title with search emoji
	title := titleStyle.Render("🔍 SkillSearch") + "\n\n"

	// Render appropriate step based on current configuration state
	switch m.step {
	case 0: // Step 1: Folder path input
		return title + promptStyle.Render("📁 Enter Folder Path:") + "\n" + m.input.View()
	case 1: // Step 2: Column header input
		return title + promptStyle.Render("📊 Enter Column Header:") + "\n" + m.input.View()
	case 2: // Step 3: Skills input
		return title + promptStyle.Render("🎯 Enter Skills (comma separated):") + "\n" + m.input.View()
	case 3: // Step 4: Mode selection with numbered options
		return title +
			modeStyle.Render("⚙️  Mode:") + "\n" +
			"  " + optionStyle.Render("1) exact") + "\n" +
			"  " + optionStyle.Render("2) exact + semantic") + "\n\n" +
			inputPromptStyle.Render("Enter your choice: ") + m.input.View()
	case 4: // Step 5: Threshold input (only for semantic mode)
		return title + promptStyle.Render("🎚️  Threshold (for semantic part):") + "\n" + m.input.View()
	}
	// Completion state
	return titleStyle.Render("✅ Done")
}

// RunTUI starts the Bubble Tea program and returns the configured model.
// This is the main entry point for the terminal user interface.
func RunTUI() model {
	p := tea.NewProgram(initialModel())
	m, _ := p.Run()
	return m.(model)
}
