package nav

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

// Sidebar representa um menu lateral vertical.
type Sidebar struct {
	items    []string
	selected int
	Width    int
	Cursor   string
	SelectedStyle   lipgloss.Style
	UnselectedStyle lipgloss.Style
}

// NewSidebar cria um novo menu lateral com os itens fornecidos.
func NewSidebar(items ...string) *Sidebar {
	return &Sidebar{
		items:    items,
		selected: 0,
		Width:    20,
		Cursor:   "▸ ",
		SelectedStyle: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			Background(lipgloss.Color("#1A1A2E")),
		UnselectedStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")),
	}
}

// SetSelected define o indice do item selecionado.
// Valores fora do intervalo sao ignorados.
func (s *Sidebar) SetSelected(index int) {
	if index >= 0 && index < len(s.items) {
		s.selected = index
	}
}

// Selected retorna o indice do item selecionado.
func (s *Sidebar) Selected() int {
	return s.selected
}

// Next avanca para o proximo item, com wraparound.
func (s *Sidebar) Next() {
	if len(s.items) == 0 {
		return
	}
	s.selected = (s.selected + 1) % len(s.items)
}

// Previous volta para o item anterior, com wraparound.
func (s *Sidebar) Previous() {
	if len(s.items) == 0 {
		return
	}
	s.selected = (s.selected - 1 + len(s.items)) % len(s.items)
}

// Items retorna a lista de itens do sidebar.
func (s *Sidebar) Items() []string {
	return s.items
}

// Render retorna o menu lateral estilizado como string.
func (s *Sidebar) Render() string {
	if len(s.items) == 0 {
		return ""
	}

	containerStyle := lipgloss.NewStyle().
		Width(s.Width).
		PaddingLeft(1).
		PaddingRight(1)

	var lines []string
	for i, item := range s.items {
		var line string
		if i == s.selected {
			line = s.SelectedStyle.Render(fmt.Sprintf("%s%s", s.Cursor, item))
		} else {
			padding := strings.Repeat(" ", len(s.Cursor))
			line = s.UnselectedStyle.Render(fmt.Sprintf("%s%s", padding, item))
		}
		lines = append(lines, line)
	}

	return containerStyle.Render(strings.Join(lines, "\n"))
}
