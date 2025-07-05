---
date: "2025-07-04"
title: "Syntax Highlighting Code with Goldmark"
description: "How to style code blocks in Go using the Goldmark library and its Goldmark Highlighting extension."
---

[Goldmark](https://github.com/yuin/goldmark) is a powerful markdown processor, which is used on this very site. You can easily add syntax highlighting using the [Goldmark-highlighting extension](https://github.com/yuin/goldmark-highlighting). Simply create a custom markdown processor to get started.

```go
import (
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

var Markdown = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"),
		),
	),
)
```

Once you've set up the custom processor, you can choose from any of the supported themes, which can be found on the extension's [repository](https://github.com/yuin/goldmark-highlighting). However, one feature that I wanted was to add the language of the code block as a data attribute to the generated markup. This would allow me to create a simple header for each code block.

After browsing through the Goldmark documentation and Github repository, I discovered the `WithWrapperRenderer` method that runs a function for every code block. This was exactly what I needed! Here's how I got it working:

```go
import (
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/util"
)

func wrapperRenderer(w util.BufWriter, ctx highlighting.CodeBlockContext, entering bool) {
	language, ok := ctx.Language()
	lang := string(language)
	// code block with a language
	if ok && lang != "" {
		if entering {
			w.WriteString("<div data-lang=" + lang + ">")
		} else {
			w.WriteString(`</div>`)
		}
		return
	}

	// code block with no language specified
	if language == nil {
		if entering {
			w.WriteString("<pre><code>")
		} else {
			w.WriteString(`</code></pre>`)
		}
	}
}

var Markdown = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"),
			highlighting.WithWrapperRenderer(wrapperRenderer), // updated code here
		),
	),
)
```

Now all code blocks will be wrapped with a `div` that contains a `data-lang` attribute specifying the language. This allows you to use this value on your frontend as you see fit!
