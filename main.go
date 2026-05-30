package main

import (
	"fmt"
	"strings"

	"charm.land/glamour/v2"
	"github.com/leonelquinteros/gotext"
)

func main() {

	// Loads the locale based on the system's locale
	locale := detectLocale(localesFS)
	l := gotext.NewLocaleFSWithPath(locale, localesFS, "locales")
	l.AddDomain("default")

	// Loads the configuration from the system's config file
	cfg := getConfig()

	// Gets the image info and OS name
	in := "# " + l.Get("Welcome to %s", getOSName()) + "\n"
	if imageInfo := getImageInfo(cfg.InfoFile); imageInfo.ImageRef != "" || imageInfo.ImageTag != "" {
		in += " 󱋩 `" + imageInfo.ImageRef + ":" + imageInfo.ImageTag + "` \n"
	} else if isBootcSystem() {
		in += " 󱋩 `" + l.Get("Unknown system") + "` \n"
	}

	// Gets the Greenboot status
	if greenboot := getGreenbootInfo(); greenboot != "" {
		in += "\n 󰟀  " + l.Get("Boot Status") + ": "
		if greenboot == "healthy" {
			in += l.Get("Healthy") + " 󰄳"
		} else {
			in += greenboot
		}
		in += " \n"
	}

	// Command list
	in += " |  " + l.Get("Command") + " | " + l.Get("Description") + " | \n"
	in += "| ------------ | ----------- |\n"
	var cmdSb strings.Builder
	for _, cmd := range cfg.Commands {
		switch cmd.Desc {
		case "cmd_list":
			cmd.Desc = l.Get("List all available commands")
		case "motd_toggle":
			cmd.Desc = l.Get("Toggle this banner on/off")
		case "sys_info":
			cmd.Desc = l.Get("View system information")
		case "cli_pkg":
			cmd.Desc = l.Get("Manage command line packages")
		case "terminal_bling":
			cmd.Desc = l.Get("Enable terminal bling")
		}
		fmt.Fprintf(&cmdSb, "| `%s`  | %s |\n", cmd.Cmd, cmd.Desc)
	}
	in += cmdSb.String()
	in += "\n"

	// Gets a random tip
	in += getRandomTip(cfg.TipsPresets...) + "\n\n"

	// Gets the links
	var linkSb strings.Builder
	for _, link := range cfg.Links {
		switch link.Name {
		case "issues":
			link.Name = "󰊤 [" + l.Get("Report an issue") + "]"
		case "docs":
			link.Name = "󰈙 [" + l.Get("Documentation") + "]"
		case "discord":
			link.Name = "󰙯 [" + l.Get("Discord") + "]"
		case "bluesky":
			link.Name = " [" + l.Get("Bluesky") + "]"
		case "discuss":
			link.Name = "󰊌 [" + l.Get("Discuss") + "]"
		case "mastodon":
			link.Name = "󰫑 [" + l.Get("Mastodon") + "]"
		default:
			link.Name = "󰌹 [" + link.Name + "]"
		}
		fmt.Fprintf(&linkSb, " - %s(%s)\n", link.Name, link.URL)
	}
	in += linkSb.String()
	in += "\n"

	// Renders the output
	out, _ := glamour.Render(in, detectTheme())
	fmt.Print(out)
}
