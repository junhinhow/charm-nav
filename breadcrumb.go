package nav

import (
	"strings"

	"charm.land/lipgloss/v2"
)

// Breadcrumb representa uma trilha de navegacao (breadcrumb).
type Breadcrumb struct {
	items         []string
	Separator     string
	ActiveStyle   lipgloss.Style
	InactiveStyle lipgloss.Style
}

// NewBreadcrumb cria um novo breadcrumb com os itens fornecidos.
func NewBreadcrumb(items ...string) *Breadcrumb {
	return &Breadcrumb{
		items:     cloneItems(items),
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
// O slice devolvido e uma copia: alterar-lo nao muda o estado interno.
func (b *Breadcrumb) Items() []string {
	return cloneItems(b.items)
}

// Add adiciona itens ao final da trilha e retorna o indice do primeiro item
// adicionado (-1 se nenhum item for adicionado).
// Equivale a Push, mas aceita varios itens de uma vez.
func (b *Breadcrumb) Add(items ...string) int {
	if len(items) == 0 {
		return -1
	}
	first := len(b.items)
	b.items = append(b.items, items...)
	return first
}

// Remove remove o item no indice informado e retorna o seu rotulo.
// Retorna string vazia quando o indice e invalido.
func (b *Breadcrumb) Remove(index int) string {
	idx := clampIndex(index, len(b.items))
	if idx < 0 {
		return ""
	}

	removed := b.items[idx]
	b.items = append(b.items[:idx], b.items[idx+1:]...)
	return removed
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
