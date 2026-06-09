package args

import (
	"umotd/internal/config"

	"github.com/leonelquinteros/gotext"
)

func TipsPresetCommands(args []string, l *gotext.Locale) {
	if len(args) < 1 {
		println(l.Get("No command specified."))
		return
	}
	switch args[0] {
	case "list":
		ListTipsPresets()
	case "add":
		AddTipsPresets(args[1:], l)
	case "remove":
		RemoveTipsPresets(args[1:], l)
	default:
		println(l.Get("Command not recognized. Try 'list', 'add', or 'remove'."))
		return
	}
}

func ListTipsPresets() {
	presets := config.ListTipsPresets()
	for _, preset := range presets {
		println(preset)
	}
}

func AddTipsPresets(args []string, l *gotext.Locale) {
	if len(args) < 1 {
		println(l.Get("No presets to add."))
		return
	}
	for _, arg := range args {
		if err := config.AddTipsPreset(arg, l); err != nil {
			println(l.Get("Failed to add preset: %s", arg))
			println(l.Get("Error ~> %s", err.Error()))
		}
	}
}

func RemoveTipsPresets(args []string, l *gotext.Locale) {
	if len(args) < 1 {
		println(l.Get("No presets to remove."))
		return
	}
	for _, arg := range args {
		if err := config.RemoveTipsPreset(arg, l); err != nil {
			println(l.Get("Failed to remove preset: %s", arg))
			println(l.Get("Error ~> %s", err.Error()))
		}
	}
}
