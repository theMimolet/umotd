package main

import (
	"embed"
	"fmt"
	"os"

	"umotd/internal"

	"github.com/leonelquinteros/gotext"
)

const version = "0.3.2"

//go:embed all:locales
var localesFS embed.FS

func main() {

	// Loads the locale based on the system's locale
	locale := internal.DetectLocale(localesFS)
	l := gotext.NewLocaleFSWithPath(locale, localesFS, "locales")
	l.AddDomain("default")

	// Handles command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		// Returns the path of the current config file
		case "config-path":
			fmt.Println(internal.TagsPath)
			return
		// Redirects to tag related commands
		case "tags":
			internal.TagsCommands(os.Args[2:], l)
			return
		// Prints the version
		case "version":
			fmt.Println(version)
			return
		default:
			fmt.Fprintln(os.Stderr, l.Get("Invalid command."))
			internal.Usage(l)
			os.Exit(1)
		}
	}

	// Loads the configuration from the system's config file
	cfg := internal.GetConfig()

	fmt.Print(internal.GetRandomMessage(l, cfg.Tags))
}
