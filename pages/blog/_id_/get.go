package blog_id_

import (
	"bytes"
	"context"
	"html/template"
	"strings"

	. "github.com/pyros2097/gromer/handlebars"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"mvdan.cc/xurls/v2"
	"pyros.sh/assets"
	not_found_404 "pyros.sh/pages/404"
)

var _ = Css(`
	.md-container {
		display: flex;
		flex-direction: column;
		width: 100%;
	}

	.md-container a {
		font-size: 1.125rem;
		line-height: 1.75rem;
		color: var(--blue);
	}

	.md-container pre {
		max-width: 64rem;
		font-family: monospace;
		font-size: 14px;
		border-radius: 16px;
		padding: 16px;
		margin: 8px;
		line-height: 20px;
		overflow-x: scroll;
	}

	.md-container p {
		line-height: 28px;
	}

	.md-container img {
		width: 100%;
	}
`)

var markdown = goldmark.New(
	goldmark.WithExtensions(
		extension.NewLinkify(
			extension.WithLinkifyAllowedProtocols([][]byte{
				[]byte("https:"),
			}),
			extension.WithLinkifyURLRegexp(
				xurls.Strict(),
			),
		),
	),
	goldmark.WithExtensions(
		highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"),
		),
	),
)

func GET(c context.Context, id string) (HtmlContent, int, error) {
	files, err := assets.FS.ReadDir("md")
	if err != nil {
		return HtmlErr(500, err)
	}
	for _, f := range files {
		title := f.Name()[11 : len(f.Name())-3]
		postId := strings.ReplaceAll(title, " ", "-")
		if postId == id {
			source, err := assets.FS.ReadFile("md/" + f.Name())
			if err != nil {
				return HtmlErr(500, err)
			}
			var buf bytes.Buffer
			if err := markdown.Convert(source, &buf); err != nil {
				return HtmlErr(500, err)
			}
			return Html(`
				{{#Page title="Blog"}}
					{{#Header}}{{/Header}}
					{{#Layout}}
						<div class="md-container">
							{{ md }}
						</div>
					{{/Layout}}
					</main>
				{{/Page}}
			`).
				Prop("md", template.HTML(buf.String())).
				Render()
		}
	}
	return not_found_404.GET(c)
}
