# charm-nav

Componentes de navegação para interfaces de terminal, estilizados com [Lipgloss](https://charm.land/lipgloss).

> [Read in English](README.md)

## Componentes

### Tabs (Abas)

Barra de abas horizontal com estilo ativo/inativo e separadores configuráveis.

```go
tabs := nav.NewTabs("Inicio", "Configuracoes", "Sobre")
tabs.Next()
fmt.Println(tabs.Render())
// Inicio | Configuracoes | Sobre
//          ^^^^^^^^^^^^^ (ativa, negrito + sublinhado)
```

**API:**

| Método | Descrição |
|--------|-----------|
| `NewTabs(items ...string)` | Cria uma nova barra de abas |
| `SetActive(index int)` | Define a aba ativa pelo índice |
| `Next()` | Avança para a próxima aba (com wraparound) |
| `Previous()` | Volta para a aba anterior (com wraparound) |
| `Active() int` | Retorna o índice da aba ativa |
| `Items() []string` | Retorna uma **cópia** de todos os rótulos das abas |
| `Add(items ...string) int` | Adiciona abas no final; retorna o índice da primeira adicionada (`-1` se nenhuma) |
| `Remove(index int) string` | Remove a aba no índice informado; retorna o rótulo (`""` se o índice for inválido) |
| `Render() string` | Renderiza a barra de abas estilizada |

**Campos customizáveis:** `Separator`, `ActiveStyle`, `InactiveStyle`

---

### Breadcrumb (Trilha de navegação)

Navegação em trilha com pilha push/pop e separador configurável.

```go
bc := nav.NewBreadcrumb("Inicio", "Configuracoes")
bc.Push("Perfil")
fmt.Println(bc.Render())
// Inicio > Configuracoes > Perfil
//                           ^^^^^^ (ativo, negrito)
```

**API:**

| Método | Descrição |
|--------|-----------|
| `NewBreadcrumb(items ...string)` | Cria uma nova trilha de navegação |
| `Push(item string)` | Adiciona um item ao final |
| `Pop() string` | Remove e retorna o último item |
| `Items() []string` | Retorna uma **cópia** de todos os itens do breadcrumb |
| `Add(items ...string) int` | Adiciona itens no final; retorna o índice do primeiro adicionado (`-1` se nenhum) |
| `Remove(index int) string` | Remove o item no índice informado; retorna o rótulo (`""` se o índice for inválido) |
| `Render() string` | Renderiza o breadcrumb estilizado |

**Campos customizáveis:** `Separator`, `ActiveStyle`, `InactiveStyle`

---

### Sidebar (Menu lateral)

Menu vertical com indicador de cursor e destaque de seleção.

```go
sidebar := nav.NewSidebar("Painel", "Usuarios", "Configuracoes", "Logs")
sidebar.SetSelected(1)
fmt.Println(sidebar.Render())
//    Painel
//  ▸ Usuarios        (selecionado, destacado)
//    Configuracoes
//    Logs
```

**API:**

| Método | Descrição |
|--------|-----------|
| `NewSidebar(items ...string)` | Cria um novo menu lateral |
| `SetSelected(index int)` | Define o item selecionado pelo índice |
| `Next()` | Avança para o próximo item (com wraparound) |
| `Previous()` | Volta para o item anterior (com wraparound) |
| `Selected() int` | Retorna o índice do item selecionado |
| `Items() []string` | Retorna uma **cópia** de todos os itens do menu |
| `Add(items ...string) int` | Adiciona itens no final; retorna o índice do primeiro adicionado (`-1` se nenhum) |
| `Remove(index int) string` | Remove o item no índice informado; retorna o rótulo (`""` se o índice for inválido) |
| `Render() string` | Renderiza o menu lateral estilizado |

**Campos customizáveis:** `Width`, `Cursor`, `SelectedStyle`, `UnselectedStyle`

**Regras de renderização:**

- O cursor é medido em **colunas de exibição**, e não em bytes — `"▸ "` ocupa 2 colunas
  mesmo tendo 4 bytes, então todos os itens ficam alinhados.
- Itens que não cabem em `Width` são **truncados**, nunca quebrados em duas linhas, para
  que um rótulo longo nunca estoure a largura do terminal.
- O destaque do item selecionado cobre a **linha inteira**, e não apenas o texto.
- O `Remove` mantém a seleção coerente: remover um item anterior ao selecionado mantém
  o mesmo item selecionado, e remover o item selecionado move a seleção para o próximo.

---

## Observações

- `Items()` devolve uma cópia: alterar o slice retornado nunca muda o componente.
- Os construtores copiam o slice variádico, então quem chamou continua sendo dono dele.
- `Add`/`Remove` estão disponíveis em todos os componentes (`Tabs`, `Breadcrumb` e `Sidebar`).

## Instalação

```bash
go get github.com/junhinhow/charm-nav
```

## Licença

MIT
