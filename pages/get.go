package pages

import (
	"context"

	. "github.com/pyros2097/gromer/handlebars"
	"pyros.sh/components"
)

var slides = []*components.Slide{
	{
		Title:       "Website",
		SubTitle:    "",
		Description: "It is hosted on BunnyCDN. It has a perfect score on the lighthouse test.",
		ImgSrc:      "/assets/images/lighthouse.png",
		ImgWidth:    "100%",
	},
	{
		Title:       "rust-embed",
		SubTitle:    "A rust macro which loads files into the rust binary at compile time during release and loads the file from the fs during dev.",
		Description: "You can use this to embed your css, js and images into a single executable which can be deployed to your servers. Also it makes it easy to build a very small docker image for you to deploy.",
		ImgSrc:      "/assets/images/rust-embed.png",
		ImgWidth:    "100%",
	},
	{
		Title:       "wapp",
		SubTitle:    "wapp is a framework to build isomorphic web apps in golang.",
		Description: "It uses a declarative syntax using funcs that allows creating and dealing with HTML elements only by using Go, and without writing any HTML markup. The syntax is inspired by react and its awesome hooks and functional component features. It is highly opioninated and integrates very well with tailwind css for now.",
		ImgSrc:      "/assets/images/rust-embed.png",
		ImgWidth:    "100%",
	},
	{
		Title:       "Pine",
		SubTitle:    "A programming language with a syntax largely inspired by nim but with the simplicity of go. It has jsx support in built by default.",
		Description: "first class functions clean syntax basics: int, float, bool, byte, enum references: string, array, map, nil, error functions: extern, proc, method, test conditions: if, elif, else, match, break, continue loops: for inbuilt: echo, assert",
		ImgSrc:      "/assets/images/pine.png",
		ImgWidth:    "40%",
	},
	{
		Title:       "Gdx Studio",
		SubTitle:    "GdxStudio is used for creating awesome games using libGdx.",
		Description: "It has all the features of libgdx built-in so you can easily,start creating games with it. Automatic Asset Loading including Atlas, TextureRegions, BitmapFonts, Music, Sound. Tools like Font Editor, Particle Editor, Texture Packer, SceneEditor, MapEditor, ActorEditor, ImagingTools are already built into it.",
		ImgSrc:      "/assets/images/gdx-studio.png",
		ImgWidth:    "100%",
	},
}

func GET(ctx context.Context) (HtmlContent, int, error) {
	return Html(`
		{{#Page title="pyros.sh"}}
			{{#Header}}{{/Header}}
			<main class="w-full h-full">
				<div class="w-full flex flex-1 flex-row justify-center">
					<div class="flex flex-row flex-1 items-center max-w-5xl text-lg font-source p-2 mt-4 sm:p-4">
						<div>
							<div>
								Hi there, I'm <strong>Peter John</strong>, a fullstack developer from Bangalore, India. I love writing code and I am lucky enough to do this as my
								job. I currently work for Equal Experts. I have a strong passion for golang but I also work with react and nodejs.
								<p>
									I like to work on open source and hobby projects. Over the course of 8 years I've accumlated a lot of useful projects used by many people around
									the world.
								</p>
							</div>
							{{#Slider slides=slides}}
							{{/Slider}}
							<div class="flex flex-col sm:flex-row">
								<div class="flex flex-1 flex-col">
									<div>
										<div class="text-2xl font-bold mt-6"><i class="icon-sitemap"></i>Interests</div>
										<div class="mt-6">These are some of the stuff I work on</div>
										<ul class="ml-6 list-disc grid grid-cols-3 gap-2 mt-6">
											<li>HTML</li>
											<li>Javascript</li>
											<li>CSS</li>
											<li>SVG</li>
											<li>Go</li>
											<li>Rust</li>
											<li>Nodejs</li>
											<li>Python</li>
											<li>Java</li>
											<li>Reactjs</li>
											<li>Serverless</li>
											<li>Web Components</li>
										</ul>
									</div>
								</div>
								<div class="flex flex-1 flex-col">
									<div>
										<div class="text-2xl font-bold mt-6"><i class="icon-chat"></i> Contact</div>
										<div class="mt-6">You can contact me through any of these methods</div>
										<ul class="mt-6 ml-6 list-disc text-lg grid grid-cols-1 gap-2">
											<li>Email: <a class="link ml-2" href="mailto:peter@pyros.sh">peter@pyros.sh</a></li>
											<li>Github: <a class="link ml-2" href="https://github.com/pyros2097">pyros2097</a></li>
											<li>LinkedIn: <a class="link ml-2" href="https://www.linkedin.com/in/peter-john-in">Peter John</a></li>
										</ul>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</main>
		{{/Page}}
		`).
		Prop("slides", slides).
		Render()
}
