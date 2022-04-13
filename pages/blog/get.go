package blog

import (
	"context"

	. "github.com/pyros2097/gromer/handlebars"
)

func GET(c context.Context) (HtmlContent, int, error) {
	return Html(`
		{{#Page title="Blog"}}
			{{#Header}}{{/Header}}
			<main class="w-full h-full">
				<div class="w-full flex flex-1 flex-row justify-center">
					<div class="flex flex-row flex-1 items-center max-w-5xl text-lg font-source p-4 mt-4">
						<div class="flex flex-1 flex-col">
							<div class="flex flex-1 flex-col">
								${data.blog
									.sort((a, b) => (b.uploadedOn > a.uploadedOn ? -1 : 1))
									.map(
										(item) => html
											<div class="flex flex-1 flex-row mt-2">
												<div class="flex-1">
													<div>
														▪
														<a class="ml-2 border-b border-black" href="${item.permaLink}"> ${item.title} </a>
													</div>
												</div>
												<div class="">${item.uploadedOn}</div>
											</div>
										,
									)}
							</div>
						</div>
					</div>
				</div>
			</main>
		{{/Page}}
		`).Render()
}
