package components

import (
	. "github.com/pyros2097/gromer/handlebars"
)

var _ = Css(`
	.ref-header {
		display: flex;
		flex-direction: column;
		background-color: var(--black);
		font-family: system-ui;
		padding-left: 8px;
	}

	.ref-header .row {
		display: flex;
		flex-direction: row;
		flex: 1;
	}
	
	.ref-header .name {
		font-size: 1.5rem;
		line-height: 2rem;
		color: var(--yellow);
	}

	.ref-header .role {
		font-size: 1.1rem;
		line-height: 1.3rem;
		color: gray;
		margin-top: 4px;
	}

	.ref-header .location {
		font-size: 1rem;
		line-height: 1.2rem;
		color: gray;
		margin-top: 4px;
	}
	
	.ref-header .contact-link {
		display: flex;
		flex: 1;
		justify-content: flex-end;
		align-items: center;
		color: black;
		font-size: 1rem;
		line-height: 1.2rem;
	}
`)

func RefHeader() *Template {
	return Html(`
		<div class="ref-header">
			<div class="row">
				<span class="name"> Peter John </span>
				<a class="contact-link" href="/work">peter.john@pyros.sh </a>
			</div>
			<div class="row">
				<span class="role" href="/"> Software Developer </span>
				<a class="contact-link" href="/ref">+919036542841 </a>
			</div>
			<div class="row">
				<span class="location" href="/"> Bengaluru, India </span>
			</div>
    </div>
	`)
}
