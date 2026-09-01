# How to translate

## Prerequisites

To translate and test Umotd, you'll need to have the following tools installed on your system:

- [`go`](https://repology.org/project/go/versions)
- [`gettext`](https://repology.org/project/gettext/versions)

## Usage

You can simply run `just translate <language code>` to extract the translatable strings and update the translations file.

Example:

```sh
just translate fr
```

Your generated/updated translation file will be located in the `locales/<language code>` directory.

> If your language already exists, it will get updated automatically. If not, a new translation file will be created for you.

Finally, use your favorite po editor to translate the strings in the `.po` file - like [Poedit](https://flathub.org/en/apps/net.poedit.Poedit), [Gtranslator](https://flathub.org/en/apps/org.gnome.Gtranslator) or [Lokalize](https://flathub.org/en/apps/org.kde.lokalize).

## Testing your translation

You can then use `LANGUAGE=<language code>` in front of the usual command to test your translation, like this:

```sh
# Run with the compiled binary -> needs to be rebuilt after translation changes
LANGUAGE=fr ./umotd
```

```sh
# Run with the source code -> not compiled, so no need to rebuild after translation changes
LANGUAGE=fr go run .
```
