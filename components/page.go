package components

import (
	"html/template"

	. "github.com/pyros2097/gromer/handlebars"
)

var _ = Css(`
	html {
		line-height: 1.15; /* 1 */
		-webkit-text-size-adjust: 100%; /* 2 */
	}

	body {
		margin: 0;
	}

	main {
		display: block;
	}

	h1 {
		font-size: 2em;
		margin: 0.67em 0;
	}

	hr {
		box-sizing: content-box; /* 1 */
		height: 0; /* 1 */
		overflow: visible; /* 2 */
	}

	pre {
		font-family: monospace, monospace; /* 1 */
		font-size: 1em; /* 2 */
	}

	a {
		background-color: transparent;
	}

	abbr[title] {
		border-bottom: none; /* 1 */
		text-decoration: underline; /* 2 */
		text-decoration: underline dotted; /* 2 */
	}

	b,
	strong {
		font-weight: bolder;
	}

	code,
	kbd,
	samp {
		font-family: monospace, monospace; /* 1 */
		font-size: 1em; /* 2 */
	}

	small {
		font-size: 80%;
	}

	sub,
	sup {
		font-size: 75%;
		line-height: 0;
		position: relative;
		vertical-align: baseline;
	}

	sub {
		bottom: -0.25em;
	}

	sup {
		top: -0.5em;
	}

	img {
		border-style: none;
	}

	button,
	input,
	optgroup,
	select,
	textarea {
		font-family: inherit; /* 1 */
		font-size: 100%; /* 1 */
		line-height: 1.15; /* 1 */
		margin: 0; /* 2 */
	}

	button,
	input { /* 1 */
		overflow: visible;
	}

	button,
	select { /* 1 */
		text-transform: none;
	}

	button,
	[type="button"],
	[type="reset"],
	[type="submit"] {
		-webkit-appearance: button;
	}

	button::-moz-focus-inner,
	[type="button"]::-moz-focus-inner,
	[type="reset"]::-moz-focus-inner,
	[type="submit"]::-moz-focus-inner {
		border-style: none;
		padding: 0;
	}

	button:-moz-focusring,
	[type="button"]:-moz-focusring,
	[type="reset"]:-moz-focusring,
	[type="submit"]:-moz-focusring {
		outline: 1px dotted ButtonText;
	}

	fieldset {
		padding: 0.35em 0.75em 0.625em;
	}

	legend {
		box-sizing: border-box; /* 1 */
		color: inherit; /* 2 */
		display: table; /* 1 */
		max-width: 100%; /* 1 */
		padding: 0; /* 3 */
		white-space: normal; /* 1 */
	}

	progress {
		vertical-align: baseline;
	}

	textarea {
		overflow: auto;
	}

	[type="checkbox"],
	[type="radio"] {
		box-sizing: border-box; /* 1 */
		padding: 0; /* 2 */
	}

	[type="number"]::-webkit-inner-spin-button,
	[type="number"]::-webkit-outer-spin-button {
		height: auto;
	}

	[type="search"] {
		-webkit-appearance: textfield; /* 1 */
		outline-offset: -2px; /* 2 */
	}

	[type="search"]::-webkit-search-decoration {
		-webkit-appearance: none;
	}

	::-webkit-file-upload-button {
		-webkit-appearance: button; /* 1 */
		font: inherit; /* 2 */
	}

	details {
		display: block;
	}

	summary {
		display: list-item;
	}

	template {
		display: none;
	}

	[hidden] {
		display: none;
	}

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
		--black-light: #313a3d;
		--yellow: #f1fa8c;
		--yellow-dark: #af9301;
		--blue: #0645ad;
		--blue-light: #0e81c7;
		--slider-bg: #f0ede2;
		--space-2: 0.5rem;
		--space-3: 0.75rem;
		--space-4: 1rem;
		--space-10: 2.5rem;
		--f-normal: 1rem;
		--f-large: 1.125rem;
		--f-xlarge: 1.25rem;
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

	a {
		font-size: 1.125rem;
		line-height: 1.75rem;
		color: var(--blue);
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

	.rounded-sm {
		border-radius: 0.125rem
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

	.p-2 {
		padding: 0.5rem
	}

	.p-4 {
		padding: 1rem
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
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Keywords    string        `json:"keywords"`
	Url         string        `json:"url"`
	Image       string        `json:"image" default:"/assets/images/icon.png"`
	Children    template.HTML `json:"children"`
}

func Page(props PageProps) *Template {
	return Html(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
					<meta charset="UTF-8" />
					<meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
					<meta content="utf-8" http-equiv="encoding" />
					<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=5, viewport-fit=cover" />
					<title>{{ props.Title }}</title>
					<meta name="description" content="{{ props.Description }}" />
					<meta name="author" content="pyrossh" />
					<meta name="image" content="{{ props.Image }}" />
					<meta name="keywords" content="{{ props.Keywords }}" />
					<meta property="og:type" content="website" />
					<meta property="og:url" content="{{ props.Url }}" />
					<meta property="og:site_name" content="pyros.sh" />
					<meta property="og:title" content="{{ props.Title }}" />
					<meta property="og:description" content="{{ props.Description }}" />
					<meta property="og:image" content="{{ props.Image }}" />
					<link rel="canonical" href="{{ props.Url }}" />
					<link rel="icon" href="/assets/images/icon.png" />
					<link rel="stylesheet" href="/styles.css" />
			</head>
			<body>
			{{ props.Children }}
			</body>
		</html>
	`)
}
