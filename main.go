package main

import (
	"fmt"
	"os"
	"strings"

	"charm.land/glamour/v2"
	"github.com/leonelquinteros/gotext"
)

const VERSION = "0.2"

func main() {

	// Loads the locale based on the system's locale
	locale := detectLocale(localesFS)
	l := gotext.NewLocaleFSWithPath(locale, localesFS, "locales")
	l.AddDomain("default")

	isDisabled := isDisabled()

	// Handles command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {

		// Prints the version
		case "--version", "-v", "version":
			fmt.Println(VERSION)
			return

		// Enables the motd
		case "enable":
			if isDisabled {
				err := os.Remove(getDisabledFile())
				if err != nil {
					fmt.Println(l.Get("Failed to enable the motd."))
					return
				}
				fmt.Println(l.Get("The motd has been enabled."))
				return
			} else {
				fmt.Println(l.Get("The motd is already enabled."))
				return
			}

		// Disables the motd
		case "disable":
			if isDisabled {
				fmt.Println(l.Get("The motd is already disabled."))
				return
			} else {
				_, err := os.Create(getDisabledFile())
				if err != nil {
					fmt.Println(l.Get("Failed to disable the motd."))
					return
				}
				fmt.Println(l.Get("The motd has been disabled."))
				return
			}
		default:
			fmt.Println(l.Get("Invalid command"))
			return
		}
	}

	// Exits if the motd is disabled
	if isDisabled {
		os.Exit(0)
	}

	// Loads the configuration from the system's config file
	cfg := getConfig()

	// Gets the image info and OS name
	in := "# " + cfg.Prefix + l.Get("Welcome to %s", getOSName()) + cfg.Suffix + "\n"
	if imageInfo := getImageInfo(cfg.InfoFile); imageInfo.ImageRef != "" || imageInfo.ImageTag != "" {
		in += " 󱋩 `" + imageInfo.ImageRef + ":" + imageInfo.ImageTag + "` \n"
	} else if isBootcSystem() {
		in += " 󱋩 `" + l.Get("Unknown system") + "` \n"
	}

	// Gets the Greenboot status
	if greenboot := getGreenbootInfo(); greenboot != "" {
		in += "\n 󰟀  " + l.Get("Boot Status") + ":"
		if greenboot == "healthy" {
			in += "`" + l.Get("Healthy") + " 󰄳`"
		} else {
			in += "`" + greenboot + "`"
		}
		in += " \n"
	}

	// Command list
	if len(cfg.Commands) > 0 {
		in += " |  " + l.Get("Command") + " | " + l.Get("Description") + " | \n"
		in += "| ------------ | ----------- |\n"
		var cmdSb strings.Builder
		for _, cmd := range cfg.Commands {
			switch cmd.Desc {
			case "cmd_list":
				cmd.Desc = l.Get("List all available commands")
			case "cli_pkg":
				cmd.Desc = l.Get("Manage command line packages")
			case "term_bling":
				cmd.Desc = l.Get("Enable terminal bling")
			case "motd_toggle":
				cmd.Desc = l.Get("Toggle this banner on/off")
			case "sys_info":
				cmd.Desc = l.Get("View system information")
			case "man_upd":
				cmd.Desc = l.Get("Manually update the system")
			}
			fmt.Fprintf(&cmdSb, "| `%s` | %s |\n", cmd.Cmd, cmd.Desc)
		}
		in += cmdSb.String()
		in += "\n"
	}

	// Gets a random tip
	in += getRandomTip(cfg.Tips, cfg.TipsPresets...) + "\n\n"

	// Gets the links
	if len(cfg.Links) > 0 {
		var linkSb strings.Builder
		for _, link := range cfg.Links {
			switch link.Name {
			case "website":
				link.Name = "󰌹 [" + l.Get("Website") + "]"
			case "issues":
				link.Name = "󰊤 [" + l.Get("Report an issue") + "]"
			case "docs":
				link.Name = "󰈙 [" + l.Get("Documentation") + "]"
			case "discuss":
				link.Name = "󰊌 [" + l.Get("Discuss") + "]"
			case "discord":
				link.Name = "󰙯 [" + l.Get("Discord") + "]"
			case "matrix":
				link.Name = "󰊌 [" + l.Get("Matrix") + "]"
			case "bluesky":
				link.Name = " [" + l.Get("Bluesky") + "]"
			case "mastodon":
				link.Name = "󰫑 [" + l.Get("Mastodon") + "]"
			case "donate":
				link.Name = "󱢏 [" + l.Get("Donate") + "]"
			default:
				link.Name = "󰌹 [" + link.Name + "]"
			}
			fmt.Fprintf(&linkSb, " - %s(%s)\n", link.Name, link.URL)
		}
		in += linkSb.String()
		in += "\n"
	}

	var out string

	colorScheme := detectTheme()
	if cfg.UseAccentColor && getDesktop() == "GNOME" {
		r, _ := glamour.NewTermRenderer(
			glamour.WithStyles(getAccentStyle()),
		)
		out, _ = r.Render(in)
	} else {
		out, _ = glamour.Render(in, colorScheme)
	}

	// Renders the output
	fmt.Print(out)
}
