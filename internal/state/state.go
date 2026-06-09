package state

import (
	"fmt"
	"os"

	"github.com/leonelquinteros/gotext"
)

func getDisabledFile() string {
	return os.ExpandEnv("$HOME/.config/umotd/disabled")
}

func IsDisabled() bool {
	_, err := os.Stat(getDisabledFile())
	return err == nil
}

func Enable(l *gotext.Locale) {
	err := os.Remove(getDisabledFile())
	if err != nil {
		fmt.Println(l.Get("Failed to enable the motd."))
		println(l.Get("Error ~> %s", err.Error()))
		return
	}
	fmt.Println(l.Get("The motd has been enabled."))
}

func Disable(l *gotext.Locale) {
	_, err := os.Create(getDisabledFile())
	if err != nil {
		fmt.Println(l.Get("Failed to disable the motd."))
		println(l.Get("Error ~> %s", err.Error()))
		return
	}
	fmt.Println(l.Get("The motd has been disabled."))
}
