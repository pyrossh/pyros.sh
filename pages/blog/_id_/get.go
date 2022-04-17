package blog_id_

import (
	"bytes"
	"context"
	"html/template"
	"strings"

	. "github.com/pyros2097/gromer/handlebars"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"mvdan.cc/xurls/v2"
	"pyros.sh/assets"
	not_found_404 "pyros.sh/pages/404"
	"pyros.sh/utils"
)

var _ = Css(`
	.md-container {
		display: flex;
		flex-direction: column;
		width: 100%;
	}

	.md-container h1 {
		color: var(--blue);
		margin: 0;
	}

	.md-container h2 {
		color: var(--blue-light);
		font-size: var(--f-large);
		font-weight: 500;
		margin-top: 20px;
		margin-bottom: 20px;
	}

	.md-container h3 {
		color: var(--blue-light);
		font-size: var(--f-large);
		font-weight: 500;
		margin: 0;
	}

	.md-container pre {
		max-width: 64rem;
		font-family: monospace;
		font-size: 14px;
		border-radius: 16px;
		padding: 16px;
		margin: 8px;
		line-height: 20px;
		overflow-x: auto;
	}

	.md-container p {
		line-height: 28px;
	}

	.md-container .title-container {
		display: flex;
		flex: 1;
		font-family: serif;
	}

	.md-container .title-container .date {
		display: flex;
		flex: 1;
		justify-content: flex-end;
		margin-right: var(--space-10);
	}
`)

var markdown = goldmark.New(
	goldmark.WithExtensions(
		meta.Meta,
	),
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
			context := parser.NewContext()
			if err := markdown.Convert(source, &buf, parser.WithContext(context)); err != nil {
				return HtmlErr(500, err)
			}
			metaData := meta.Get(context)
			description := metaData["description"]
			image := metaData["image"]
			date := metaData["date"]
			tags := metaData["tags"].([]interface{})
			keywords := []string{}
			keywords = append(keywords, "pyros.sh", "pyrossh", "blog post")
			keywords = append(keywords, strings.Split(id, "-")...)
			for _, t := range tags {
				keywords = append(keywords, t.(string))
			}
			return Html(`
				{{#Page url=url title=title description=description keywords=keywords image=image}}
					{{#Header}}{{/Header}}
					{{#Layout}}
						<div class="md-container">
							<div class="title-container">
								<div>
									<h1>{{ title }}</h1>
									<h2>{{ description }}</h2>
								</div>
								<div class="date">
									<h3>{{ date }}</h3>
								</div>
							</div>
							{{ md }}
						</div>
					{{/Layout}}
					</main>
				{{/Page}}
			`).
				Prop("title", title).
				Prop("description", description).
				Prop("image", image).
				Prop("date", date).
				Prop("keywords", strings.Join(keywords, ",")).
				Prop("md", template.HTML(buf.String())).
				Prop("url", utils.GetUrl(c)).
				Render()
		}
	}
	return not_found_404.GET(c)
}
