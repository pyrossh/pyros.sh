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

	.md-container img {
		width: 100%;
	}

	.md-container p {
		line-height: 28px;
	}

	.md-container .title-container {
		display: flex;
		flex: 1;
		flex-direction: column;
		font-family: serif;
	}

	.md-container .title-container h1 {
		color: var(--black-light);
		margin: 0;
		line-height: 3rem;
	}

	.md-container .title-container h2 {
		color: var(--yellow-dark);
		font-size: 1.20rem;
		font-weight: 500;
		margin-top: 20px;
		margin-bottom: 20px;
	}

	.md-container .title-container h3 {
		color: var(--yellow-dark);
		font-size: 1.20rem;
		font-weight: 500;
		margin: 0;
	}

	.md-container .title-container .date {
		display: flex;
		flex: 1;
		justify-content: flex-start;
	}

	.md-container .tags-container {
		margin-top: var(--space-4);
		margin-bottom: var(--space-4);
	}

	@media (min-width: 640px) {
		.md-container img {
			width: auto;
		}

		.md-container .title-container {
			flex-direction: row;
		}

		.md-container .title-container .date {
			flex-direction: row;
			justify-content: flex-end;
			margin-right: var(--space-10);
		}

		.md-container .tags-container {
			margin-top: var(--space-1);
			margin-bottom: var(--space-1);
		}
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
			description := metaData["description"].(string)
			image := metaData["image"].(string)
			date := metaData["date"].(string)
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
							<div class="tags-container">
								{{#each tags as |tag|}}
									{{#Tag text=tag}}
									{{/Tag}}
								{{/each}}
							</div>
							{{ md }}
						</div>
					{{/Layout}}
					</main>
				{{/Page}}
			`).
				Prop("title", title).
				Prop("description", description).
				Prop("image", utils.GetImageUrl(image)).
				Prop("date", date).
				Prop("keywords", strings.Join(keywords, ",")).
				Prop("md", template.HTML(buf.String())).
				Prop("url", utils.GetUrl(c)).
				Prop("tags", tags).
				Render()
		}
	}
	return not_found_404.GET(c)
}
