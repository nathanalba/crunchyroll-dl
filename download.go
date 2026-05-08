package main

import tea "github.com/charmbracelet/bubbletea"

type downloadDoneMsg struct{}
type downloadErrMsg struct{ err error }
type progressMsg struct{ text string }

func startDownload(m model) tea.Cmd {
	return func() tea.Msg {
		etpRt = m.cookieInput.Value()
		token = GetAccessToken(etpRt)
		processUrl(
			m.urlInput.Value(),
			m.selectedVideoQuality(),
			m.selectedAudioQuality(),
			m.selectedAudioLang(),
			m.selectedSubtitleLang(),
		)
		return downloadDoneMsg{}
	}
}
