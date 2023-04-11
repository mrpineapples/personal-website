package utils

import (
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
)

var Markdown = goldmark.New(
	goldmark.WithExtensions(extension.GFM, highlighting.NewHighlighting(
		highlighting.WithStyle("dracula"),
		highlighting.WithFormatOptions(
			html.WithLineNumbers(true),
		),
	)),
)
