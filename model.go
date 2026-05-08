// model.go
package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type screen int

const (
	screenMain screen = iota
	screenSettings
	screenDownload
)

type model struct {
	screen      screen
	urlInput    textinput.Model
	cookieInput textinput.Model
	inputFocus  int // 0 = url, 1 = cookie

	settingsFocus int
	cursors       [4]int // video quality, audio quality, audio lang, subs lang

	err string
}

var videoQualities = []string{"1080p", "720p", "480p", "360p", "240p"}
var audioQualities = []string{"192k", "128k", "96k"}
var audioLangs = []string{"en-US", "ja-JP", "es-LA", "fr-FR", "de-DE", "pt-BR"}
var subtitleLangs = []string{"en-US", "ja-JP", "es-LA", "fr-FR", "de-DE", "pt-BR", "off"}

func initialModel() model {
	url := textinput.New()
	url.Placeholder = "https://www.crunchyroll.com/series/..."
	url.Focus()
	url.Width = 60

	cookie := textinput.New()
	cookie.Placeholder = "etp_rt cookie value"
	cookie.Width = 60
	cookie.EchoMode = textinput.EchoPassword

	return model{
		screen:      screenMain,
		urlInput:    url,
		cookieInput: cookie,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "1":
			m.screen = screenMain
			m.inputFocus = 0
			m.urlInput.Focus()
			m.cookieInput.Blur()

		case "2":
			m.screen = screenSettings

		case "tab":
			if m.screen == screenMain {
				if m.inputFocus == 0 {
					m.inputFocus = 1
					m.urlInput.Blur()
					m.cookieInput.Focus()
				} else {
					m.inputFocus = 0
					m.cookieInput.Blur()
					m.urlInput.Focus()
				}
			}

		case "up":
			if m.screen == screenSettings && m.settingsFocus > 0 {
				m.settingsFocus--
			}
		case "down":
			if m.screen == screenSettings && m.settingsFocus < 3 {
				m.settingsFocus++
			}

		case "left":
			if m.screen == screenSettings {
				lengths := [4]int{len(videoQualities), len(audioQualities), len(audioLangs), len(subtitleLangs)}
				if m.cursors[m.settingsFocus] > 0 {
					m.cursors[m.settingsFocus]--
				} else {
					m.cursors[m.settingsFocus] = lengths[m.settingsFocus] - 1
				}
			}
		case "right":
			if m.screen == screenSettings {
				lengths := [4]int{len(videoQualities), len(audioQualities), len(audioLangs), len(subtitleLangs)}
				m.cursors[m.settingsFocus] = (m.cursors[m.settingsFocus] + 1) % lengths[m.settingsFocus]
			}

		case "enter":
			if m.screen == screenMain {
				if m.urlInput.Value() != "" && m.cookieInput.Value() != "" {
					m.screen = screenDownload
					return m, startDownload(m)
				}
			}
		}
	}

	var cmd tea.Cmd
	if m.screen == screenMain {
		if m.inputFocus == 0 {
			m.urlInput, cmd = m.urlInput.Update(msg)
		} else {
			m.cookieInput, cmd = m.cookieInput.Update(msg)
		}
	}

	return m, cmd
}

func (m model) View() string {
	switch m.screen {
	case screenMain:
		return mainView(m)
	case screenSettings:
		return settingsView(m)
	case screenDownload:
		return downloadView(m)
	}
	return ""
}

func (m model) selectedVideoQuality() string { return videoQualities[m.cursors[0]] }
func (m model) selectedAudioQuality() string { return audioQualities[m.cursors[1]] }
func (m model) selectedAudioLang() string    { return audioLangs[m.cursors[2]] }
func (m model) selectedSubtitleLang() string { return subtitleLangs[m.cursors[3]] }
