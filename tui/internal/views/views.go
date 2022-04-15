package views

import tea "github.com/charmbracelet/bubbletea"

type ViewName string

const (
	NoView       = ""
	LoginView    = "loginView"
	MainMenuView = "mainMenuView"
)

type View interface {
	tea.Model
	GetTransitionRequest() (ViewName, map[string]string)
	HandleTransition(map[string]string) error
	GetHelpView() string
}

type errMsg error
