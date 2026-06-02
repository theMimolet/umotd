package main

import (
	"math/rand"

	"github.com/leonelquinteros/gotext"
)

func getRandomTip(l *gotext.Locale, customTips []string, preset ...string) string {

	if len(preset) == 0 && len(customTips) == 0 {
		return ""
	}

	tips := []string{}

	for _, p := range preset {
		switch p {
		case "aurora":
			tips = append(tips, []string{
				l.Get("**Love Aurora?** Help keep it going by [donating](%s)", "https://docs.getaurora.dev/project-docs/credits"),
				l.Get("**Need more in-depth technical information?** — Check out the [Aurora docs](%s)", "https://docs.getaurora.dev"),
				l.Get("Don't forget to check the [release notes](%s)", "https://github.com/ublue-os/aurora/releases"),
				l.Get("The wallpapers are made by **Chandeleer** — [Support their work](%s)", "https://ko-fi.com/chandeleer"),
				l.Get("Use `ujust rebase-helper` to roll back to a specific image or switch channels — [See the docs for details](%s)", "https://docs.getaurora.dev/guides/release-streams/#switching-between-streams"),
				l.Get("**Are you into dinosaurs?** Then try `ujust toggle-dinosaurs`."),
				l.Get("Are you a developer? Then the `Developer Experience` is made for you! — [Check it out](%s)", "https://docs.getaurora.dev/dx/aurora-dx-intro"),
			}...)
		case "bazzite":
			tips = append(tips, []string{
				l.Get("**Love Bazzite?** Help keep it going by [donating](%s)", "https://docs.bazzite.gg/donations/"),
				l.Get("**Did an update break something?** You can roll back, pin the previous release, or rebase to an older build — [view the guide](%s)", "https://docs.bazzite.gg/Installing_and_Managing_Software/Updates_Rollbacks_and_Rebasing/rolling_back_system_updates/"),
				l.Get("**Care about game preservation?** Support the Stop Killing Games initiative — [Find out more](%s)", "https://www.stopkillinggames.com/"),
				l.Get("**Prefer Distrobox over rpm-ostree for installing packages** — It's safer and easier to manage. `ujust distrobox` gets you started! — [More info](%s)", "https://docs.bazzite.gg/Installing_and_Managing_Software/Distrobox/"),
				l.Get("**Installing a Windows game that isn't on Steam?** Lutris is pre-installed and recommended for better handling of Wine prefixes — [View gaming guide](%s)", "https://docs.bazzite.gg/Gaming/"),
				l.Get("Bazzite uses BTRFS for internal and external drives (including MicroSD). Note: **NTFS and exFAT are not supported.**"),
				l.Get("**Looking to set up Waydroid?** — [View our documentation](%s)", "https://docs.bazzite.gg/Installing_and_Managing_Software/Waydroid_Setup_Guide/"),
				l.Get("Bazzite uses ZSTD-compressed BTRFS with automatic deduplication -> **meaning more storage space for your games.**"),
				l.Get("**Managing a large ROM library?** The ROM Properties Page shell extension comes pre-installed, adding **thumbnails and metadata** to all your files."),
				l.Get("The `bazzite-rollback-helper` command can guide you for rolling back or rebasing to older Bazzite builds — [More info](%s)", "https://docs.bazzite.gg/Installing_and_Managing_Software/Updates_Rollbacks_and_Rebasing/rolling_back_system_updates/"),
				l.Get("ProtonPlus can be used to install and update custom versions of Proton. We recommend Proton-GE for problematic Steam games and Wine-GE for all other use cases outside of Steam."),
				l.Get("**Want to customize and theme your desktop?** — [Follow our guide](%s) to get started.", "https://docs.bazzite.gg/General/Desktop_Environment_Tweaks/"),
			}...)
		case "bluefin":
			tips = append(tips, []string{
				l.Get("**Love Bluefin?** Help keep it going by [donating](%s)", "https://docs.projectbluefin.io/donations"),
				l.Get("Need more in-depth technical information? — Check out the [Bluefin Administrator's Guide](%s)", "https://docs.projectbluefin.io/administration"),
				l.Get("Don't forget to check the [release notes](%s)", "https://github.com/ublue-os/bluefin/releases"),
				l.Get("Use DistroShelf to run apps from other distros in isolated containers — no commitment required — [Check it out](%s)", "appstream://com.ranfdev.DistroShelf"),
				l.Get("Use `ujust rebase-helper` to roll back to a specific image or switch channels — [See the docs for details](%s)", "https://docs.projectbluefin.io/administration/#switching-between-streams"),
				l.Get("The **Gradia Capture** extension is made by **Alexander Vanhee** — [Support their work](%s)", "https://ko-fi.com/alexandervanhee"),
				l.Get("Almost all the wallpapers are made by **Jacob Schnurr** — [Check out his Etsy](%s)", "https://www.etsy.com/shop/JSchnurrCommissions"),
				l.Get("Are you a developer? Then the `Developer Experience` is made for you! — [Check it out](%s)", "https://docs.projectbluefin.io/bluefin-dx/"),
			}...)
		case "bazzite-deck":
			tips = append(tips, []string{
				l.Get("**Want to install Decky Loader?** There's a `ujust` command for that! `ujust setup-decky install`"),
				l.Get("󰓓 The Steam game mode updater also updates Bazzite, Flatpak apps, and Distrobox containers. The changelog shown is from the Steam client."),
				l.Get("**Install a game with Lutris?** Right-click on it and `Create steam shortcut` to easily play it in 󰓓 Steam game mode."),
				l.Get("**Games missing icons?** The `SteamGridDB` plugin for Decky Loader makes it easy to add missing art — [More info](%s)", "https://github.com/SteamGridDB/decky-steamgriddb"),
				l.Get("**Confused about what games are compatible with Linux?** The `ProtonDB Badges` plugin for Decky Loader adds community-powered game compatibility badges to your entire 󰓓 Steam Library — [More info](%s)", "https://github.com/OMGDuke/protondb-decky"),
				l.Get("**Looking for more security?** Tailscale and OpenVPN are built-in, and can be controlled from 󰓓 Steam game mode by the `Tailscale Control` and `TunnelDeck` Decky loader plugins."),
				l.Get("**Using a handheld that doesn't have enough buttons?** 󰓓 Steam game mode now supports touch gestures to slide out the menus, thanks to [ChimeraOS](%s)", "https://chimeraos.org/"),
			}...)
		case "bazzite-gnome":
			tips = append(tips, []string{
				l.Get("**Missing the hot corner?** Move your mouse to the bottom edge of the screen to trigger the Activities overview — or re-enable it in Settings."),
			}...)
		case "gnome":
			tips = append(tips, []string{
				l.Get("**GNOME makes your desktop!** — [Donate to GNOME](%s)", "https://donate.gnome.org"),
				l.Get("**Want to control your device from your phone?** Look for the `GSConnect` extension in the [Extension Manager](%s)", "appstream://com.mattjakeman.ExtensionManager"),
				l.Get("**Miss the wobbly windows from the early 2000s?** Look for the `Compiz windows effect` extension in the [Extension Manager](%s)", "appstream://com.mattjakeman.ExtensionManager"),
				l.Get("Manage desktop extensions with [Extension Manager](%s)", "appstream://com.mattjakeman.ExtensionManager"),
				l.Get("**Need more control over your Flatpak apps permissions?** [Check out Flatseal](%s)", "appstream://com.github.tchx84.Flatseal"),
			}...)
		case "kde":
			tips = append(tips, []string{
				l.Get("**KDE powers your desktop!** — [Donate to KDE](%s)", "https://kde.org/donate"),
				l.Get("The KDE mascot is a dragon named `Konqi`!"),
				l.Get("**Want to control your device from your phone?** Then the KDE Connect app is made for you — [More info](%s)", "https://kdeconnect.kde.org/"),
				l.Get("**Want more control over Flatpak app permissions?** Find them under `Security and Privacy → App Permissions` in System Settings."),
			}...)
		case "ublue":
			tips = append(tips, []string{
				l.Get("Use `Ctrl + Alt + T` to quickly open a terminal."),
				l.Get("**Did an update break something?** You can roll back with `bootc rollback`."),
				l.Get("Use `brew search` and `brew install` to install packages. %s will take care of the updates automatically.", getOSName()),
				l.Get("`ujust --choose` will show you each ujust shortcut and the script they're running."),
				l.Get("`ujust changelogs` shows a summary of the package changes since the last update."),
				l.Get("The **Bazaar** app store is made by **Kolunmi** — [Support their work](%s)", "https://ko-fi.com/kolunmi"),
				l.Get("Switch shells safely: change your shell in Terminal settings instead of system-wide — [Read more](%s)", "https://tim.siosm.fr/blog/2023/12/22/dont-change-defaut-login-shell/"),
				l.Get("Packages installed in Distrobox can be exported to appear like any other application — [View documentation](%s)", "https://distrobox.it/usage/distrobox-export/"),
				l.Get("*%s isn't a distro*, this is a custom image built on  Fedora Atomic Desktop technology — [View our mission](%s)", getOSName(), "https://ublue.it/mission/"),
				l.Get("**Support indie game preservation and OSS developers!** — [Join Hit Save!'s Patreon](%s)", "https://patreon.com/hitsave"),
				l.Get("**H.264 hardware acceleration works out of the box** — no tweaks necessary!"),
				l.Get("**No Flatpak available?** Gear Lever is pre-installed for easy AppImage management — [Get it here](%s)", "appstream://it.mijorus.gearlever"),
				l.Get("**Tailscale is included**, check out [their docs](%s)", "https://tailscale.com/kb/1017/install"),
				l.Get("**Need to manage your Flatpak repositories and data?** — [Check out Warehouse](%s)", "appstream://io.github.flattool.Warehouse"),
				l.Get("**Open a folder with Clapgrep** for super powerful search — [Install it here](%s)", "appstream://de.leopoldluley.Clapgrep"),
				l.Get("**Do you love our wallpapers?** Check out the full [Universal Blue artwork collection](%s)", "https://docs.projectbluefin.io/artwork/"),
			}...)
		case "ublue-dev":
			tips = append(tips, []string{
				l.Get("**Do you like servers?** Then check out [ucore](%s)", "https://github.com/ublue-os/ucore"),
				l.Get("Check out `ujust bbrew` for a curated selection of IDEs, development and command line apps."),
				l.Get("`ujust jetbrains-toolbox` installs JetBrains tools in your home directory, all ready to go!"),
				l.Get("Install `tealdeer` with Homebrew to have a basic rundown on commands for your tools."),
				l.Get("`ujust bbrew` and select `k8s-tools` that will get you started with Kubernetes development tools like `kind` and `kubectl`."),
				l.Get("**Container development is OS-agnostic** — your devcontainers work on Linux, macOS, and Windows."),
				l.Get("**Performance profiling is built-in** — try `sysprof`, `bpftrace`, or `perf` to dig into what your system is doing."),
				l.Get("Prefer `docker compose` for multi-container setups where a single devcontainer isn't enough."),
				l.Get("**%s is your gateway to Kubernetes** 󱃾 `kind create cluster` to [get started](%s)", getOSName(), "https://kind.sigs.k8s.io/"),
				l.Get("**%s is your gateway to Cloud Native** — find your flock at [landscape.cncf.io](%s)", getOSName(), "https://l.cncf.io"),
				l.Get("**%s separates the OS from your development environment** — take full advantage of the cloud-native workflow!", getOSName()),
				l.Get("**Develop with devcontainers** — use `devcontainer.json` files in your projects for isolated, reproducible environments! [Get started here](%s)", "https://code.visualstudio.com/docs/devcontainers/tutorial"),
			}...)
		case "default":
			tips = append(tips, []string{
				l.Get("The Linux penguin is named `Tux`!"),
				l.Get("Maybe **the real %s** was the friends we made along the way.", getOSName()),
				l.Get("**Your distro is valid!** It's your computer, choose what works best for you.👍"),
			}...)
		}
	}

	tips = append(tips, customTips...)

	return tips[rand.Intn(len(tips))]
}
