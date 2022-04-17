package blog

import (
	"context"
	"sort"
	"strings"

	. "github.com/pyros2097/gromer/handlebars"

	"pyros.sh/assets"
	"pyros.sh/utils"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

func GET(c context.Context) (HtmlContent, int, error) {
	files, err := assets.FS.ReadDir("md")
	if err != nil {
		return HtmlErr(500, err)
	}
	posts := []*Post{}
	for _, f := range files {
		title := f.Name()[11 : len(f.Name())-3]
		id := strings.ReplaceAll(title, " ", "-")
		posts = append(posts, &Post{
			ID:    id,
			Title: title,
			Date:  f.Name()[0:10],
		})
	}
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date > posts[j].Date
	})
	return Html(`
		{{#Page url=url title="Blog" description="List of posts" keywords="pyros.sh,pyrossh,blog"}}
			{{#Header}}{{/Header}}
			{{#Layout}}
				<div class="flex flex-1 flex-col">
					<div class="flex flex-1 flex-col">
						{{#each posts as |post| }}
							<div class="flex flex-1 flex-row mt-2">
								<div class="flex-1">
									<div>
										▪
										<a class="ml-2 border-b border-black" href="/blog/{{ post.ID }}"> {{ post.Title }} </a>
									</div>
								</div>
								<div class="">{{ post.Date }}</div>
							</div>
						{{/each}}
					</div>
				</div>
			{{/Layout}}
		{{/Page}}
		`).
		Prop("posts", posts).
		Prop("url", utils.GetUrl(c)).
		Render()
}
