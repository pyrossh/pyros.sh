package components

import (
	"html/template"

	. "github.com/pyros2097/gromer/handlebars"
)

type PageProps struct {
	Title    string        `json:"title"`
	Children template.HTML `json:"children"`
}

func Page(props PageProps) *Template {
	return Html(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
					<meta charset="UTF-8" />
					<meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
					<meta content="utf-8" http-equiv="encoding" />
					<title>{{ props.Title }}</title>
					<meta name="description" content="Hi there, I'm Peter John, a fullstack developer from Bangalore, India." />
					<meta name="author" content="Peter John,pyros.sh" />
					<meta name="keywords" content="peter john,pyros.sh,full stack developer,bangalore,india" />
					<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover" />
					<link rel="icon" href="/assets/icon.png" />
					<link rel="stylesheet" href="/assets/css/styles.css" />
					<script defer src="/assets/js/alpinejs@3.9.6.js"></script>
			</head>
			<body>
			{{ props.Children }}
			</body>
		</html>
	`)
}
