package nav

import (
	"strings"

	"charm.land/lipgloss/v2"
)

// Breadcrumb representa uma trilha de navegacao (breadcrumb).
type Breadcrumb struct {
	items     []string
	Separator string
	ActiveStyle   lipgloss.Style
	InactiveStyle lipgloss.Style
}

// NewBreadcrumb cria um novo breadcrumb com os itens fornecidos.
func NewBreadcrumb(items ...string) *Breadcrumb {
	return &Breadcrumb{
		items:     items,
		Separator: " > ",
		ActiveStyle: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")),
		InactiveStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")),
	}
}

// Push adiciona um item ao final do breadcrumb.
func (b *Breadcrumb) Push(item string) {
	b.items = append(b.items, item)
}

// Pop remove e retorna o ultimo item do breadcrumb.
// Retorna string vazia se nao houver itens.
func (b *Breadcrumb) Pop() string {
	if len(b.items) == 0 {
		return ""
	}
	last := b.items[len(b.items)-1]
	b.items = b.items[:len(b.items)-1]
	return last
}

// Items retorna a lista de itens do breadcrumb.
func (b *Breadcrumb) Items() []string {
	return b.items
}

// Render retorna o breadcrumb estilizado como string.
// O ultimo item e exibido com estilo ativo (item atual).
func (b *Breadcrumb) Render() string {
	if len(b.items) == 0 {
		return ""
	}

	var parts []string
	last := len(b.items) - 1
	for i, item := range b.items {
		if i == last {
			parts = append(parts, b.ActiveStyle.Render(item))
		} else {
			parts = append(parts, b.InactiveStyle.Render(item))
		}
	}

	return strings.Join(parts, b.Separator)
}
