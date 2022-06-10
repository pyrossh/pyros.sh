package components

import (
	"html/template"

	"github.com/pyros2097/gromer"
	. "github.com/pyros2097/gromer/handlebars"
	"pyros.sh/assets"
)

var _ = Css(`
	@font-face {
  	font-family: 'fa-elv-sh';
  	src: url('data:application/octet-stream;base64,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') format('woff');
	}

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
		--yellow-dark: #6b5a04;
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
		font-size: 1.125rem;
		line-height: 1.75rem;
		font-weight: 400;
		font-family: serif;
		direction: ltr;
		font-synthesis: none;
		text-rendering: optimizeLegibility;
	}

	a {
		font-size: 1.125rem;
		line-height: 1.75rem;
		color: var(--blue);
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

	.w-full {
		width: 100%
	}

	.flex-1 {
		flex: 1 1 0%
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

	@page {
  	size: 8.5in 11in; 
    margin: 0.5in;

				/* size: A4;
		margin: 48px 24px; */
	}

	@media print {
		body {
			font-size: 16px;
			line-height: 24px;
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
					<link rel="icon" href="{{ iconUrl }}" />
					<link rel="stylesheet" href="{{ stylesUrl }}" />
			</head>
			<body>
			{{ props.Children }}
			</body>
		</html>
	`).Props(
		"iconUrl", gromer.GetAssetUrl(assets.FS, "images/icon.png"),
		"stylesUrl", gromer.GetStylesUrl(),
		"picoCssUrl", gromer.GetAssetUrl(assets.FS, "css/pico.min.css"),
	)
}
