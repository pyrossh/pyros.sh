package not_found_404

import (
	"context"

	. "github.com/pyros2097/gromer/handlebars"
)

func GET(c context.Context) (HtmlContent, int, error) {
	return Html(`
		{{#Page title="Page Not Found" description="Page Not Found"}}
			{{#Header}}{{/Header}}
			{{#Layout}}
				<h1 class="text-5xl">Page Not Found</h1>
				<h2 class="text-lg mt-20">
					<a class="underline" href="/">Go Back</a>
				</h1>
			{{/Layout}}
		{{/Page}}
		`).RenderWithStatus(404)
}
