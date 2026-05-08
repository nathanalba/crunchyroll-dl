package main

import tea "github.com/charmbracelet/bubbletea"

type downloadDoneMsg struct{}
type downloadErrMsg struct{ err error }

func startDownload(m model) tea.Cmd {
	return func() tea.Msg {
		// wire up your actual download logic here
		// using m.urlInput.Value(), m.cookieInput.Value(),
		// videoQualities[m.qualityCursor], m.audioQuality
		return downloadDoneMsg{}
	}
}
