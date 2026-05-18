package main

import (
	"os/exec"
	"strings"

	"charm.land/glamour/v2/ansi"
)

const defaultMargin uint = 2

func stringPtr(s string) *string { return &s }
func boolPtr(b bool) *bool       { return &b }
func uintPtr(u uint) *uint       { return &u }

var accentColors = map[string]string{
	"blue":   "#8be9fd",
	"teal":   "#50fa7b",
	"green":  "#50fa7b",
	"yellow": "#f1fa8c",
	"orange": "#ffb86c",
	"red":    "#ff5555",
	"pink":   "#ff79c6",
	"purple": "#bd93f9",
	"slate":  "#6272a4",
}

func getAccentColor() string {
	defaultColor := accentColors["blue"]
	cmd := exec.Command("gsettings", "get org.gnome.desktop.interface accent-color")
	output, err := cmd.Output()
	if err != nil {
		return defaultColor
	}
	accent := strings.Trim(strings.TrimSpace(string(output)), "'")
	if color, ok := accentColors[accent]; ok {
		return color
	}
	return defaultColor
}

func getDefaultStyle() ansi.StyleConfig {
	accent := getAccentColor()

	return ansi.StyleConfig{
		Document: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockPrefix: "\n",
				BlockSuffix: "\n",
			},
			Margin: uintPtr(defaultMargin),
		},
		BlockQuote: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color:  stringPtr("#f1fa8c"),
				Italic: boolPtr(true),
			},
			Indent:      uintPtr(defaultMargin),
			IndentToken: stringPtr("| "),
		},
		List: ansi.StyleList{
			LevelIndent: defaultMargin,
		},
		Heading: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockSuffix: "\n",
				Color:       stringPtr(accent), // ← accent color on headings
				Bold:        boolPtr(true),
			},
		},
		H1: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: " ",
				Suffix: " ",
				Color:  stringPtr("#FFFFFF"),
			},
		},
		H2: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "## ",
			},
		},
		H3: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "### ",
			},
		},
		H4: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "#### ",
			},
		},
		H5: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "##### ",
			},
		},
		H6: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "###### ",
				Bold:   boolPtr(false),
			},
		},
		Text: ansi.StylePrimitive{
			Color: stringPtr("default"),
		},
		Strikethrough: ansi.StylePrimitive{
			CrossedOut: boolPtr(true),
		},
		Emph: ansi.StylePrimitive{
			Italic: boolPtr(true),
		},
		Strong: ansi.StylePrimitive{
			Bold: boolPtr(true),
		},
		HorizontalRule: ansi.StylePrimitive{
			Color:  stringPtr(accent),
			Format: "\n--------\n",
		},
		Item: ansi.StylePrimitive{
			BlockPrefix: "• ",
		},
		Enumeration: ansi.StylePrimitive{
			BlockPrefix: ". ",
			Color:       stringPtr(accent), // ← accent on enumerations too
		},
		Task: ansi.StyleTask{
			StylePrimitive: ansi.StylePrimitive{},
			Ticked:         "[✓] ",
			Unticked:       "[ ] ",
		},
		Link: ansi.StylePrimitive{
			Color:     stringPtr(accent), // ← accent on links
			Underline: boolPtr(true),
		},
		LinkText: ansi.StylePrimitive{
			Bold: boolPtr(true),
		},
		Image: ansi.StylePrimitive{
			Color:     stringPtr(accent),
			Underline: boolPtr(true),
		},
		ImageText: ansi.StylePrimitive{
			Color:  stringPtr("#ff79c6"),
			Format: "Image: {{.text}} →",
		},
		Code: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color: stringPtr(accent),
			},
		},
		CodeBlock: ansi.StyleCodeBlock{},
		Table: ansi.StyleTable{
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{},
			},
		},
		DefinitionDescription: ansi.StylePrimitive{
			BlockPrefix: "\n🠶 ",
		},
	}
}
