# Umotd

Umotd is a proof of concept for a program-based MOTD that handles translations.

It's written in **Go** and uses the `gotext` and `glamour` libraries.

**WIP** : The translations are not yet complete and the style is pretty much a work in progress - use a dark themed terminal for the best experience for now. 

If you want to contribute, you're welcome to submit a pull request or open an issue - it's very much appreciated.

## How to try

To try Umotd, first clone the repository and build the binary (you'll need the latest version of Go installed):

```
git clone https://github.com/theMimolet/umotd
cd umotd
go build
./umotd
```

## How to translate

To translate Umotd, you'll need to have go, gettext and xgotext (`go install github.com/leonelquinteros/gotext/cli/xgotext@latest`)

Then you can simply run the translators.sh script to extract and update the translations.
```
sh ./translators.sh <language code>
```

If your language already exists, it will be updated automatically. If not, a new language file will be created for you at `locales/<language code>/LC_MESSAGES/default.po`.

Use your favorite po editor to translate the strings in the `.po` file. 
You can then use `LANGUAGE=<language code>` in front of the `./umotd` or `go run .` command to test your translation.


## How to configure

Umotd has default built-in configs, but you may be more interested in having a custom config file.

You can create a custom config file at `/etc/umotd/config.json` or `~/.config/umotd/config.json`.

Note : There are built-in presets for tips, commands descriptions and links labels - those are used to get translated strings.

Here's an example config file with all the currently available options (as of May 18th):
```json
{
  "info_file": "/usr/share/ublue-os/image-info.json",
  "commands": [
    {
      "cmd": "ujust --choose",
      "desc": "cmd_list"
    },
    {
      "cmd": "ujust toggle-user-motd",
      "desc": "motd_toggle"
    },
    {
      "cmd": "ujust aurora-cli",
      "desc": "terminal_bling"
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
      "desc": "Custom command field (won't be translated)"
    }
  ],
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
      "name": "Custom Link (won't be translated)",
      "url": "https://www.innersloth.com/games/among-us/"
    }
  ],
  "tips_presets": ["aurora", "dev", "kde", "ublue"]
}
```
