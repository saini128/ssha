package ui

import (
	"fmt"
	"os"
	"os/exec"
	"ssha/db"
	"ssha/models"
	"ssha/sshc"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the UI state
type model struct {
	table    table.Model
	hosts    []models.Host
	filtered []models.Host
	search   textinput.Model
	mode     string
	selected int
	err      error
}

func initialModel() model {
	hosts, _ := db.GetHosts()

	columns := []table.Column{
		{Title: "Alias", Width: 30},
		{Title: "Hostname", Width: 30},
		{Title: "User", Width: 15},
		{Title: "Port", Width: 5},
	}
	rows := makeRows(hosts)
	tbl := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	search := textinput.New()
	search.Placeholder = "Type to search..."
	search.Focus()

	return model{
		table:    tbl,
		hosts:    hosts,
		filtered: hosts,
		search:   search,
		mode:     "search",
	}
}

func makeRows(hosts []models.Host) []table.Row {
	var rows []table.Row
	for _, h := range hosts {
		rows = append(rows, table.Row{h.Alias, h.Hostname, h.User, fmt.Sprintf("%d", h.Port)})
	}
	return rows
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+n":
			fmt.Println("[ADD HOST] Dialog should open.")
		case "ctrl+k":
			fmt.Println("[UPDATE HOST] Dialog for", m.filtered[m.selected])
		case "ctrl+d":
			fmt.Println("[DELETE HOST] Confirm deleting", m.filtered[m.selected])
		case "enter":
			selectedRow := m.table.SelectedRow()
			if selectedRow != nil {
				index := m.table.Cursor()
				host := m.hosts[index]
				// fmt.Print("\033[H\033[2J")
				cmd := exec.Command("clear")
				cmd.Stdout = os.Stdout
				_ = cmd.Run()
				if host.SecurityType == "password" {
					m.err = sshc.SSHintoHostViaPassword(host)
				} else {
					m.err = sshc.SSHintoHostViaPrivateKey(host)
				}
				if m.err != nil {
					return m, nil
				}
				return m, tea.Quit
			}
			return m, tea.Quit
		default:
			if msg.String() == "up" || msg.String() == "down" {
				m.table, _ = m.table.Update(msg)
				m.selected = m.table.Cursor()
			} else {
				m.search, _ = m.search.Update(msg)
				m.filterHosts()
			}
		}
	}
	return m, nil
}

func (m *model) filterHosts() {
	query := strings.ToLower(m.search.Value())
	var filtered []models.Host
	for _, host := range m.hosts {
		if strings.Contains(strings.ToLower(host.Alias), query) {
			filtered = append(filtered, host)
		}
	}
	m.filtered = filtered
	m.table.SetRows(makeRows(filtered))
	m.table.SetCursor(0)
	m.selected = 0
}

func (m model) View() string {
	return fmt.Sprintf("%s\n%s", m.search.View(), m.table.View())
}

func Run() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
