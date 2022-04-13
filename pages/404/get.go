package not_found

import (
	"context"

	. "github.com/pyros2097/gromer/handlebars"
)

func GET(c context.Context) (HtmlContent, int, error) {
	return Html(`
		{{#Page title="Page Not Found"}}
			{{#Header}}{{/Header}}
			<main class="flex flex-1 flex-col mt-20 items-center">
				<h1 class="text-5xl">Page Not Found</h1>
				<h2 class="text-lg mt-20">
					<a class="underline" href="/">Go Back</a>
				</h1>
			</main>
		{{/Page}}
		`).Render()
}
