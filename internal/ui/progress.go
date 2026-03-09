package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	progress progress.Model
	percent  float64
	done     bool
	src      string
	dst      string
}

type ProgressMsg float64
type DoneMsg struct{}

type tickMsg time.Time

func NewModel(src, dst string) Model {
	return Model{
		progress: progress.New(progress.WithDefaultGradient()),
		src:      src,
		dst:      dst,
	}
}

func tick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tickMsg:
		if m.percent >= 1.0 {
			m.done = true
			return m, tea.Quit
		}
		m.percent += 0.05
		return m, tick()
	case tea.KeyMsg:
		return m, tea.Quit
	}
	var cmd tea.Cmd
	progressModel, cmd := m.progress.Update(msg)
	m.progress = progressModel.(progress.Model)
	return m, cmd
}

func (m Model) View() string {
	station := CurrentStation(m.percent)

	if m.done {
		end := EndStation()
		return fmt.Sprintf(
			"\n出発地点：[%s] ── 到着地点：[%s]\n宿場町[%s] 「%s」\n\n",
			m.src, m.dst,
			end.Name, end.Quote,
		)
	}

	start := StartStation()
	end := EndStation()
	bar := m.progress.ViewAs(m.percent)

	return fmt.Sprintf(
		"\n出発地点：[%s] ── 到着地点：[%s]\n宿場町[%s]%s宿場町[%s]\n🏃💨 %s\n「%s」\n",
		m.src, m.dst,
		start.Name,
		"　　　　　　　　　　　　　　　　　　　　",
		end.Name,
		bar,
		station.Quote,
	)
}
