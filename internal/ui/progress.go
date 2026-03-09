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
}

type ProgressMsg float64
type DoneMsg struct{}

// tickMsg は一定間隔で進捗を進めるタイマーイベント
type tickMsg time.Time

func NewModel() Model {
	return Model{
		progress: progress.New(progress.WithDefaultGradient()),
	}
}

// tick は0.1秒後に tickMsg を送るコマンド
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
	if m.done {
		return "ガッテンだ！無事に荷を届けたぜ。\n"
	}
	return fmt.Sprintf("\n🏃💨 走るぜ！\n%s\n", m.progress.ViewAs(m.percent))
}
