package components

import "github.com/charmbracelet/bubbles/textinput"

func MakeButton(text string) textinput.Model {
	button := textinput.New()

	button.Prompt = " [ " + text + " ]"
	button.CharLimit = 0
	button.EchoMode = textinput.EchoNone
	button.Width = 0

	return button
}
