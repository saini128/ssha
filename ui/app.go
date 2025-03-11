package ui

import (
	"fmt"
	"os"
	"ssha/db"
	"ssha/models"
	"ssha/sshc" // Or ssha/ssh, make sure it matches your folder name

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	table table.Model
	hosts []models.Host
	err   error // Add error field
}

func initialModel() model {
	hosts, _ := db.GetHosts()

	columns := []table.Column{
		{Title: "Alias", Width: 30},
		{Title: "Hostname", Width: 30},
		{Title: "User", Width: 15},
		{Title: "Port", Width: 5},
	}

	var rows []table.Row
	for _, h := range hosts {
		rows = append(rows, table.Row{h.Alias, h.Hostname, h.User, fmt.Sprintf("%d", h.Port)})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return model{table: t, hosts: hosts}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return m, tea.Quit
		case "enter":
			selectedRow := m.table.SelectedRow()
			if selectedRow != nil {
				index := m.table.Cursor()
				host := m.hosts[index]
				if host.SecurityType == "password" {
					m.err = sshc.SSHintoHostViaPassword(host) // Store the error

				} else {
					m.err = sshc.SSHintoHostViaPrivateKey(host) // Store the error

				}
				if m.err != nil {
					return m, nil // Don't quit on error, show the error
				}
				return m, tea.Quit // Quit only on successful connection
			}
		}
	}

	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n\n%s", m.err, m.table.View()) // Display error
	}
	return m.table.View()
}

func Run() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
