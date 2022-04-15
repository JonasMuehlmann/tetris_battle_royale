package helpers

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/textinput"
)

var (
	FocusedStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	UnfocusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	NoStyle               = lipgloss.NewStyle()
	FullScreenRender      = lipgloss.NewStyle().Width(160).Height(90).Render
	BorderedRenderer      = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).Render
	ThickBorderedRenderer = lipgloss.NewStyle().Border(lipgloss.ThickBorder(), true).Render
	TitleRenderer         = FocusedStyle.Copy().Bold(true).Underline(true).Render
)

func FocusTextInput(element *textinput.Model) {
	element.Focus()
	element.TextStyle = FocusedStyle
	element.PromptStyle = FocusedStyle
}

func UnfocusTextInput(element *textinput.Model) {
	element.Blur()
	element.TextStyle = UnfocusedStyle
	element.PromptStyle = UnfocusedStyle
}
