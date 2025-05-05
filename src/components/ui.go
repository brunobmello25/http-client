package components

import (
	"fmt"
	"strings"

	"github.com/brunobmello25/http-client/src/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Styles
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#25A065")).
				Bold(true)

	normalItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5"))

	borderStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#25A065"))
)

type UI struct {
	collection *models.Collection
	cursor     int
	width      int
	height     int
	activePane int // 0: left, 1: center
}

func NewUI(collection *models.Collection) *UI {
	return &UI{
		collection: collection,
		cursor:     0,
		activePane: 0,
	}
}

func (ui *UI) Init() tea.Cmd {
	return nil
}

func (ui *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ui.width = msg.Width
		ui.height = msg.Height
		return ui, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return ui, tea.Quit
		case "tab":
			ui.activePane = (ui.activePane + 1) % 2
		case "up", "k":
			if ui.activePane == 0 && ui.cursor > 0 {
				ui.cursor--
			}
		case "down", "j":
			if ui.activePane == 0 && ui.cursor < len(ui.collection.Requests)-1 {
				ui.cursor++
			}
		case "enter":
			// TODO: Implement request execution
			return ui, nil
		}
	}
	return ui, nil
}

func (ui *UI) renderRequestList() string {
	var sb strings.Builder
	sb.WriteString(titleStyle.Render("Requests"))
	sb.WriteString("\n\n")

	for i, req := range ui.collection.Requests {
		cursor := " "
		if ui.cursor == i {
			cursor = ">"
		}

		style := normalItemStyle
		if ui.cursor == i {
			style = selectedItemStyle
		}

		sb.WriteString(fmt.Sprintf("%s %s %s\n", cursor, style.Render(req.Method), req.Name))
	}

	// Fill remaining height with empty lines
	remainingHeight := ui.height - strings.Count(sb.String(), "\n") - 3 // 3 for title, help text, and border
	for i := 0; i < remainingHeight; i++ {
		sb.WriteString("\n")
	}

	return borderStyle.Width(ui.width / 3).Height(ui.height - 2).Render(sb.String())
}

func (ui *UI) renderRequestDetails() string {
	if ui.cursor >= len(ui.collection.Requests) {
		return borderStyle.Width(ui.width * 2 / 3).Height(ui.height - 2).Render("No request selected")
	}

	req := ui.collection.Requests[ui.cursor]
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("Request Details"))
	sb.WriteString("\n\n")

	// Method and URL
	sb.WriteString(fmt.Sprintf("Method: %s\n", req.Method))
	sb.WriteString(fmt.Sprintf("URL: %s\n\n", req.URL))

	// Headers
	sb.WriteString("Headers:\n")
	for key, value := range req.Headers {
		sb.WriteString(fmt.Sprintf("  %s: %s\n", key, value))
	}
	sb.WriteString("\n")

	// Body
	if req.Body != "" {
		sb.WriteString("Body:\n")
		sb.WriteString(req.Body)
	}

	// Fill remaining height with empty lines
	remainingHeight := ui.height - strings.Count(sb.String(), "\n") - 3 // 3 for title, help text, and border
	for i := 0; i < remainingHeight; i++ {
		sb.WriteString("\n")
	}

	return borderStyle.Width(ui.width * 2 / 3).Height(ui.height - 2).Render(sb.String())
}

func (ui *UI) View() string {
	if ui.collection == nil {
		return "No collection loaded"
	}

	// Calculate layout
	leftPane := ui.renderRequestList()
	centerPane := ui.renderRequestDetails()

	// Combine panes
	layout := lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftPane,
		centerPane,
	)

	// Add help text
	helpText := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Render("Tab: Switch Panes • ↑/↓: Navigate • Enter: Execute • q: Quit")

	// Combine everything
	finalLayout := lipgloss.JoinVertical(
		lipgloss.Left,
		layout,
		helpText,
	)

	return finalLayout
}
