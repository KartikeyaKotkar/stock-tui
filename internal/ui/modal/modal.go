package modal

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var overlayStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#7D56F4")).
	Padding(1, 2).
	Background(lipgloss.Color("#1a1a2e"))

type Model struct {
	title   string
	content string
	width   int
	height  int
	visible bool
}

func New(title string) Model {
	return Model{title: title}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "esc" || msg.String() == "q" || msg.String() == "?" {
			m.visible = false
		}
	}
	return m, nil
}

func (m *Model) SetContent(content string) {
	m.content = content
}

func (m *Model) SetSize(w, h int) {
	m.width = w
	m.height = h
}

func (m *Model) Show() {
	m.visible = true
}

func (m *Model) Hide() {
	m.visible = false
}

func (m *Model) Toggle() {
	m.visible = !m.visible
}

func (m Model) Visible() bool {
	return m.visible
}

func (m Model) View() string {
	if !m.visible {
		return ""
	}

	modalWidth := min(m.width-10, 60)
	modalHeight := min(m.height-6, 20)

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).
		MarginBottom(1)

	content := titleStyle.Render(m.title) + "\n\n" + m.content

	modal := overlayStyle.
		Width(modalWidth).
		Height(modalHeight).
		Render(content)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, modal)
}

func Overlay(base, modal string, w, h int) string {
	if modal == "" {
		return base
	}

	baseLines := strings.Split(base, "\n")
	modalLines := strings.Split(modal, "\n")

	startY := (h - len(modalLines)) / 2
	startX := (w - lipgloss.Width(modalLines[0])) / 2

	for i, line := range modalLines {
		y := startY + i
		if y >= 0 && y < len(baseLines) {
			baseRunes := []rune(baseLines[y])
			lineRunes := []rune(line)

			for j, r := range lineRunes {
				x := startX + j
				if x >= 0 && x < len(baseRunes) {
					baseRunes[x] = r
				} else if x >= len(baseRunes) {
					baseRunes = append(baseRunes, r)
				}
			}
			baseLines[y] = string(baseRunes)
		}
	}

	return strings.Join(baseLines, "\n")
}
