package nav

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

// larguraEstilizada devolve a largura de exibicao do trecho estilizado da
// linha, ou seja, o conteudo entre o codigo ANSI inicial e o reset final.
func larguraEstilizada(linha string) int {
	fim := strings.Index(linha, "\x1b[m")
	if fim < 0 {
		fim = strings.Index(linha, "\x1b[0m")
	}
	if fim < 0 {
		return ansi.StringWidth(linha)
	}
	return ansi.StringWidth(linha[:fim])
}

// colunaDoTexto devolve a coluna em que o texto do item comeca.
func colunaDoTexto(t *testing.T, linha, texto string) int {
	t.Helper()
	plano := ansi.Strip(linha)
	idx := strings.Index(plano, texto)
	if idx < 0 {
		t.Fatalf("texto %q nao encontrado na linha %q", texto, plano)
	}
	return ansi.StringWidth(plano[:idx])
}

func TestSidebarAlinhaColunasComCursorMultibyte(t *testing.T) {
	// O cursor padrao "▸ " tem 4 bytes e 2 colunas. A medicao por bytes
	// deslocava o item selecionado 2 colunas para a esquerda.
	sb := NewSidebar("Dashboard", "Usuarios", "Configuracoes")
	sb.SetSelected(1)

	linhas := strings.Split(sb.Render(), "\n")
	if len(linhas) != 3 {
		t.Fatalf("esperava 3 linhas, obteve %d", len(linhas))
	}

	colunas := []int{
		colunaDoTexto(t, linhas[0], "Dashboard"),
		colunaDoTexto(t, linhas[1], "Usuarios"),
		colunaDoTexto(t, linhas[2], "Configuracoes"),
	}
	for i, col := range colunas {
		if col != colunas[0] {
			t.Errorf("item %d comeca na coluna %d, esperado %d", i, col, colunas[0])
		}
	}
	if colunas[0] != 3 {
		t.Errorf("esperava texto na coluna 3 (1 de padding + 2 do cursor), obteve %d", colunas[0])
	}
}

func TestSidebarAlinhaColunasComCursorASCII(t *testing.T) {
	sb := NewSidebar("Dashboard", "Usuarios")
	sb.Cursor = "> "
	sb.SetSelected(0)

	linhas := strings.Split(sb.Render(), "\n")
	col0 := colunaDoTexto(t, linhas[0], "Dashboard")
	col1 := colunaDoTexto(t, linhas[1], "Usuarios")
	if col0 != col1 {
		t.Errorf("itens desalinhados: %d != %d", col0, col1)
	}
}

func TestSidebarRespeitaLarguraDoCursor(t *testing.T) {
	// Cursor de 3 colunas: o texto de todos os itens deve continuar
	// alinhado.
	sb := NewSidebar("Dashboard", "Usuarios")
	sb.Cursor = "=> "
	sb.SetSelected(1)

	linhas := strings.Split(sb.Render(), "\n")
	col0 := colunaDoTexto(t, linhas[0], "Dashboard")
	col1 := colunaDoTexto(t, linhas[1], "Usuarios")
	if col0 != col1 {
		t.Errorf("itens desalinhados: %d != %d", col0, col1)
	}
	if col0 != 4 {
		t.Errorf("esperava texto na coluna 4, obteve %d", col0)
	}
}

func TestSidebarTruncaItemLongo(t *testing.T) {
	longo := "Este item tem um rotulo comprido demais para o menu lateral"
	sb := NewSidebar(longo, "Usuarios")

	linhas := strings.Split(sb.Render(), "\n")
	if len(linhas) != 2 {
		t.Fatalf("esperava 2 linhas (sem quebra de linha), obteve %d: %q", len(linhas), sb.Render())
	}

	for i, linha := range linhas {
		if w := lipgloss.Width(linha); w > sb.Width {
			t.Errorf("linha %d tem largura %d, maior que Width %d", i, w, sb.Width)
		}
	}

	// O texto do item precisa ter sido realmente truncado.
	plano := ansi.Strip(linhas[0])
	if strings.Contains(plano, longo) {
		t.Errorf("esperava o item truncado, obteve %q", plano)
	}
	if !strings.Contains(plano, "Este item") {
		t.Errorf("esperava o inicio do rotulo preservado, obteve %q", plano)
	}
}

func TestSidebarTruncaItemLongoComCursorMultibyte(t *testing.T) {
	longo := "Item com acentuacao e emoji: configuracoes avancadas"
	sb := NewSidebar(longo)
	sb.Cursor = "▸ "

	linhas := strings.Split(sb.Render(), "\n")
	if len(linhas) != 1 {
		t.Fatalf("esperava 1 linha, obteve %d", len(linhas))
	}
	if w := lipgloss.Width(linhas[0]); w > sb.Width {
		t.Errorf("linha tem largura %d, maior que Width %d", w, sb.Width)
	}
}

func TestSidebarDestaqueCobreLinhaInteira(t *testing.T) {
	sb := NewSidebar("Dashboard", "Usuarios")
	sb.SetSelected(0)

	linhas := strings.Split(sb.Render(), "\n")
	selecionada := linhas[0]

	if w := larguraEstilizada(selecionada); w != sb.Width {
		t.Errorf("destaque cobre %d colunas, esperado %d", w, sb.Width)
	}
	if w := lipgloss.Width(selecionada); w != sb.Width {
		t.Errorf("linha tem largura %d, esperado %d", w, sb.Width)
	}
}

func TestSidebarLinhasTodasComMesmaLargura(t *testing.T) {
	sb := NewSidebar("A", "Item bem mais longo que o primeiro", "C")
	sb.SetSelected(1)

	for i, linha := range strings.Split(sb.Render(), "\n") {
		if w := lipgloss.Width(linha); w != sb.Width {
			t.Errorf("linha %d tem largura %d, esperado %d", i, w, sb.Width)
		}
	}
}

func TestItemsNaoExpoemSliceInterno(t *testing.T) {
	sb := NewSidebar("a", "b")
	items := sb.Items()
	items[0] = "ALTERADO"

	if got := sb.Items()[0]; got != "a" {
		t.Errorf("estado interno alterado via Items(): obteve %q", got)
	}

	tabs := NewTabs("a", "b")
	itemsTabs := tabs.Items()
	itemsTabs[0] = "ALTERADO"
	if got := tabs.Items()[0]; got != "a" {
		t.Errorf("estado interno alterado via Items(): obteve %q", got)
	}

	bc := NewBreadcrumb("a", "b")
	itemsBc := bc.Items()
	itemsBc[0] = "ALTERADO"
	if got := bc.Items()[0]; got != "a" {
		t.Errorf("estado interno alterado via Items(): obteve %q", got)
	}
}

func TestConstrutoresCopiaSliceVariadico(t *testing.T) {
	baseSidebar := []string{"a", "b"}
	sb := NewSidebar(baseSidebar...)
	baseSidebar[0] = "ALTERADO"
	if got := sb.Items()[0]; got != "a" {
		t.Errorf("sidebar guardou referencia do slice original: obteve %q", got)
	}

	baseTabs := []string{"a", "b"}
	tabs := NewTabs(baseTabs...)
	baseTabs[1] = "ALTERADO"
	if got := tabs.Items()[1]; got != "b" {
		t.Errorf("tabs guardou referencia do slice original: obteve %q", got)
	}

	baseBc := []string{"a", "b"}
	bc := NewBreadcrumb(baseBc...)
	baseBc[0] = "ALTERADO"
	if got := bc.Items()[0]; got != "a" {
		t.Errorf("breadcrumb guardou referencia do slice original: obteve %q", got)
	}
}

func TestSidebarAddRemove(t *testing.T) {
	sb := NewSidebar("a", "b")

	if idx := sb.Add("c"); idx != 2 {
		t.Errorf("Add devolveu indice %d, esperado 2", idx)
	}
	if got := sb.Items(); len(got) != 3 || got[2] != "c" {
		t.Errorf("esperava [a b c], obteve %v", got)
	}

	if got := sb.Remove(1); got != "b" {
		t.Errorf("Remove devolveu %q, esperado %q", got, "b")
	}
	if got := sb.Items(); len(got) != 2 || got[1] != "c" {
		t.Errorf("esperava [a c], obteve %v", got)
	}

	if got := sb.Remove(99); got != "" {
		t.Errorf("Remove com indice invalido devolveu %q, esperado string vazia", got)
	}
	if got := sb.Remove(-1); got != "" {
		t.Errorf("Remove com indice negativo devolveu %q, esperado string vazia", got)
	}
}

func TestSidebarRemoveAjustaSelecao(t *testing.T) {
	// Remover um item anterior ao selecionado mantem o mesmo item.
	sb := NewSidebar("a", "b", "c")
	sb.SetSelected(2)
	sb.Remove(0)
	if sb.Selected() != 1 {
		t.Errorf("selecao esperada 1, obteve %d", sb.Selected())
	}
	if got := sb.Items()[sb.Selected()]; got != "c" {
		t.Errorf("esperava 'c' selecionado, obteve %q", got)
	}

	// Remover o item selecionado move a selecao para o proximo.
	sb = NewSidebar("a", "b", "c")
	sb.SetSelected(1)
	sb.Remove(1)
	if sb.Selected() != 1 {
		t.Errorf("selecao esperada 1, obteve %d", sb.Selected())
	}
	if got := sb.Items()[sb.Selected()]; got != "c" {
		t.Errorf("esperava 'c' selecionado, obteve %q", got)
	}

	// Remover o ultimo item selecionado limita a selecao ao ultimo restante.
	sb = NewSidebar("a", "b", "c")
	sb.SetSelected(2)
	sb.Remove(2)
	if sb.Selected() != 1 {
		t.Errorf("selecao esperada 1, obteve %d", sb.Selected())
	}
	if got := sb.Items()[sb.Selected()]; got != "b" {
		t.Errorf("esperava 'b' selecionado, obteve %q", got)
	}
}

func TestTabsAddRemove(t *testing.T) {
	tabs := NewTabs("a", "b")

	if idx := tabs.Add("c", "d"); idx != 2 {
		t.Errorf("Add devolveu indice %d, esperado 2", idx)
	}
	if got := tabs.Items(); len(got) != 4 {
		t.Errorf("esperava 4 itens, obteve %v", got)
	}

	if got := tabs.Remove(0); got != "a" {
		t.Errorf("Remove devolveu %q, esperado %q", got, "a")
	}
	if got := tabs.Items(); len(got) != 3 || got[0] != "b" {
		t.Errorf("esperava [b c d], obteve %v", got)
	}
}

func TestTabsRemoveAjustaAbaAtiva(t *testing.T) {
	tabs := NewTabs("a", "b", "c")
	tabs.SetActive(2)

	tabs.Remove(0)
	if tabs.Active() != 1 {
		t.Errorf("aba ativa esperada 1, obteve %d", tabs.Active())
	}
	if got := tabs.Items()[tabs.Active()]; got != "c" {
		t.Errorf("esperava 'c' ativa, obteve %q", got)
	}
}

func TestBreadcrumbAddRemove(t *testing.T) {
	bc := NewBreadcrumb("Inicio")

	if idx := bc.Add("Config", "Perfil"); idx != 1 {
		t.Errorf("Add devolveu indice %d, esperado 1", idx)
	}
	if got := bc.Items(); len(got) != 3 {
		t.Errorf("esperava 3 itens, obteve %v", got)
	}

	if got := bc.Remove(1); got != "Config" {
		t.Errorf("Remove devolveu %q, esperado %q", got, "Config")
	}
	if got := bc.Items(); len(got) != 2 || got[1] != "Perfil" {
		t.Errorf("esperava [Inicio Perfil], obteve %v", got)
	}

	// Push/Pop continuam funcionando.
	bc.Push("Extra")
	if got := bc.Pop(); got != "Extra" {
		t.Errorf("Pop devolveu %q, esperado %q", got, "Extra")
	}
}

func TestRenderVazio(t *testing.T) {
	if got := NewSidebar().Render(); got != "" {
		t.Errorf("esperava string vazia, obteve %q", got)
	}
	if got := NewTabs().Render(); got != "" {
		t.Errorf("esperava string vazia, obteve %q", got)
	}
	if got := NewBreadcrumb().Render(); got != "" {
		t.Errorf("esperava string vazia, obteve %q", got)
	}
}

func TestSidebarSetSelectedIgnoraIndicesInvalidos(t *testing.T) {
	sb := NewSidebar("a", "b")
	sb.SetSelected(-1)
	if sb.Selected() != 0 {
		t.Errorf("selecao esperada 0, obteve %d", sb.Selected())
	}
	sb.SetSelected(5)
	if sb.Selected() != 0 {
		t.Errorf("selecao esperada 0, obteve %d", sb.Selected())
	}
}

func TestSidebarLargurasMinimas(t *testing.T) {
	for _, width := range []int{1, 2, 3, 5, 10} {
		sb := NewSidebar("Dashboard", "Usuarios")
		sb.Width = width

		for i, linha := range strings.Split(sb.Render(), "\n") {
			if w := lipgloss.Width(linha); w != width {
				t.Errorf("width=%d: linha %d tem largura %d", width, i, w)
			}
		}
	}
}

func TestSidebarComCursorVazio(t *testing.T) {
	sb := NewSidebar("Dashboard", "Usuarios")
	sb.Cursor = ""

	linhas := strings.Split(sb.Render(), "\n")
	if len(linhas) != 2 {
		t.Fatalf("esperava 2 linhas, obteve %d", len(linhas))
	}
	if col0, col1 := colunaDoTexto(t, linhas[0], "Dashboard"), colunaDoTexto(t, linhas[1], "Usuarios"); col0 != col1 {
		t.Errorf("itens desalinhados: %d != %d", col0, col1)
	}
}

func TestSidebarComCursorMaiorQueAAreaUtil(t *testing.T) {
	sb := NewSidebar("Aa", "Bb")
	sb.Cursor = ">>>>>>>>>>"

	linhas := strings.Split(sb.Render(), "\n")
	for i, linha := range linhas {
		if w := lipgloss.Width(linha); w > sb.Width {
			t.Errorf("linha %d tem largura %d, maior que Width %d", i, w, sb.Width)
		}
	}
	if col0, col1 := colunaDoTexto(t, linhas[0], "Aa"), colunaDoTexto(t, linhas[1], "Bb"); col0 != col1 {
		t.Errorf("itens desalinhados: %d != %d", col0, col1)
	}
}
