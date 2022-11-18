package create

import "strings"

func DiscordStyle(s string) string {
	if !strings.ContainsAny(s, "[]()") {
		return ">>> " + s
	}
	s = strings.ReplaceAll(s, "[", "**__")
	s = strings.ReplaceAll(s, "]", "__**")
	s = strings.ReplaceAll(s, "(", "**")
	s = strings.ReplaceAll(s, ")", "**")

	return ">>> " + s
}
