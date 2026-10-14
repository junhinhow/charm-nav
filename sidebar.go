package nav

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

// Sidebar representa um menu lateral vertical.
type Sidebar struct {
	items           []string
	selected        int
	Width           int
	Cursor          string
	SelectedStyle   lipgloss.Style
	UnselectedStyle lipgloss.Style
}

// NewSidebar cria um novo menu lateral com os itens fornecidos.
func NewSidebar(items ...string) *Sidebar {
	return &Sidebar{
		items:    cloneItems(items),
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
// O slice devolvido e uma copia: alterar-lo nao muda o estado interno.
func (s *Sidebar) Items() []string {
	return cloneItems(s.items)
}

// Render retorna o menu lateral estilizado como string.
//
// O cursor e medido pela largura de exibicao (e nao em bytes), os itens que
// nao cabem sao truncados em vez de quebrar linha e o destaque do item
// selecionado cobre a linha inteira.
func (s *Sidebar) Render() string {
	if len(s.items) == 0 {
		return ""
	}

	// Padding lateral de 1 coluna em cada lado.
	width := max(s.Width, 1)
	avail := max(width-2, 0)

	// O cursor ocupa a mesma quantidade de colunas em todos os itens, mesmo
	// quando ele e multibyte (ex: "▸ " tem 4 bytes e 2 colunas).
	cursorWidth := min(lipgloss.Width(s.Cursor), avail)
	prefixWidth := cursorWidth
	textWidth := max(avail-cursorWidth, 0)

	// Prefixo do item selecionado, ja limitado a largura disponivel.
	selectedPrefix := ansi.Truncate(s.Cursor, prefixWidth, "")
	unselectedPrefix := strings.Repeat(" ", prefixWidth)

	var lines []string
	for i, item := range s.items {
		prefix := unselectedPrefix
		style := s.UnselectedStyle
		if i == s.selected {
			prefix = selectedPrefix
			style = s.SelectedStyle
		}

		text := ansi.Truncate(item, textWidth, "")

		// O padding lateral esquerdo entra dentro do trecho estilizado para
		// que o fundo do item selecionado cubra a linha inteira.
		line := " " + prefix + text
		if faltam := width - lipgloss.Width(line); faltam > 0 {
			line += strings.Repeat(" ", faltam)
		}

		lines = append(lines, style.Render(line))
	}

	return strings.Join(lines, "\n")
}
