package views

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"tui/internal/components"
	"tui/internal/helpers"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyMap struct {
	Up      key.Binding
	Down    key.Binding
	Help    key.Binding
	Quit    key.Binding
	Confirm key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Confirm}, // first column
		{k.Help, k.Quit},          // second column
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "shift+tab"),
		key.WithHelp("↑/⇧⭾", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "tab"),
		key.WithHelp("↓/⭾ ", "move down"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("⏎", "confirm"),
	),
}

type LoginModel struct {
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

func (m *LoginModel) GetTransitionRequest() (viewToTransitionTo ViewName, transitionInfo map[string]string) {
	viewToTransitionTo = m.viewToTransitionTo
	transitionInfo = m.transitionInfo

	m.viewToTransitionTo = NoView

	return
}

func (m *LoginModel) HandleTransition(transitionInfo map[string]string) error {
	return nil
}

func NewLoginModel() *LoginModel {
	model := &LoginModel{
		elements:           make([]*textinput.Model, 0),
		viewToTransitionTo: NoView,
		err:                nil,
		keymap:             keys,
		help:               help.New(),
	}

	usernameInput := textinput.New()
	usernameInput.CharLimit = 20
	usernameInput.Width = 20
	usernameInput.Prompt = "Username: "
	model.usernameInput = usernameInput
	model.elements = append(model.elements, &model.usernameInput)

	passwordInput := textinput.New()
	passwordInput.Width = 20
	passwordInput.CharLimit = 20
	passwordInput.EchoMode = textinput.EchoPassword
	passwordInput.Prompt = "Password: "
	model.passwordInput = passwordInput
	model.elements = append(model.elements, &model.passwordInput)

	model.registerButton = components.MakeButton("Register")
	model.elements = append(model.elements, &model.registerButton)

	model.loginButton = components.MakeButton("Login")
	model.elements = append(model.elements, &model.loginButton)

	return model
}
func (m *LoginModel) Init() tea.Cmd {
	for _, element := range m.elements {
		helpers.UnfocusTextInput(element)
	}

	helpers.FocusTextInput(m.elements[m.focusedElement])

	return textinput.Blink
}
func (m *LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// NOTE: This **** is very important!!! if you forget it, you will waste 2h debugging why the extended help is no showing...
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		// Handle keybinds
		switch {
		case key.Matches(msg, m.keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.Confirm):
			if m.elements[m.focusedElement] == &m.loginButton {
				// TODO: login
				client := &http.Client{}
				bodyRaw := map[string]string{"username": "foo", "password": "bar"}
				body, _ := json.Marshal(bodyRaw)
				res, err := client.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(body))
				if err != nil {
					log.Panic(err)
				}

				m.viewToTransitionTo = MainMenuView

				var response map[string]string
				responseRaw, err := ioutil.ReadAll(res.Body)
				err = json.Unmarshal(responseRaw, &response)
				if err != nil {
					log.Panic(err)
				}

				m.transitionInfo = response

			} else if m.elements[m.focusedElement] == &m.registerButton {
				// TODO: register
				m.viewToTransitionTo = MainMenuView
			}
		case key.Matches(msg, m.keymap.Up):
			helpers.UnfocusTextInput(m.elements[m.focusedElement])

			if m.focusedElement == 0 {
				m.focusedElement = len(m.elements) - 1
			} else {
				m.focusedElement--
			}

			helpers.FocusTextInput(m.elements[m.focusedElement])
		case key.Matches(msg, m.keymap.Down):
			helpers.UnfocusTextInput(m.elements[m.focusedElement])

			if m.focusedElement == len(m.elements)-1 {
				m.focusedElement = 0
			} else {
				m.focusedElement++
			}

			helpers.FocusTextInput(m.elements[m.focusedElement])
		case key.Matches(msg, m.keymap.Help):
			m.help.ShowAll = !m.help.ShowAll
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

func (m *LoginModel) View() string {

	style := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true)

	return style.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			lipgloss.JoinVertical(lipgloss.Left, m.usernameInput.View(), m.passwordInput.View()),
			lipgloss.NewStyle().MarginTop(1).Render(lipgloss.JoinHorizontal(lipgloss.Center, m.registerButton.View(), m.loginButton.View())),
		),
	)
}

func (m *LoginModel) GetHelpView() string {
	maxHeight := lipgloss.Height(m.help.FullHelpView(m.keymap.FullHelp()))
	maxWidth := lipgloss.Width(m.help.FullHelpView(m.keymap.FullHelp()))

	return lipgloss.Place(maxWidth, maxHeight, lipgloss.Center, lipgloss.Bottom, m.help.View(m.keymap))
}
