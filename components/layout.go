package components

import (
	"html/template"

	. "github.com/pyros2097/gromer/handlebars"
)

type LayoutProps struct {
	Children template.HTML `json:"children"`
}

func Layout(props LayoutProps) *Template {
	return Html(`
		<main class="w-full h-full">
				<div class="w-full flex flex-1 flex-row justify-center">
					<div class="flex flex-row flex-1 items-center max-w-5xl text-lg p-2 mt-4 sm:p-4">
						{{ props.Children }}
					</div>
				</div>
		</main>
	`)
}
