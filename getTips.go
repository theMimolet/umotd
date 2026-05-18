package main

import (
	"math/rand"

	"github.com/leonelquinteros/gotext"
)

func getRandomTip(preset ...string) string {

	if len(preset) == 0 {
		return ""
	}

	locale := detectLocale(localesFS)
	l := gotext.NewLocaleFSWithPath(locale, localesFS, "locales")
	l.AddDomain("default")

	tips := []string{}

	for _, p := range preset {
		switch p {
		case "aurora":
			tips = append(tips, []string{
				l.Get("Aurora is your gateway to Kubernetes 󱃾 `kind create cluster` to [get started]") + "(https://kind.sigs.k8s.io/)",
				l.Get("Aurora is your gateway to Cloud Native - find your flock at [landscape.cncf.io]") + "(https://l.cncf.io)",
				l.Get("Need more in-depth technical information?~Check out the [Aurora docs]") + "(https://docs.getaurora.dev)",
				l.Get("Don't forget to check the [release notes]") + "(https://github.com/ublue-os/aurora/releases)",
				l.Get("Help keep Aurora alive and healthy, consider [donating]") + "(https://docs.getaurora.dev/project-docs/credits)",
				l.Get("Aurora separates the OS from your development environment - embrace the cloud-native workflow"),
				l.Get("Really love our wallpapers? Donate to [Chandeleer]") + "(https://ko-fi.com/chandeleer)",
			}...)
		case "kde":
			tips = append(tips, []string{
				l.Get("KDE powers your desktop! Donate to [KDE]") + "(https://kde.org/donate)",
			}...)
		case "dev":
			tips = append(tips, []string{
				l.Get("Use `Ctrl`-`Alt`-`T` to quickly open a terminal"),
				l.Get("Like servers? Check out [ucore](https://github.com/ublue-os/ucore)"),
				l.Get("Tailscale is included, check out [their docs](https://tailscale.com/kb/1017/install)"),
				l.Get("Check out `ujust bbrew` for curated selections of development and command line apps"),
				l.Get("`ujust jetbrains-toolbox` installs JetBrains tools in your home directory, all ready to go!"),
				l.Get("Try tealdeer to have a basic rundown on commands for a given tool: `tldr vim`"),
				l.Get("`ujust bbrew` and select `k8s-tools` that will get you started with Kubernetes development tools like kind and kubectl"),
				l.Get("Container development is OS-agnostic - your devcontainers work on Linux, macOS, and Windows"),
			}...)
		case "ublue":
			tips = append(tips, []string{
				l.Get("Update break something? You can roll back with `bootc rollback`"),
				l.Get("Use `brew search` and `brew install` to install packages. Your system will take care of the updates automatically"),
				l.Get("`ujust --choose` will show you each shortcut and the script it's running"),
				l.Get("`ujust rebase-helper` can help you roll back to a specific image, or to a different channel entirely, check the docs for more info"),
				l.Get("`ujust changelogs` shows a summary of the package changes since the last update"),
				l.Get("Like dinosaurs? `ujust toggle-dinosaurs`"),
				l.Get("Support the app store! Donate to  [Bazaar]") + "(https://github.com/kolunmi/bazaar)!",
				l.Get("Switch shells safely: change your shell in [Terminal settings instead of system-wide]") + "(https://tim.siosm.fr/blog/2023/12/22/dont-change-defaut-login-shell/)",
				l.Get("Open a folder with Clapgrep (Found in the Bazaar App Store) for super powerful search"),
				l.Get("Check out `ujust bbrew` for curated selections of development and command line apps"),
			}...)
		default:
			// no-op
		}
	}

	return tips[rand.Intn(len(tips))]
}
