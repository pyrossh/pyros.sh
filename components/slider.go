package components

import (
	. "github.com/pyros2097/gromer/handlebars"
)

type Slide struct {
	Title       string `json:"title"`
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
        <ul class="flex flex-row mb-4">
					{{#each props.Slides as |slide| }}
						<li class="px-3 py-1 cursor-pointer select-none bg-black text-white hover:bg-gray-300">
                {{@index}}
            </li>
					{{/each}}
        </ul>
        <div class="bg-slider mb-8 rounded-sm overflow-hidden">
					{{#each props.Slides as |slide| }}
						<div class="flex flex-row relative w-full overflow-hidden">
							<div class="flex flex-col sm:flex-row w-full p-4 sm:p-8">
								<div class="sm:w-5/12 flex flex-1 flex-col mr-8">
									<p class="text-2xl font-bold mb-4">{{ slide.Title }}</p>
									<p class="leading-6 mb-4">{{ slide.SubTitle }}</p>
									<p class="leading-6">{{ slide.Description }}</p>
								</div>
								<div class="flex-1 mt-8 sm:mt-0 sm:w-4/12 bg-white border-black border rounded-sm flex flex-col items-center justify-center">
									<div class="p-4">
										<img src="{{ slide.ImgSrc }}" alt={{ slide.Title }} width={{ slide.ImgWidth }} />
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

// ${title === 'Website'
//             ? html`<div class="flex flex-row relative w-full overflow-hidden">
//                 <div class="flex flex-col sm:flex-row w-full p-4 sm:p-8">
//                   <div class="sm:w-5/12 flex flex-1 flex-col mr-8">
//                     <p class="text-2xl font-bold mb-4">${title}</p>
//                     <p class="leading-6 mb-4">
//                       This website is built with
//                       <a class="link" href="https://github.com/pyros2097/density-ssg" target="_blank" rel="noopener noreferrer">density-ssg</a>, a web
//                       components based opinionated Static Site Generator (SSG).
//                     </p>
//                     <p class="leading-6">${description}</p>
//                   </div>
//                   <div class="flex-1 mt-8 sm:mt-0 sm:w-4/12 bg-white border-black border rounded-sm flex flex-col items-center justify-center">
//                     <div class="p-4">
//                       <img src=${imgSrc} alt=${title} width=${imgWidth} />
//                     </div>
//                   </div>
//                 </div>
//               </div>`
//             : html``}
