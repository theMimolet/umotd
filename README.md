# Umotd

*Umotd is a translatable and configurable MOTD made for your favorite Linux systems!*

**WIP** : Some features still need testing.

**Contributions are welcome!** If you want to contribute, you're welcome to submit a pull request or open an issue - it's very much appreciated ❤️

## Roadmap

Here are features that are planned for the future:

- Auto-detect or config option to prevent broken symbols
- Shows a warning if the installed system image is over a certain amount of time (e.g. one month - configurable)
- Shows a warning if the installed system image is unverified (could be disabled via configuration)

> Both are already done in Bazzite, so it's only a matter of porting them over to umotd.

## How to try

### Install it with Homebrew

If you have Homebrew installed on your system, you can install Umotd with the following command:
```
brew install themimolet/tap/umotd
```

It's the recommended way to install Umotd on your system, as it will automatically update via Homebrew.

### Download it from the releases page

You can download the latest release from the [releases page](https://github.com/theMimolet/umotd/releases).

You can then rename it to `umotd` and place it in your usual `/bin` folder.

> Note: You won't receive automatic updates, you will have to download Umotd at each new release.

### Compile from source

You'll need to have [`go`](https://repology.org/project/go/versions) installed on your system to compile Umotd from source.

Then you'll have to simply clone the repository and then build the binary:
```
git clone https://github.com/theMimolet/umotd
cd umotd
go build
./umotd
```

You'll then have the `umotd` binary in the current directory, which you can just drop into your usual `/bin` folder and it will work without any further setup (except for the configuration file if you want to customize it).

## How to translate

### Prerequisites

To translate and test Umotd, you'll need to have the following tools installed on your system:

- [`go`](https://repology.org/project/go/versions)
- [`gettext`](https://repology.org/project/gettext/versions)
- [`xgotext`](https://pkg.go.dev/github.com/leonelquinteros/gotext/cli/xgotext) - `go install github.com/leonelquinteros/gotext/cli/xgotext@latest`

### Usage

You can simply run the translators.sh script to extract the translatable strings and update the translations files.

```
./translators.sh <language code>
```

Your translation files are located in the `locales/<language code>/LC_MESSAGES/default.po` directory.

> If your language already exists, it will be updated automatically. If not, a new language file will be created for you.

Finally, use your favorite po editor to translate the strings in the `.po` file - like [Poedit](appstream://net.poedit.poedit), [Gtranslator](appstream://org.gnome.Gtranslator) or [Lokalize](appstream://org.kde.lokalize).

### Testing your translation

You can then use `LANGUAGE=<language code>` in front of the usual command to test your translation, like this:

```sh
# Run with the compiled binary - needs to be rebuilt after translation changes
LANGUAGE=fr ./umotd
```

```sh
# Run with the source code - not compiled, so no need to rebuild after translation changes
LANGUAGE=fr go run .
```

## How to configure Umotd

Umotd has a default built-in look, but it's actually made to be highly customizable.
If you're managing a custom system, it might interest you.

### Where to put your config

You can create a custom config file at `/etc/umotd/config.json` (system-wide) or `~/.config/umotd/config.json` (user-specific).

### Translations ?!

Umotd supports translations for specific tips, command descriptions and link names.
They have specific names / codes that are used to get translated strings.

Any other option not listed won't be translated.

### Breaking down the configuration file

Here's a breakdown of the config file options - there's also the example folder if you want to see concrete use cases.

#### Commands

This option allows you to define a list of commands to display in the MOTD.

Here are the unique codes you can use to get translated strings for command descriptions : 

- `cmd_list`: "List of available commands"
- `cli_pkg`: "Manage command line packages"
- `term_bling`: "Enable terminal bling"
- `motd_toggle`: "Toggle this banner on/off" (there are no built-in commands for this)
- `sys_info`: "View system info"
- `man_upd`: "Manually update the system"

```json
{
  "commands": [
    {
      "cmd": "ujust aurora-cli",
      "desc": "term_bling"
    },
    {
      "cmd": "fastfetch",
      "desc": "sys_info"
    },
    {
      "cmd": "brew help",
      "desc": "cli_pkg"
    },
    {
      "cmd": "cowsay",
      "desc": "Display a cow saying something"
    }
  ]
}
```

#### Links

This option allows to add custom links to the MOTD.

There are unique names you can use to get a translated name for the link :

- `website` : "Website"
- `issues` : "Report an issue"
- `docs` : "Documentation"
- `discuss` : "Discuss"
- `discord` : "Discord"
- `matrix` : "Matrix"
- `bluesky` : "Bluesky"
- `mastodon` : "Mastodon"
- `donate` : "Donate"

```json
{
  "links": [
    {
      "name": "issues",
      "url": "https://issues.bazzite.gg/"
    },
    {
      "name": "docs",
      "url": "https://docs.bazzite.gg/"
    },
    {
      "name": "discord",
      "url": "https://discord.gg/bazzite"
    },
    {
      "name": "bluesky",
      "url": "https://bluesky.bazzite.gg/"
    },
    {
      "name": "discuss",
      "url": "https://github.com/ublue-os/aurora/discussions"
    },
    {
      "name": "Custom Link (it won't get translated)",
      "url": "https://www.innersloth.com/games/among-us/"
    }
  ]
}
```

#### Prefix and Suffix

These options allow to customize the prefix and suffix of the welcome message.

```json
{
  "prefix": "> ",
  "suffix": " !"
}
```

Example:

```
`> Welcome to UBlue !`
```

#### Tips

The `tips` option allows to add custom tips to the MOTD.
But if you want to use the predefined and translated tips included in umotd, you can use the `tips-presets` option.

Currently, there are the following presets available:

- `aurora`
- `bazzite`
- `bazzite-deck`
- `bazzite-gnome`
- `bluefin`
- `default`
- `gnome`
- `kde`
- `ublue`
- `ublue-dev`

```json
{
  "tips": [
    "This is a custom tip by yours truly ! :D",
    "This is another custom tip (they won't get translated)"
  ],
  "tips-presets": [
    "gnome",
    "bluefin",
    "ublue",
    "ublue-dev"
  ]
}
```

#### Use Accent Color

This option allows umotd to use the accent color of the system.

> Note: It's only available for the GNOME desktop as it relies on `gsettings`.

```json
{
  "use-accent-color": true
}
```

#### Info File

This option allows to redirect the info file path to a custom location.

> Note: This option is specifically tailored for bootc / ublue images, because it uses the image info file of their system to display information about the system image the user is currently running.


```json
{
  "info-file": "/usr/share/ublue-os/image-info.json"
}
```
