package console

import (
	"fmt"
	"image/color"
	"strings"
	"sync"
	"time"

	"charm.land/lipgloss/v2"
)

var (
	mauve  = lipgloss.Color("#C6A0F6")
	pink   = lipgloss.Color("#F5BDE6")
	green  = lipgloss.Color("#A6DA95")
	yellow = lipgloss.Color("#EED49F")
	red    = lipgloss.Color("#ED8796")
	blue   = lipgloss.Color("#8AADF4")
	dim    = lipgloss.Color("#6E738D")
	text   = lipgloss.Color("#CAD3F5")
)

type level struct {
	name  string
	icon  string
	color color.Color
}

var (
	lDebug   = level{"debug", "·", dim}
	lInfo    = level{"info", "●", blue}
	lSuccess = level{"success", "✔", green}
	lWarn    = level{"warn", "▲", yellow}
	lError   = level{"error", "✘", red}
)

var mu sync.Mutex

func fg(c color.Color) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(c)
}

func emit(s string) {
	mu.Lock()
	defer mu.Unlock()
	lipgloss.Println(s)
}

func stamp() string {
	return fg(dim).Render(time.Now().Format("15:04:05"))
}

func write(l level, format string, args []any) {
	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	tag := fg(l.color).Bold(true).Width(9).Render(l.icon + " " + l.name)
	emit(stamp() + "  " + tag + " " + fg(text).Render(msg))
}

func Debug(format string, args ...any)   { write(lDebug, format, args) }
func Info(format string, args ...any)    { write(lInfo, format, args) }
func Success(format string, args ...any) { write(lSuccess, format, args) }
func Warn(format string, args ...any)    { write(lWarn, format, args) }
func Error(format string, args ...any)   { write(lError, format, args) }

func Command(name, user, guild string, took time.Duration) {
	speed := green
	if took > time.Second {
		speed = red
	} else if took > 200*time.Millisecond {
		speed = yellow
	}

	emit(stamp() + "  " +
		fg(pink).Bold(true).Width(9).Render("✦ cmd") + " " +
		fg(text).Bold(true).Render("/"+name) + "  " +
		fg(blue).Render(user) +
		fg(dim).Render(" in "+guild) + "  " +
		fg(speed).Render(took.Round(time.Millisecond).String()))
}

func Banner(name, version string) {
	letters := strings.Split(name, "")
	if len(letters) == 0 {
		return
	}

	shades := lipgloss.Blend1D(len(letters), mauve, pink)
	spaced := make([]string, len(letters))
	for i, l := range letters {
		spaced[i] = fg(shades[i]).Bold(true).Render(l)
	}

	body := lipgloss.JoinVertical(lipgloss.Center,
		strings.Join(spaced, ""),
		"",
		fg(dim).Italic(true).Render(version),
	)

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForegroundBlend(mauve, pink).
		Padding(1, 6)

	lipgloss.Println()
	lipgloss.Println(box.Render(body))
	lipgloss.Println()
}

func Ready(tag string, guilds int, ping time.Duration) {
	rows := [][2]string{
		{"user", tag},
		{"guilds", fmt.Sprint(guilds)},
		{"ping", ping.Round(time.Millisecond).String()},
		{"since", time.Now().Format("3:04 PM")},
	}

	label := lipgloss.NewStyle().Foreground(dim).Width(9)
	value := fg(text).Bold(true)

	lines := []string{fg(green).Bold(true).Render("● online"), ""}
	for _, r := range rows {
		lines = append(lines, label.Render(r[0])+value.Render(r[1]))
	}

	card := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(dim).
		Padding(1, 3)

	lipgloss.Println(card.Render(lipgloss.JoinVertical(lipgloss.Left, lines...)))
	lipgloss.Println()
}

func Goodbye() {
	lipgloss.Println()
	lipgloss.Println(fg(pink).Italic(true).Render("✦ | goodbye... for now...\n"))
}
