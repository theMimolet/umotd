package args

import (
	"umotd/internal/config"

	"github.com/leonelquinteros/gotext"
)

func TipsPresetCommands(args []string, l *gotext.Locale) {
	if len(args) < 1 {
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
		for _, arg := range args {
			println(arg)
		}
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
		config.AddTipsPreset(arg, l)
	}
}

func RemoveTipsPresets(args []string, l *gotext.Locale) {
	if len(args) < 1 {
		println(l.Get("No presets to remove."))
		return
	}
	for _, arg := range args {
		config.RemoveTipsPreset(arg, l)
	}
}
