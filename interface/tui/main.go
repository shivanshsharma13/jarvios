package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	input   string
	history []string
	waiting bool
}

var (
	userStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Bold(true)
	ariaStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	inputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
)

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if strings.TrimSpace(m.input) == "" {
				return m, nil
			}
			userMsg := m.input
			m.history = append(m.history, "you: "+userMsg)
			m.input = ""
			m.waiting = true
			return m, sendMessage(userMsg)
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			m.input += msg.String()
		}
	case replyMsg:
		m.waiting = false
		m.history = append(m.history, "aria: "+string(msg))
	}
	return m, nil
}

func (m model) View() string {
	var sb strings.Builder
	sb.WriteString("\n  ARIA OS — type a message, Ctrl+C to quit\n\n")
	for _, line := range m.history {
		if strings.HasPrefix(line, "you:") {
			sb.WriteString("  " + userStyle.Render(line) + "\n")
		} else {
			sb.WriteString("  " + ariaStyle.Render(line) + "\n")
		}
	}
	if m.waiting {
		sb.WriteString("\n  aria is thinking...\n")
	}
	sb.WriteString("\n  " + inputStyle.Render("> "+m.input))
	return sb.String()
}

type replyMsg string

func sendMessage(msg string) tea.Cmd {
	return func() tea.Msg {
		body, _ := json.Marshal(map[string]string{"message": msg})
		resp, err := http.Post("http://localhost:7777/chat",
			"application/json", bytes.NewReader(body))
		if err != nil {
			return replyMsg("error: daemon not running")
		}
		defer resp.Body.Close()
		var result map[string]string
		data, _ := io.ReadAll(resp.Body)
		json.Unmarshal(data, &result)
		return replyMsg(result["reply"])
	}
}

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
