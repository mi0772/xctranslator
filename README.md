

# xcTranslate

**xcTranslate** is a command-line tool written in Go for translating Apple's `.xcstrings` localization files using Google Translate.

It reads a source `.xcstrings` file (in JSON format), detects all untranslated base strings, translates them into multiple target languages, and generates a new `.xcstrings` file with completed localizations.

---

## Features

- Reads and writes valid `.xcstrings` (Apple localization JSON format)
- Automatic source language detection
- Supports translation into multiple languages
- Preserves placeholders such as `%1$@`, `%lld`, `{name}`, etc.
- Simple CLI usage
- Thread-safe translation engine using mutex protection

---

## Installation

```bash
go install github.com/yourusername/xctranslate@latest
```

---

## Usage

```bash
xctranslate -input=Localizable.xcstrings -output=Translated.xcstrings -langs=it,fr,de
```

### CLI Flags

| Flag      | Description                                |
|-----------|--------------------------------------------|
| `-input`  | Path to the input `.xcstrings` file        |
| `-output` | Path to save the translated `.xcstrings`   |
| `-langs`  | Comma-separated list of target languages (e.g. `it,fr,de`) |

---

## Output format

The translated file will preserve the original structure and include localizations in the form:

```json
"Hello, world!" : {
  "localizations" : {
    "it" : {
      "stringUnit" : {
        "state" : "translated",
        "value" : "Ciao, mondo!"
      }
    },
    "fr" : {
      "stringUnit" : {
        "state" : "translated",
        "value" : "Bonjour le monde!"
      }
    }
  }
}
```

---

## Limitations

- Uses the [unofficial `gtranslate`](https://github.com/bregydoc/gtranslate) package, which scrapes the Google Translate web interface
- Not recommended for high-volume or production use
- Translations are processed sequentially (no concurrent translation engine)

---

## License

MIT License — see [`LICENSE`](LICENSE)

---

## Credits

- [bregydoc/gtranslate](https://github.com/bregydoc/gtranslate)
- Inspired by real-world localization challenges in Apple app development