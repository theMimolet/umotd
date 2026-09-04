package internal

import (
	"math/rand"
	"slices"

	"github.com/leonelquinteros/gotext"
)

func GetRandomMessage(l *gotext.Locale, tags []string) string {

	messages := []string{}

	// General messages accross all Universal Blue systems

	messages = append(messages, []string{
		l.Get("Press `Ctrl + Alt + T` to quickly open a terminal."),
		l.Get("**Did an update break something?** Roll back with `bootc rollback`."),
		l.Get("Search for and install packages with `brew search` and `brew install`. %s will handle the updates automatically.", GetOSName()),
		l.Get("Use `ujust --choose` to see all ujust shortcuts and their associated scripts."),
		l.Get("The **Bazaar** app store is created by **Kolunmi** — [Support their work](%s)", "https://ko-fi.com/kolunmi"),
		l.Get("Switch shells safely: Change your shell in your Terminal's settings (not system-wide) — [Read more](%s)", "https://tim.siosm.fr/blog/2023/12/22/dont-change-defaut-login-shell/"),
		l.Get("Export Distrobox packages to make them appear like native applications — [View documentation](%s)", "https://distrobox.it/usage/distrobox-export/"),
		l.Get("**H.264 hardware acceleration works out of the box** — no tweaks needed!"),
		l.Get("**No Flatpak available?** Use Gear Lever for easy AppImage management — [Check it out](%s)", "appstream://it.mijorus.gearlever"),
		l.Get("**Tailscale is included** — check out [the docs](%s)", "https://tailscale.com/docs/how-to/quickstart"),
		l.Get("**Need to manage Flatpak repositories and data?** — Try [Warehouse](%s)", "appstream://io.github.flattool.Warehouse"),
		l.Get("Use Clapgrep for **powerful folder searches** — [Check it out](%s)", "appstream://de.leopoldluley.Clapgrep"),
		l.Get("**Love our wallpapers?** Explore the full [Universal Blue artwork collection](%s)", "https://docs.projectbluefin.io/artwork/"),
		l.Get("Run `sl` makes a Steam locomotive appear in your terminal! Install it with `brew install sl`"), // idea from bigredsponge
		l.Get("**Are you a developer?** — Try `Developer Mode` for container tooling, virtualization, and IDEs — run `ujust devmode`"),
		l.Get("**Like servers?** Check out [ucore](%s)", "https://github.com/ublue-os/ucore"),
		l.Get("Use `ujust bbrew` to try **BBrew**, a simplified Homebrew package manager."),
		l.Get("Install out **tealdeer** (`brew install tealdeer`) for a quick rundown of command-line tools."),
		// l.Get("`cowsay <something>` makes a cow say something right in your terminal! Install it with `brew install cowsay`"), - homebrew cowsay requires to have perl installed.. but it's not automatically installed with cowsay as a dependency
		l.Get("These messages are part of umotd — [Check it out](%s)", "https://github.com/projectbluefin/umotd"),
		l.Get("**Want to add new messages?** File an issue on the [umotd GitHub repository](%s)", "https://github.com/projectbluefin/umotd"),
	}...)

	// Desktop related messages

	if slices.Contains(tags, "gnome") {
		messages = append(messages, []string{
			l.Get("**GNOME powers your desktop** — [Donate to GNOME](%s)", "https://donate.gnome.org"),
			l.Get("**Control your device from your phone** with the `GSConnect` extension. Look for it in the [Extension Manager](%s)", "appstream://com.mattjakeman.ExtensionManager"),
			l.Get("**Miss the wobbly windows from the early 2000s?** Look for the `Compiz windows effect` extension in the [Extension Manager](%s)", "appstream://com.mattjakeman.ExtensionManager"),
			l.Get("Do you like **the blur**? The `Blur my Shell` extension is pre-installed ! Configure it to your liking in the [Extension Manager](%s)", "appstream://com.mattjakeman.ExtensionManager"),
			l.Get("**Want to see Bluetooth device battery levels?** Look for the `Bluetooth Battery Meter` extension in the [Extension Manager](%s)", "appstream://com.mattjakeman.ExtensionManager"),
			l.Get("Manage desktop extensions using the [Extension Manager](%s)", "appstream://com.mattjakeman.ExtensionManager"),
			l.Get("**Need more control over Flatpak app permissions?** Tweak them to your liking with [Flatseal](%s)", "appstream://com.github.tchx84.Flatseal"),
		}...)
	}

	if slices.Contains(tags, "kde") {
		messages = append(messages, []string{
			l.Get("**KDE powers your desktop** — [Donate to KDE](%s)", "https://kde.org/donate"),
			l.Get("The KDE mascot is a dragon named **Konqi**!"),
			l.Get("**Control your device from your phone** with the `KDE Connect` app — [More info](%s)", "https://kdeconnect.kde.org/"),
			l.Get("**Adjust Flatpak app permissions** in System Settings → Security and Privacy → App Permissions."),
		}...)
	}

	// System related messages

	if slices.Contains(tags, "aurora") {
		messages = append(messages, []string{
			l.Get("**Love Aurora?** Help keep it going by [donating](%s)", "https://docs.getaurora.dev/project-docs/credits"),
			l.Get("**Need technical details?** Check out the [Aurora docs](%s)", "https://docs.getaurora.dev"),
			l.Get("The wallpapers are made by **Chandeleer** — [Support their work](%s)", "https://ko-fi.com/chandeleer"),
			l.Get("Use `ujust rebase-helper` to roll back or switch channels — [See the docs](%s)", "https://docs.getaurora.dev/guides/release-streams/#switching-between-streams"),
			l.Get("**Like dinosaurs?** Try `ujust toggle-dinosaurs` for a Mesozoic surprise!"),
			l.Get("Are you a developer? Try the `Developer Experience` — [Check it out](%s)", "https://docs.getaurora.dev/dx/aurora-dx-intro"),
			l.Get("**%s is not a distro**, this is a custom image built on  Fedora Atomic Desktop technology — [View our mission](%s)", GetOSName(), "https://ublue.it/mission/"),
		}...)
	}
	if slices.Contains(tags, "bazzite") {
		messages = append(messages, []string{
			l.Get("**Love Bazzite?** Help keep it going by [donating](%s)", "https://docs.bazzite.gg/donations/"),
			l.Get("**Did an update break something?** Roll back, pin the previous release, or rebase to an older build — [view the guide](%s)", "https://docs.bazzite.gg/Installing_and_Managing_Software/Updates_Rollbacks_and_Rebasing/rolling_back_system_updates/"),
			l.Get("**Care about game preservation?** Support the Stop Killing Games initiative — [Learn more](%s)", "https://www.stopkillinggames.com/"),
			l.Get("**Prefer Distrobox for package installation!** It’s safer and easier — start with `ujust distrobox`. [More info](%s)", "https://docs.bazzite.gg/Installing_and_Managing_Software/Distrobox/"),
			l.Get("**Installing a non-Steam Windows game?** Use Lutris (pre-installed) for better Wine prefix handling — [View gaming guide](%s)", "https://docs.bazzite.gg/Gaming/"),
			l.Get("Bazzite uses BTRFS for all drives (including MicroSD). Note: **NTFS and exFAT are not supported.**"),
			l.Get("**Looking to set up Waydroid?** — [View our documentation](%s)", "https://docs.bazzite.gg/Installing_and_Managing_Software/Waydroid_Setup_Guide/"),
			l.Get("Bazzite uses ZSTD-compressed BTRFS with automatic deduplication — **more storage space for all your games!**"),
			l.Get("**Managing a large ROM library?** The ROM Properties Page shell extension (pre-installed) adds **thumbnails and metadata** to your files."),
			l.Get("Use the `bazzite-rollback-helper` to roll back or rebase to older builds — [More info](%s)", "https://docs.bazzite.gg/Installing_and_Managing_Software/Updates_Rollbacks_and_Rebasing/rolling_back_system_updates/"),
			l.Get("Use ProtonPlus to install custom Proton versions. We recommend Proton-GE for Steam games and Wine-GE for non-Steam use cases."),
			l.Get("**Want to customize your desktop?** — [Follow our guide](%s)", "https://docs.bazzite.gg/General/Desktop_Environment_Tweaks/"),
			l.Get("**Performance profiling is built-in** — use `sysprof`, `bpftrace`, or `perf` to analyze your system."),
			l.Get("**%s is not a distro** — it’s a custom image built on  Fedora Atomic Desktop — [View our mission](%s)", GetOSName(), "https://ublue.it/mission/"),
			l.Get("**Support indie game preservation and OSS developers** — [Join Hit Save!'s Patreon](%s)", "https://patreon.com/hitsave"),
			l.Get("Use `ujust changelogs` to view a summary of the package changes since the last update."),
		}...)
	}
	if slices.Contains(tags, "bazzite-gnome") {
		messages = append(messages, []string{
			l.Get("**Missing the hot corner?** Move your mouse to the edge at the bottom of your screen to trigger the Activities overview, or re-enable it in Settings."),
		}...)
	}
	if slices.Contains(tags, "bazzite-deck") {
		messages = append(messages, []string{
			l.Get("**Want to install Decky Loader?** Use `ujust setup-decky install`"),
			l.Get("󰓓 The Steam game mode updater also updates Bazzite, Flatpak apps, and Distrobox containers. The changelog is from the Steam client."),
			l.Get("**Install a game with Lutris?** Right-click on it and `Create Steam Shortcut` to play it in 󰓓 Steam game mode."),
			l.Get("**Missing game icons?** Use the `SteamGridDB` plugin for Decky Loader to add missing art — [More info](%s)", "https://github.com/SteamGridDB/decky-steamgriddb"),
			l.Get("**Not sure which games are Linux-compatible?** The `ProtonDB Badges` plugin for Decky Loader adds community-powered compatibility badges to your entire 󰓓 Steam Library — [More info](%s)", "https://github.com/OMGDuke/protondb-decky"),
			l.Get("**Want more security?** Use **Tailscale** and **OpenVPN** (built-in) with the `Tailscale Control` and `TunnelDeck` Decky Loader plugins in 󰓓 Steam game mode."),
			l.Get("**Using a handheld with limited buttons?** 󰓓 Steam game mode supports touch gestures to slide out the menus, thanks to [ChimeraOS](%s)", "https://chimeraos.org/"),
		}...)
	}
	if slices.Contains(tags, "bluefin") {
		messages = append(messages, []string{
			l.Get("**Love Bluefin?** Help keep it going by [donating](%s)", "https://docs.projectbluefin.io/donations"),
			l.Get("Need technical details? — Check out the [Bluefin Administrator's Guide](%s)", "https://docs.projectbluefin.io/administration"),
			l.Get("Use DistroShelf to run apps from other distros in isolated containers — no commitment required — [Check it out](%s)", "appstream://com.ranfdev.DistroShelf"),
			l.Get("Use `ujust rebase-helper` to roll back or switch channels — [See the docs](%s)", "https://docs.projectbluefin.io/administration/#switching-between-streams"),
			l.Get("The `Gradia Capture` extension helps you take better screenshots! Thank Alexander Vanhee — [Support his work!](%s)", "https://ko-fi.com/alexandervanhee"),
			l.Get("Most wallpapers are by **Jacob Schnurr** — [Check out his Etsy](%s)", "https://www.etsy.com/shop/JSchnurrCommissions"),
			l.Get("Are you a developer? Try the `Developer Experience` — [Check it out](%s)", "https://docs.projectbluefin.io/bluefin-dx/"),
			// l.Get("**Bluefin Dakotaraptor** is the newest addition to the Bluefin family! — [Check it out](%s)", "https://docs.projectbluefin.io/dakota/"),
			// l.Get("**Bluefin Dakotaraptor** builds your desktop directly from GNOME! — [Check it out](%s)", "https://docs.projectbluefin.io/dakota/"),
		}...)
	}

	// Dev related messages

	if slices.Contains(tags, "vscode") {
		messages = append(messages, []string{
			l.Get("**Container development is OS-agnostic** — devcontainers work on Linux, macOS, and Windows."),
			l.Get("**Develop with devcontainers** — use `devcontainer.json` to create isolated, reproducible environments! - Activate the workflow with `ujust devmode` and [get started today !](%s)", "https://code.visualstudio.com/docs/devcontainers/tutorial"),
			l.Get("For multi-container setups, use `docker compose` instead of a single devcontainer."),
		}...)
	}

	if slices.Contains(tags, "containers") {
		messages = append(messages, []string{
			l.Get("**%s is your gateway to 󱃾 Kubernetes** — run `kind create cluster` to [get started](%s)", GetOSName(), "https://kind.sigs.k8s.io/"),
			l.Get("**%s is your gateway to Cloud Native** — find your flock at [landscape.cncf.io](%s)", GetOSName(), "https://l.cncf.io"),
			l.Get("%s separates the OS from your development environment — **embrace the cloud-native workflow**!", GetOSName()),
			l.Get("Use `ujust bbrew` and select `k8s-tools` to install Kubernetes tools like `kind` and `kubectl`."),
		}...)
	}

	return messages[rand.Intn(len(messages))]
}
