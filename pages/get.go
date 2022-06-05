package pages

import (
	"context"

	"github.com/pyros2097/gromer"
	. "github.com/pyros2097/gromer/handlebars"
	"pyros.sh/assets"
	"pyros.sh/components"
	"pyros.sh/utils"
)

var slides = []*components.Slide{
	{
		Title:       "Rust-embed",
		Link:        "https://github.com/pyros2097/rust-embed",
		SubTitle:    "A rust macro which loads files into the rust binary at compile time during release and loads the file from the fs during dev.",
		Description: "You can use this to embed your css, js and images into a single executable which can be deployed to your servers. Also it makes it easy to build a very small docker image for you to deploy.",
		ImgSrc:      gromer.GetAssetUrl(assets.FS, "images/rust-embed.png"),
	},
	{
		Title:       "Gromer",
		Link:        "https://github.com/pyros2097/gromer",
		SubTitle:    "Gromer is a framework and cli to build web apps in golang.\nIt uses a declarative syntax using inline handlebar templates for components and pages.",
		Description: "It also generates http handlers for your routes which follow a particular folder structure. Similar to other frameworks like nextjs, sveltekit. These handlers are also normal functions and can be imported in other packages to call them directly.",
		ImgSrc:      gromer.GetAssetUrl(assets.FS, "images/gromer.png"),
	},
	{
		Title:       "Pine",
		Link:        "https://github.com/pyros2097/pine",
		SubTitle:    "A programming language with a syntax largely inspired by pyret and hush but with the simplicity of go.",
		Description: "It has support first class functions and has a clean syntax and can call C library functions directly. Pine programs can be compiled to AMD64 and ARM64 and this is done using QBE a small and fast compiler backend.",
		ImgSrc:      gromer.GetAssetUrl(assets.FS, "images/pine.png"),
	},
	{
		Title:       "Gdx Studio",
		Link:        "https://github.com/pyros2097/gdx-studio",
		SubTitle:    "GdxStudio is used for creating awesome games using libGDX.",
		Description: "It has all the features of libGDX built-in so you can easily,start creating games with it. Automatic Asset Loading including Atlas, TextureRegions, BitmapFonts, Music, Sound. Tools like Font Editor, Particle Editor, Texture Packer, SceneEditor, MapEditor, ActorEditor, ImagingTools are already built into it.",
		ImgSrc:      gromer.GetAssetUrl(assets.FS, "images/gdx-studio.png"),
	},
}

var _ = Css(`
	.home-container li {
		list-style-type: disc;
	}

	.icon-chat:before {
		content: '\e800';
	}

	.icon-sitemap:before {
    content: '\f0e8';
	}

		[class^="icon-"]:before, [class*=" icon-"]:before {
			font-family: "fa-elv-sh";
			font-style: normal;
			font-weight: normal;
			speak: never;
			display: inline-block;
			text-decoration: inherit;
			width: 1em;
			margin-right: 0.2em;
			text-align: center;
			/* opacity: .8; */
			font-variant: normal;
			text-transform: none;
			line-height: 1em;
			margin-left: 0.2em;
			/* font-size: 120%; */
			-webkit-font-smoothing: antialiased;
			-moz-osx-font-smoothing: grayscale;
			/* text-shadow: 1px 1px 1px rgb(127 127 127 / 30%); */
	}
`)

func GET(c context.Context) (HtmlContent, int, error) {
	return Html(`
		{{#Page url=url title="pyros.sh" description="Hi there, I'm Peter John, a fullstack developer from Bangalore, India." keywords="peter john,pyros.sh,pyrossh,full stack developer,bangalore,india"}}
			{{#Header}}{{/Header}}
			{{#Layout}}
				<div class="home-container">
					<div>
						Hi there, I'm <strong>Peter John</strong>, a fullstack developer from Bangalore, India. I love writing code and I am lucky enough to do this as my
						job. I currently work for Equal Experts. I have a strong passion for golang but I also work with react and nodejs.
						<p>
							I like to work on open source and hobby projects. Over the course of 8 years I've accumlated a lot of useful projects used by many people around
							the world. These are some of them,
						</p>
					</div>
					{{#Slider slides=slides}}
					{{/Slider}}
					<div class="flex flex-col sm:flex-row">
						<div class="flex flex-1 flex-col">
							<div>
								<div class="text-2xl font-bold mt-6"><i class="icon-sitemap"></i>Interests</div>
								<div class="mt-6">These are some of the stuff I work on</div>
								<ul class="ml-6 grid grid-cols-3 gap-2 mt-6">
									<li>HTML</li>
									<li>Javascript</li>
									<li>CSS</li>
									<li>SVG</li>
									<li>Go</li>
									<li>Rust</li>
									<li>Node</li>
									<li>Java</li>
									<li>React</li>
									<li>Postgres</li>
									<li>k8s</li>
									<li>Serverless</li>
								</ul>
							</div>
						</div>
						<div class="flex flex-1 flex-col">
							<div>
								<div class="text-2xl font-bold mt-6"><i class="icon-chat"></i> Contact</div>
								<div class="mt-6">You can contact me through any of these methods</div>
								<ul class="mt-6 ml-6 text-lg grid grid-cols-1 gap-2">
									<li>Email: <a class="ml-2" href="mailto:peter.john@pyros.sh">peter.john@pyros.sh</a></li>
									<li>Github: <a class="ml-2" href="https://github.com/pyrossh">pyrossh</a></li>
									<li>LinkedIn: <a class="ml-2" href="https://www.linkedin.com/in/peter-john-in">Peter John</a></li>
								</ul>
							</div>
						</div>
					</div>
				</div>
			{{/Layout}}
		{{/Page}}
		`).
		Prop("slides", slides).
		Prop("url", utils.GetUrl(c)).
		Render()
}
