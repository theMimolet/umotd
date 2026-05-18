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
				l.Get("Help keep Aurora alive and healthy, consider [donating](%s)", "https://docs.getaurora.dev/project-docs/credits"),
				l.Get("Need more in-depth technical information? ~ Check out the [Aurora docs](%s)", "https://docs.getaurora.dev"),
				l.Get("Don't forget to check the [release notes](%s)", "https://github.com/ublue-os/aurora/releases"),
				l.Get("Really love our wallpapers? Donate to [Chandeleer](%s)", "https://ko-fi.com/chandeleer"),
			}...)
		case "bazzite":
			tips = append(tips, []string{
				l.Get("Help keep Bazzite alive and healthy, consider [donating](%s)", "https://docs.bazzite.gg/donations/"),
				l.Get("*Update break something?* You can roll back and pin the previous release or rebase by build date ~ [View our guide](%s)", "https://docs.bazzite.gg/Installing_and_Managing_Software/Updates_Rollbacks_and_Rebasing/rolling_back_system_updates/"),
				l.Get("**Protect your video games!** ~ [Visit Stop Killing Games](%s)", "https://www.stopkillinggames.com/"),
				l.Get("It is **always** better to install packages with Distrobox rather than layer them with rpm-ostree. `ujust distrobox` makes it easy! ~ [More info](%s)", "https://docs.bazzite.gg/Installing_and_Managing_Software/Distrobox/"),
				l.Get("*Installing a Windows game that isn't on Steam?* Lutris is pre-installed and recommended for better handling of wine prefixes ~ [View gaming guide](%s)", "https://docs.bazzite.gg/Gaming/"),
				l.Get("BTRFS is used by default for internal drives, and we recommend BTRFS for external drives including MicroSD cards. *NTFS and exFAT are not supported.*"),
				l.Get("*Looking to setup Waydroid?* ~ [View our documentation](%s)", "https://docs.bazzite.gg/Installing_and_Managing_Software/Waydroid_Setup_Guide/"),
				l.Get("Bazzite uses ZSTD compression in BTRFS by default, and deduplicates files across your entire drive. **More space for your games!**"),
				l.Get("*Have a large library of ROMs to manage?* ROM Properties Page shell extension is installed by default and makes it much easier, with thumbnails and additional info for all of your files"),
				l.Get("The `bazzite-rollback-helper` command can guide you for rolling back or rebasing to older Bazzite builds"),
				l.Get("ProtonPlus can be used to install and update custom versions of Proton. We recommend Proton-GE for problematic Steam games."),
			}...)
		case "bluefin":
			tips = append(tips, []string{
				l.Get("Help keep Bluefin alive and healthy, consider [donating](%s)", "https://docs.projectbluefin.io/donations"),
				l.Get("Need more in-depth technical information? ~ Check out the [Bluefin Administrator's Guide](%s)", "https://docs.projectbluefin.io/administration"),
				l.Get("Don't forget to check the [release notes](%s)", "https://github.com/ublue-os/bluefin/releases"),
				l.Get("Use DistroShelf to create pet containers for different distros"),
				l.Get("VSCode comes with the `devcontainers` extension pre-installed - perfect for containerized development"),
				l.Get("Check out `ujust bbrew` for curated selections of development and command line apps"),
			}...)
		case "kde":
			tips = append(tips, []string{
				l.Get("KDE powers your desktop! Donate to [KDE](%s)", "https://kde.org/donate"),
				l.Get("*Want to control your device from your phone?* The KDE Connect app functions with all images ~ [More info](%s)", "https://kdeconnect.kde.org/"),
				l.Get("*Want to customize and theme your desktop?* Tweak your desktop by following our guide ~ [More info](%s)", "https://docs.bazzite.gg/General/Desktop_Environment_Tweaks/"),
			}...)
		case "gnome":
			tips = append(tips, []string{
				l.Get("GNOME makes your desktop! Donate to [GNOME](%s)", "https://donate.gnome.org"),
				l.Get("*Want to control your device from your phone?* The GSConnect extension works with your KDE Connect app ! [More info](%s)", "https://github.com/GSConnect/gnome-shell-extension-gsconnect/wiki"),
				l.Get("*Looking for some nostalgia?* Enable `Compiz windows effect` from the Extension Manager."),
				l.Get("*Missing the top left hot corner?* Apply pressure to the bottom edge of your screen with your mouse. You can also re-enable the hot corner from settings if desired."),
				l.Get("Manage desktop extensions with *Extension Manager*."),
				l.Get("ProtonPlus can be used to install and update custom versions of Proton. We recommend Proton-GE for problematic Steam games and Wine-GE for all other use cases outside of Steam."),
			}...)
		case "dev":
			tips = append(tips, []string{
				l.Get("Like servers? Check out [ucore](%s)", "https://github.com/ublue-os/ucore"),
				l.Get("`ujust bbrew` and select `ide` for our curated selections of development environments (e.g. VSCode, JetBrains, NeoVim)"),
				l.Get("Check out `ujust bbrew` for curated selections of development and command line apps"),
				l.Get("`ujust jetbrains-toolbox` installs JetBrains tools in your home directory, all ready to go!"),
				l.Get("Install tealdeer with brew to have a basic rundown on commands for a given tool: `tldr vim`"),
				l.Get("`ujust bbrew` and select `k8s-tools` that will get you started with Kubernetes development tools like `kind` and `kubectl`"),
				l.Get("Container development is OS-agnostic - your devcontainers work on Linux, macOS, and Windows"),
				l.Get("Performance profiling tools are built-in: try `sysprof`, `bpftrace`, and other debugging tools"),
				l.Get("Use `docker compose` for multi-container development if devcontainers don't fit your workflow"),
				l.Get("%s is your gateway to Kubernetes 󱃾 `kind create cluster` to [get started](%s)", getOSName(), "https://kind.sigs.k8s.io/"),
				l.Get("%s is your gateway to Cloud Native - find your flock at [landscape.cncf.io](%s)", getOSName(), "https://l.cncf.io"),
				l.Get("%s separates the OS from your development environment - embrace the cloud-native workflow !", getOSName()),
				l.Get("Develop with devcontainers - use `devcontainer.json` files in your projects for isolated, reproducible environments ! [Get started here](%s)", "https://code.visualstudio.com/docs/devcontainers/tutorial"),
			}...)
		case "deck":
			tips = append(tips, []string{
				l.Get("*Want to install Decky Loader?* There's a `ujust` command for that! `ujust setup-decky install`"),
				l.Get("The updater built into 󰓓 Steam game mode has been modified to update Bazzite, Flatpaks, and Distrobox containers. *Just ignore the changelog.*"),
				l.Get("*Install a game with Lutris?* Right click on it and `Create steam shortcut` to easily play it in 󰓓 Steam game mode."),
				l.Get("*Games missing icons?* The `SteamGridDB` plugin for Decky Loader makes it easy to add missing art ~ [More info](%s)", "https://github.com/SteamGridDB/decky-steamgriddb"),
				l.Get("*Confused about what games are compatible with Linux?* The `ProtonDB Badges` plugin for Decky Loader adds community-powered game compatibility badges to your entire 󰓓 Steam Library ~ [More info](%s)", "https://github.com/OMGDuke/protondb-decky"),
				l.Get("*Looking for more security?* Tailscale and OpenVPN are built in, and can be controlled from 󰓓 Steam game mode by the `Tailscale Control` and `TunnelDeck` Decky loader plugins."),
				l.Get("*Using a handheld that doesn't have enough buttons?* 󰓓 Steam game mode now supports touch gestures to slide out the menus, thanks to [ChimeraOS](%s)", "https://chimeraos.org/"),
			}...)
		case "ublue":
			tips = append(tips, []string{
				l.Get("Use `Ctrl`-`Alt`-`T` to quickly open a terminal"),
				l.Get("Update break something? You can roll back with `bootc rollback`"),
				l.Get("Use `brew search` and `brew install` to install packages. %s will take care of the updates automatically", getOSName()),
				l.Get("`ujust --choose` will show you each shortcut and the script it's running"),
				l.Get("`ujust rebase-helper` can help you roll back to a specific image, or to a different channel entirely, check the docs for more info"),
				l.Get("`ujust changelogs` shows a summary of the package changes since the last update"),
				l.Get("Like dinosaurs? `ujust toggle-dinosaurs`"),
				l.Get("Support the app store! Donate to  [Bazaar](%s)", "https://github.com/kolunmi/bazaar"),
				l.Get("Switch shells safely: change your shell in [Terminal settings instead of system-wide](%s)", "https://tim.siosm.fr/blog/2023/12/22/dont-change-defaut-login-shell/"),
				l.Get("Packages installed in Distrobox can be exported to appear like any other application ~ [View documentation](%s)", "https://distrobox.it/usage/distrobox-export/"),
				l.Get("*%s isn't a distro*, this is a custom image built on  Fedora Atomic Desktop technology ~ [View our mission](%s)", getOSName(), "https://ublue.it/mission/"),
				l.Get("**Support indie game preservation and OSS developers!** ~ [Join Hit Save!'s Patreon](%s)", "https://patreon.com/hitsave"),
				l.Get("**H.264 hardware acceleration is supported out of the box.** No tweaks necessary!"),
				l.Get("*No Flatpak or distro packaging available?* The Gear Lever app is included to make managing and integrating AppImages easy! ~ [Install it here](%s)", "appstream://it.mijorus.gearlever"),
				l.Get("Tailscale is included, check out [their docs](%s)", "https://tailscale.com/kb/1017/install"),
				l.Get("*Need more control over your Flatpaks?* Check out the Warehouse and Flatseal applications to manage them"),
				l.Get("Open a folder with Clapgrep for super powerful search ~ [Install it here](%s)", "appstream://de.leopoldluley.Clapgrep"),
			}...)
		default:
			// no-op
		}
	}

	return tips[rand.Intn(len(tips))]
}
