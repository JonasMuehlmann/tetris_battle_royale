package views

import (
	"fmt"
	"log"
	"strings"
	"tui/internal/components"
	"tui/internal/helpers"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainMenuModel struct {
	focusedElement     int
	elements           []*textinput.Model
	usernameInput      textinput.Model
	passwordInput      textinput.Model
	registerButton     textinput.Model
	loginButton        textinput.Model
	viewToTransitionTo ViewName
	transitionInfo     map[string]string
	help               help.Model
	keymap             keyMap
	err                error
}

func (m *MainMenuModel) GetTransitionRequest() (viewToTransitionTo ViewName, transitionInfo map[string]string) {
	viewToTransitionTo = m.viewToTransitionTo
	transitionInfo = m.transitionInfo

	m.viewToTransitionTo = NoView

	return
}

func (m *MainMenuModel) HandleTransition(transitionInfo map[string]string) error {
	sessionID, ok := transitionInfo["sessionID"]
	if !ok {
		log.Panic("No session ID")
	}

	userID, ok := transitionInfo["userID"]
	if !ok {
		log.Panic("No user ID")
	}

	username, ok := transitionInfo["username"]
	if !ok {
		log.Panic("No username")
	}

	userIDButton := components.MakeButton(userID)
	sessionIDButton := components.MakeButton(sessionID)
	usernameButton := components.MakeButton(username)

	m.elements = append(m.elements, &userIDButton)
	m.elements = append(m.elements, &sessionIDButton)
	m.elements = append(m.elements, &usernameButton)

	return nil
}

func NewMainMenuModel() *MainMenuModel {
	model := &MainMenuModel{
		elements:           make([]*textinput.Model, 0),
		viewToTransitionTo: NoView,
		err:                nil,
	}

	usernameInput := textinput.New()
	usernameInput.CharLimit = 80
	usernameInput.Width = 80
	usernameInput.Prompt = "SessionID: "
	model.usernameInput = usernameInput
	model.elements = append(model.elements, &usernameInput)

	passwordInput := textinput.New()
	passwordInput.CharLimit = 80
	passwordInput.Width = 80
	passwordInput.Prompt = "UserID: "
	model.passwordInput = passwordInput
	model.elements = append(model.elements, &passwordInput)

	model.registerButton = components.MakeButton("Register")
	model.elements = append(model.elements, &model.registerButton)
	return model
}

func (m *MainMenuModel) Init() tea.Cmd {

	for _, element := range m.elements {
		helpers.UnfocusTextInput(element)
	}

	helpers.FocusTextInput(m.elements[m.focusedElement])

	return textinput.Blink
}
func (m *MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle keybinds
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if m.elements[m.focusedElement] == &m.loginButton {
				// TODO: login
				m.viewToTransitionTo = LoginView
			} else if m.elements[m.focusedElement] == &m.registerButton {
				// TODO: register
				m.viewToTransitionTo = LoginView
			}
		case "shift+tab", "up":
			helpers.UnfocusTextInput(m.elements[m.focusedElement])

			if m.focusedElement == 0 {
				m.focusedElement = len(m.elements) - 1
			} else {
				m.focusedElement--
			}

			helpers.FocusTextInput(m.elements[m.focusedElement])
		case "tab", "down":
			helpers.UnfocusTextInput(m.elements[m.focusedElement])

			if m.focusedElement == len(m.elements)-1 {
				m.focusedElement = 0
			} else {
				m.focusedElement++
			}

			helpers.FocusTextInput(m.elements[m.focusedElement])
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg

		return m, nil
	}

	for i, element := range m.elements {
		*m.elements[i], cmd = element.Update(msg)
	}

	return m, cmd
}

func (m *MainMenuModel) View() string {
	skeleton := `
%s
`[1:]

	subviews := make([]string, 0, len(m.elements))

	for _, element := range m.elements {
		subviews = append(subviews, element.View())
	}

	return fmt.Sprintf(skeleton,
		strings.Join(subviews, "\n"),
	)
}

func (m *MainMenuModel) GetHelpView() string {
	maxHeight := lipgloss.Height(m.help.FullHelpView(m.keymap.FullHelp()))
	maxWidth := lipgloss.Width(m.help.FullHelpView(m.keymap.FullHelp()))

	return lipgloss.Place(maxWidth, maxHeight, lipgloss.Center, lipgloss.Bottom, m.help.View(m.keymap))
}
