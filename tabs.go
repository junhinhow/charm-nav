// Pacote nav fornece componentes de navegacao TUI estilizados com lipgloss.
package nav

import (
	"strings"

	"charm.land/lipgloss/v2"
)

// Tabs representa uma barra de abas horizontal para navegacao.
type Tabs struct {
	items         []string
	active        int
	Separator     string
	ActiveStyle   lipgloss.Style
	InactiveStyle lipgloss.Style
}

// NewTabs cria uma nova barra de abas com os itens fornecidos.
func NewTabs(items ...string) *Tabs {
	return &Tabs{
		items:     cloneItems(items),
		active:    0,
		Separator: " | ",
		ActiveStyle: lipgloss.NewStyle().
			Bold(true).
			Underline(true).
			Foreground(lipgloss.Color("#7D56F4")),
		InactiveStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")),
	}
}

// SetActive define o indice da aba ativa.
// Valores fora do intervalo sao ignorados.
func (t *Tabs) SetActive(index int) {
	if index >= 0 && index < len(t.items) {
		t.active = index
	}
}

// Active retorna o indice da aba ativa.
func (t *Tabs) Active() int {
	return t.active
}

// Next avanca para a proxima aba, com wraparound.
func (t *Tabs) Next() {
	if len(t.items) == 0 {
		return
	}
	t.active = (t.active + 1) % len(t.items)
}

// Previous volta para a aba anterior, com wraparound.
func (t *Tabs) Previous() {
	if len(t.items) == 0 {
		return
	}
	t.active = (t.active - 1 + len(t.items)) % len(t.items)
}

// Items retorna a lista de itens das abas.
// O slice devolvido e uma copia: alterar-lo nao muda o estado interno.
func (t *Tabs) Items() []string {
	return cloneItems(t.items)
}

// Add adiciona itens ao final da lista de abas e retorna o indice do
// primeiro item adicionado (-1 se nenhum item for adicionado).
func (t *Tabs) Add(items ...string) int {
	if len(items) == 0 {
		return -1
	}
	first := len(t.items)
	t.items = append(t.items, items...)
	return first
}

// Remove remove a aba no indice informado e retorna o seu rotulo.
// Retorna string vazia quando o indice e invalido.
//
// A aba ativa e ajustada: remover uma aba anterior a ativa mantem a mesma
// aba ativa e remover a propria aba ativa move a selecao para a proxima.
func (t *Tabs) Remove(index int) string {
	idx := clampIndex(index, len(t.items))
	if idx < 0 {
		return ""
	}

	removed := t.items[idx]
	t.items = append(t.items[:idx], t.items[idx+1:]...)

	switch {
	case idx < t.active:
		t.active--
	case t.active >= len(t.items) && len(t.items) > 0:
		t.active = len(t.items) - 1
	}

	return removed
}

// Render retorna a barra de abas estilizada como string.
func (t *Tabs) Render() string {
	if len(t.items) == 0 {
		return ""
	}

	var parts []string
	for i, item := range t.items {
		if i == t.active {
			parts = append(parts, t.ActiveStyle.Render(item))
		} else {
			parts = append(parts, t.InactiveStyle.Render(item))
		}
	}

	return strings.Join(parts, t.Separator)
}
