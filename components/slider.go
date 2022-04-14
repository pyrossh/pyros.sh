package components

import (
	. "github.com/pyros2097/gromer/handlebars"
)

type Slide struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	SubTitle    string `json:"subTitle"`
	Description string `json:"description"`
	ImgSrc      string `json:"imgSrc"`
	ImgWidth    string `json:"imgWidth"`
}

type SliderProps struct {
	Slides []*Slide `json:"slides"`
}

func Slider(props SliderProps) *Template {
	return Html(`
		<article>
      <div class="pt-10">
        <div class="bg-slider mb-8 rounded-sm overflow-hidden">
					{{#each props.Slides as |slide| }}
						<div class="flex flex-row relative w-full overflow-hidden">
							<div class="flex flex-col sm:flex-row w-full p-4 sm:p-8">
								<div class="sm:w-5/12 flex flex-1 flex-col mr-8">
									<a href="{{ slide.Link }}" class="underline" target="_blank" rel="noopener noreferrer">
										<p class="text-2xl font-bold mb-4">{{ slide.Title }}</p>
									</a>
									<p class="leading-6 mb-4">{{ slide.SubTitle }}</p>
									<p class="leading-6">{{ slide.Description }}</p>
								</div>
								<div class="flex-1 mt-8 sm:mt-0 sm:w-4/12 bg-white border-black border rounded-sm flex flex-col items-center justify-center">
									<div>
										<img src="{{ slide.ImgSrc }}" alt="{{ slide.Title }}" width="100%" />
									</div>
								</div>	
							</div>
						</div>
					{{/each}}
        </div>
      </div>
    </article>
	`)
}
