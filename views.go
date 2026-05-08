// views.go
package main

import "fmt"

func navbar(active screen) string {
	views := []struct {
		key   string
		label string
		s     screen
	}{
		{"1", "main", screenMain},
		{"2", "settings", screenSettings},
	}

	bar := "  "
	for _, v := range views {
		if v.s == active {
			bar += fmt.Sprintf("[%s] %s  ", v.key, v.label)
		} else {
			bar += fmt.Sprintf(" %s  %s  ", v.key, v.label)
		}
	}
	return bar + "\n\n"
}

func mainView(m model) string {
	s := navbar(screenMain)
	s += "  URL\n"
	s += "  " + m.urlInput.View() + "\n\n"
	s += "  etp_rt cookie\n"
	s += "  " + m.cookieInput.View() + "\n\n"
	s += "  [tab] switch field  [enter] start download  [ctrl+c] quit\n"
	return s
}

func settingsView(m model) string {
	s := navbar(screenSettings)

	settings := []struct {
		label  string
		values []string
		cur    int
	}{
		{"Video Quality ", videoQualities, m.cursors[0]},
		{"Audio Quality ", audioQualities, m.cursors[1]},
		{"Audio Language", audioLangs, m.cursors[2]},
		{"Subtitle Lang ", subtitleLangs, m.cursors[3]},
	}

	for i, setting := range settings {
		if i == m.settingsFocus {
			s += fmt.Sprintf("  > %s   < %s >\n", setting.label, setting.values[setting.cur])
		} else {
			s += fmt.Sprintf("    %s     %s\n", setting.label, setting.values[setting.cur])
		}
	}

	s += "\n  [up/down] select  [left/right] change  [1] back to main\n"
	return s
}

func downloadView(m model) string {
	if m.err != "" {
		return fmt.Sprintf("  Error: %s\n", m.err)
	}
	return "  Downloading...\n"
}
