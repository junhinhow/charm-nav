# charm-nav

Navigation components for terminal UIs, styled with [Lipgloss](https://charm.land/lipgloss).

> [Leia em Portugues](README.pt-br.md)

## Components

### Tabs

Horizontal tab bar with active/inactive styling and configurable separators.

```go
tabs := nav.NewTabs("Home", "Settings", "About")
tabs.Next()
fmt.Println(tabs.Render())
// Home | Settings | About
//         ^^^^^^^^ (active, bold + underline)
```

**API:**

| Method | Description |
|--------|-------------|
| `NewTabs(items ...string)` | Create a new tab bar |
| `SetActive(index int)` | Set the active tab by index |
| `Next()` | Move to the next tab (wraps around) |
| `Previous()` | Move to the previous tab (wraps around) |
| `Active() int` | Get the active tab index |
| `Items() []string` | Get all tab labels |
| `Render() string` | Render the styled tab bar |

**Customizable fields:** `Separator`, `ActiveStyle`, `InactiveStyle`

---

### Breadcrumb

Trail navigation with push/pop stack and configurable separator.

```go
bc := nav.NewBreadcrumb("Home", "Settings")
bc.Push("Profile")
fmt.Println(bc.Render())
// Home > Settings > Profile
//                    ^^^^^^^ (active, bold)
```

**API:**

| Method | Description |
|--------|-------------|
| `NewBreadcrumb(items ...string)` | Create a new breadcrumb trail |
| `Push(item string)` | Append an item |
| `Pop() string` | Remove and return the last item |
| `Items() []string` | Get all breadcrumb items |
| `Render() string` | Render the styled breadcrumb |

**Customizable fields:** `Separator`, `ActiveStyle`, `InactiveStyle`

---

### Sidebar

Vertical menu with cursor indicator and selection highlighting.

```go
sidebar := nav.NewSidebar("Dashboard", "Users", "Settings", "Logs")
sidebar.SetSelected(1)
fmt.Println(sidebar.Render())
//   Dashboard
// ▸ Users        (selected, highlighted)
//   Settings
//   Logs
```

**API:**

| Method | Description |
|--------|-------------|
| `NewSidebar(items ...string)` | Create a new sidebar menu |
| `SetSelected(index int)` | Set the selected item by index |
| `Next()` | Move to the next item (wraps around) |
| `Previous()` | Move to the previous item (wraps around) |
| `Selected() int` | Get the selected item index |
| `Items() []string` | Get all menu items |
| `Render() string` | Render the styled sidebar |

**Customizable fields:** `Width`, `Cursor`, `SelectedStyle`, `UnselectedStyle`

---

## Install

```bash
go get github.com/junhinhow/charm-nav
```

## License

MIT
