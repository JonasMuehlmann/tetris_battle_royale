package main

import (
	"log"
	"tui/internal/helpers"
	"tui/internal/views"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	Width  = 90
	Height = 30
)

func main() {
	programModel := NewProgramModel()

	p := tea.NewProgram(programModel)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

type ProgramModel struct {
	curView views.View
	views   map[views.ViewName]views.View
}

func NewProgramModel() *ProgramModel {
	model := &ProgramModel{
		views: make(map[views.ViewName]views.View),
	}

	// Register all views
	model.views[views.LoginView] = views.NewLoginModel()
	model.views[views.MainMenuView] = views.NewMainMenuModel()

	model.curView = model.views[views.LoginView]

	return model
}

func (m *ProgramModel) Init() tea.Cmd {
	return m.curView.Init()
}

func (m *ProgramModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	newModel, cmd := m.curView.Update(message)

	viewToTransitionTo, transitionInfo := newModel.(views.View).GetTransitionRequest()

	if viewToTransitionTo != views.NoView {
		m.curView = m.views[viewToTransitionTo]
		m.curView.Init()
		m.curView.HandleTransition(transitionInfo)
	}

	return m, cmd
}

func (m *ProgramModel) View() string {
	title := helpers.TitleRenderer("Tetris Battle Royale")
	help := m.curView.GetHelpView()
	center := lipgloss.Place(Width, Height-lipgloss.Height(title)-lipgloss.Height(help), lipgloss.Center, lipgloss.Center, m.curView.View())

	return helpers.ThickBorderedRenderer(lipgloss.JoinVertical(lipgloss.Center, title, center, help))
}
