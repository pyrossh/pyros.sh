package components

import (
	. "github.com/pyros2097/gromer/handlebars"
)

var _ = Css(`
	.tag {
		background-color: var(--black-light);
		color: white;
		display: inline-block;
		padding-left: 8px;
		padding-right:8px;
		text-align:center;
	}
`)

type TagProps struct {
	Text string `json:"text"`
}

func Tag(props TagProps) *Template {
	return Html(`
		<span class="tag">{{ props.Text }}</span>
	`)
}
