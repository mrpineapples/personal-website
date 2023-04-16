package utils

import (
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/util"
)

var Markdown = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"),
			highlighting.WithWrapperRenderer(wrapperRenderer),
		),
	),
)

var langMap = map[string]string{
	"css":        "CSS",
	"go":         "Go",
	"html":       "HTML",
	"javascript": "JavaScript",
	"js":         "JavaScript",
	"json":       "JSON",
	"jsx":        "JavaScript",
	"sql":        "SQL",
	"ts":         "TypeScript",
	"tsx":        "TypeScript",
	"typescript": "TypeScript",
}

func wrapperRenderer(w util.BufWriter, ctx highlighting.CodeBlockContext, entering bool) {
	language, isLanguageSet := ctx.Language()
	lang := string(language)
	// code block with a language
	if isLanguageSet && lang != "" {
		formattedLang, ok := langMap[lang]
		if !ok {
			formattedLang = lang
		}

		if entering {
			w.WriteString("<div data-lang=" + formattedLang + ">")
		} else {
			w.WriteString(`</div>`)
		}
		return
	}

	// Code block with no language specified
	if language == nil {
		if entering {
			w.WriteString("<pre><code>")
		} else {
			w.WriteString(`</code></pre>`)
		}
	}
}
