package components

import (
	"html/template"

	. "github.com/pyros2097/gromer/handlebars"
)

var _ = Css(`
	main {
		width: 100%;
		height: 100%;
		font-family: serif;
	}

	main > div {
		display: flex;
		width: 100%;
		flex: 1 1 0%;
		flex-direction: row;
		justify-content: center;
	}

	main > div > div {
		display: flex;
		width: 100%;
		flex: 1 1 0%;
		flex-direction: row;
		align-items: center;
		max-width: var(--5xl);
		font-size: 1.125rem;
		line-height: 1.75rem;
		padding: var(--space-2);
		margin-top: var(--space-4);
	}

	@media (min-width: 640px) {
		main > div > div {
			padding: var(--space-4);
		}
	}
`)

type LayoutProps struct {
	Children template.HTML `json:"children"`
}

func Layout(props LayoutProps) *Template {
	return Html(`
		<main>
				<div>
					<div>
						{{ props.Children }}
					</div>
				</div>
		</main>
	`)
}
