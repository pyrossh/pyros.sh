package components

import (
	"html/template"

	. "github.com/pyros2097/gromer/handlebars"
)

var _ = Css(`
	*,
	::before,
	::after {
		box-sizing: border-box;
		/* 1 */
		border-width: 0;
		/* 2 */
		border-style: solid;
		/* 2 */
		border-color: #e5e7eb;
		/* 2 */
	}

	:root {
		--black: #1a1a1a;
		--yellow: #f1fa8c;
		--blue: #0645ad;
		--slider-bg: #f0ede2;
		--space-2: 0.5rem;
		--space-3: 0.75rem;
		--space-4: 1rem;
		--space-10: 2.5rem;
		--5xl: 64rem;
	}

	a {
		color: inherit;
		text-decoration: inherit;
	}

	ol,
	ul,
	menu {
		list-style: none;
		margin: 0;
		padding: 0;
	}

	html {
		height: 100%;
		width: 100%;
		line-height: 1.5;
		-webkit-text-size-adjust: 100%;
		-moz-tab-size: 4;
		-o-tab-size: 4;
		tab-size: 4;
	}

	body {
		margin: 0;
		display: flex;
		height: 100%;
		min-height: 100vh;
		width: 100%;
		flex: 1 1 0%;
		flex-direction: column;
		background-color: rgb(255 255 255);
		font-size: 1rem;
		line-height: 1.5rem;
		font-weight: 400;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;
		direction: ltr;
		font-synthesis: none;
		text-rendering: optimizeLegibility;
	}

	.link {
		font-size: 1.125rem;
		line-height: 1.75rem;
		color: var(--blue);
	}

	pre {
		max-width: 64rem;
		font-family: monospace;
		font-size: 14px;
		border-radius: 16px;
		padding: 16px;
		margin: 8px;
		line-height: 20px;
		overflow-x: scroll;
	}

	.fixed {
		position: fixed
	}

	.relative {
		position: relative
	}

	.my-3 {
		margin-top: 0.75rem;
		margin-bottom: 0.75rem
	}

	.my-6 {
		margin-top: 1.5rem;
		margin-bottom: 1.5rem
	}

	.mt-4 {
		margin-top: 1rem
	}

	.mr-8 {
		margin-right: 2rem
	}

	.mb-4 {
		margin-bottom: 1rem
	}

	.mt-8 {
		margin-top: 2rem
	}

	.mt-6 {
		margin-top: 1.5rem
	}

	.ml-6 {
		margin-left: 1.5rem
	}

	.ml-2 {
		margin-left: 0.5rem
	}

	.mt-20 {
		margin-top: 5rem
	}

	.mt-2 {
		margin-top: 0.5rem
	}

	.mt-1 {
		margin-top: 0.25rem
	}

	.mr-4 {
		margin-right: 1rem
	}

	.block {
		display: block
	}

	.inline {
		display: inline
	}

	.flex {
		display: flex
	}

	.grid {
		display: grid
	}

	.h-full {
		height: 100%
	}

	.w-full {
		width: 100%
	}

	.max-w-5xl {
		max-width: 64rem
	}

	.flex-1 {
		flex: 1 1 0%
	}

	.list-outside {
		list-style-position: outside
	}

	.list-disc {
		list-style-type: disc
	}

	.grid-cols-3 {
		grid-template-columns: repeat(3, minmax(0, 1fr))
	}

	.grid-cols-1 {
		grid-template-columns: repeat(1, minmax(0, 1fr))
	}

	.flex-row {
		flex-direction: row
	}

	.flex-col {
		flex-direction: column
	}

	.items-center {
		align-items: center
	}

	.justify-center {
		justify-content: center
	}

	.gap-2 {
		gap: 0.5rem
	}

	.overflow-hidden {
		overflow: hidden
	}

	.rounded-sm {
		border-radius: 0.125rem
	}

	.rounded-md {
		border-radius: 0.375rem
	}

	.border {
		border-width: 1px
	}

	.border-b {
		border-bottom-width: 1px
	}

	.border-black {
		--tw-border-opacity: 1;
		border-color: rgb(0 0 0 / var(--tw-border-opacity))
	}

	.bg-white {
		--tw-bg-opacity: 1;
		background-color: rgb(255 255 255 / var(--tw-bg-opacity))
	}

	.p-3 {
		padding: 0.75rem
	}

	.p-2 {
		padding: 0.5rem
	}

	.p-4 {
		padding: 1rem
	}

	.py-1 {
		padding-top: 0.25rem;
		padding-bottom: 0.25rem
	}

	.px-4 {
		padding-left: 1rem;
		padding-right: 1rem
	}

	.pt-10 {
		padding-top: 2.5rem
	}

	.pb-4 {
		padding-bottom: 1rem
	}

	.text-base {
		font-size: 1rem;
		line-height: 1.5rem
	}

	.text-lg {
		font-size: 1.125rem;
		line-height: 1.75rem
	}

	.text-2xl {
		font-size: 1.5rem;
		line-height: 2rem
	}

	.text-5xl {
		font-size: 3rem;
		line-height: 1
	}

	.text-xl {
		font-size: 1.25rem;
		line-height: 1.75rem
	}

	.text-sm {
		font-size: 0.875rem;
		line-height: 1.25rem
	}

	.font-bold {
		font-weight: 700
	}

	.font-semibold {
		font-weight: 600
	}

	.leading-6 {
		line-height: 1.5rem
	}

	.text-white {
		--tw-text-opacity: 1;
		color: rgb(255 255 255 / var(--tw-text-opacity))
	}

	.text-black {
		--tw-text-opacity: 1;
		color: rgb(0 0 0 / var(--tw-text-opacity))
	}

	.underline {
		-webkit-text-decoration-line: underline;
						text-decoration-line: underline
	}

	@media (min-width: 640px) {
		.sm\:mt-0 {
			margin-top: 0px
		}

		.sm\:ml-10 {
			margin-left: 2.5rem
		}

		.sm\:w-5\/12 {
			width: 41.666667%
		}

		.sm\:w-4\/12 {
			width: 33.333333%
		}

		.sm\:flex-row {
			flex-direction: row
		}

		.sm\:p-4 {
			padding: 1rem
		}

		.sm\:p-8 {
			padding: 2rem
		}

		.sm\:text-lg {
			font-size: 1.125rem;
			line-height: 1.75rem
		}
	}
`)

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
					<link rel="stylesheet" href="/assets/css/normalize@8.0.1.css" />
					<link rel="stylesheet" href="/styles.css" />
					<script defer src="/assets/js/alpinejs@3.9.6.js"></script>
			</head>
			<body>
			{{ props.Children }}
			</body>
		</html>
	`)
}

//     <title>${item.title}</title>
//     <meta name="title" content="${item.title}" />
//     <meta name="description" content="${config.description}" />
//     <meta name="image" content="${config.image}" />
//     <meta name="keywords" content="${config.keywords}" />
//     <meta name="author" content="${config.author}" />
//     <meta property="og:type" content="website" />
//     <meta property="og:url" content="${config.url}" />
//     <meta property="og:site_name" content="${config.siteName}" />
//     <meta property="og:title" content="${item.title}" />
//     <meta property="og:description" content="${config.description}" />
//     <meta property="og:image" content="${config.image}" />
//     <link rel="canonical" href="${config.url}" />
//     <link rel="stylesheet" href="/assets/css/dracula.css" />
