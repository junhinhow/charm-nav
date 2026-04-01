# charm-nav

Componentes de navegacao para interfaces de terminal, estilizados com [Lipgloss](https://charm.land/lipgloss).

> [Read in English](README.md)

## Componentes

### Tabs (Abas)

Barra de abas horizontal com estilo ativo/inativo e separadores configuraveis.

```go
tabs := nav.NewTabs("Inicio", "Configuracoes", "Sobre")
tabs.Next()
fmt.Println(tabs.Render())
// Inicio | Configuracoes | Sobre
//          ^^^^^^^^^^^^^ (ativa, negrito + sublinhado)
```

**API:**

| Metodo | Descricao |
|--------|-----------|
| `NewTabs(items ...string)` | Cria uma nova barra de abas |
| `SetActive(index int)` | Define a aba ativa pelo indice |
| `Next()` | Avanca para a proxima aba (com wraparound) |
| `Previous()` | Volta para a aba anterior (com wraparound) |
| `Active() int` | Retorna o indice da aba ativa |
| `Items() []string` | Retorna todos os rotulos das abas |
| `Render() string` | Renderiza a barra de abas estilizada |

**Campos customizaveis:** `Separator`, `ActiveStyle`, `InactiveStyle`

---

### Breadcrumb (Trilha de navegacao)

Navegacao em trilha com pilha push/pop e separador configuravel.

```go
bc := nav.NewBreadcrumb("Inicio", "Configuracoes")
bc.Push("Perfil")
fmt.Println(bc.Render())
// Inicio > Configuracoes > Perfil
//                           ^^^^^^ (ativo, negrito)
```

**API:**

| Metodo | Descricao |
|--------|-----------|
| `NewBreadcrumb(items ...string)` | Cria uma nova trilha de navegacao |
| `Push(item string)` | Adiciona um item ao final |
| `Pop() string` | Remove e retorna o ultimo item |
| `Items() []string` | Retorna todos os itens do breadcrumb |
| `Render() string` | Renderiza o breadcrumb estilizado |

**Campos customizaveis:** `Separator`, `ActiveStyle`, `InactiveStyle`

---

### Sidebar (Menu lateral)

Menu vertical com indicador de cursor e destaque de selecao.

```go
sidebar := nav.NewSidebar("Painel", "Usuarios", "Configuracoes", "Logs")
sidebar.SetSelected(1)
fmt.Println(sidebar.Render())
//   Painel
// ▸ Usuarios      (selecionado, destacado)
//   Configuracoes
//   Logs
```

**API:**

| Metodo | Descricao |
|--------|-----------|
| `NewSidebar(items ...string)` | Cria um novo menu lateral |
| `SetSelected(index int)` | Define o item selecionado pelo indice |
| `Next()` | Avanca para o proximo item (com wraparound) |
| `Previous()` | Volta para o item anterior (com wraparound) |
| `Selected() int` | Retorna o indice do item selecionado |
| `Items() []string` | Retorna todos os itens do menu |
| `Render() string` | Renderiza o menu lateral estilizado |

**Campos customizaveis:** `Width`, `Cursor`, `SelectedStyle`, `UnselectedStyle`

---

## Instalacao

```bash
go get github.com/junhinhow/charm-nav
```

## Licenca

MIT
