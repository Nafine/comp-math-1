package model

import "github.com/charmbracelet/lipgloss"

const (
	hotPink = lipgloss.Color("#FF06B7")
	hotRed  = lipgloss.Color("#FF0000")
)

var (
	inputStyle = lipgloss.NewStyle().Foreground(hotPink)
	errorStyle = lipgloss.NewStyle().Foreground(hotRed)
	cellStyle  = lipgloss.NewStyle().
			Width(10).
			PaddingRight(1)
)
