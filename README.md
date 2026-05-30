# Umotd

Umotd is a translatable and configurable MOTD - made for your favorite Bootc system !

It's written in **Go** and uses the `gotext` and `glamour` libraries.

**WIP** : The translations are not yet complete and some features still need testing. 

If you want to contribute, you're welcome to submit a pull request or open an issue - it's very much appreciated.

## What's next

Here are features that are planned for the future:

- Shows a warning if the installed image is over a certain amount of time (e.g. one month - configurable)
- Shows a warning if the installed image is unverified (could be disabled via configuration)

> Both are already done in Bazzite, so it's only a matter of porting them over to Umotd.

- Custom configurable tips

## How to try

To try Umotd, first clone the repository and build the binary (you'll need the latest version of Go installed):

```
git clone https://github.com/theMimolet/umotd
cd umotd
go build
./umotd
```

> You can also just drop the `umotd` binary into your usual `/bin` folder and it will work from anywhere.

## How to translate

To translate Umotd, you'll need to have `go`, `gettext` and `xgotext` 

You can get them all using the following commands: 
- `go`: `brew install go`
- `gettext`: `brew install gettext` (actually might be preinstalled)
- `xgotext`: `go install github.com/leonelquinteros/gotext/cli/xgotext@latest`

Then you can simply run the translators.sh script to extract and update the translations.
```
./translators.sh <language code>
```

If your language already exists, it will be updated automatically. If not, a new language file will be created for you at `locales/<language code>/LC_MESSAGES/default.po`.

Use your favorite po editor to translate the strings in the `.po` file. 
You can then use `LANGUAGE=<language code>` in front of the `./umotd` or `go run .` command to test your translation.


## How to configure

Umotd has default built-in configs, but you may be more interested in having a custom config file.

You can create a custom config file at `/etc/umotd/config.json` or `~/.config/umotd/config.json`.

Note : There are built-in presets for tips, commands descriptions and links labels - those are used to get translated strings.

Here's an example config file with all the currently available options (there's the example folder if you want to see concrete use cases):
```json
{
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
  "info-file": "/usr/share/ublue-os/image-info.json",
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
  "symbol": "!",
  "tips-presets": ["aurora", "dev", "kde", "ublue"],
  "use-accent-color": true
}
```
