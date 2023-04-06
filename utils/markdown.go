package utils

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

var Markdown = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
)
