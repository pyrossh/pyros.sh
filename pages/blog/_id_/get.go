package blog_id_

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"strings"

	. "github.com/pyros2097/gromer/handlebars"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"pyros.sh/assets"
)

var markdown = goldmark.New(
	goldmark.WithExtensions(
		highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"),
		),
	),
)

func GET(c context.Context, id string) (HtmlContent, int, error) {
	files, err := assets.FS.ReadDir("md")
	if err != nil {
		return HtmlErr(500, err)
	}
	for _, f := range files {
		title := f.Name()[11 : len(f.Name())-3]
		postId := strings.ReplaceAll(title, " ", "-")
		if postId == id {
			source, err := assets.FS.ReadFile("md/" + f.Name())
			if err != nil {
				return HtmlErr(500, err)
			}
			var buf bytes.Buffer
			if err := markdown.Convert(source, &buf); err != nil {
				return HtmlErr(500, err)
			}
			return Html(`
				{{#Page title="Blog"}}
					{{#Header}}{{/Header}}
					{{#Layout}}
						<div class="block">
							{{ md }}
						</div>
					{{/Layout}}	
					</main>
				{{/Page}}
			`).
				Prop("md", template.HTML(buf.String())).
				Render()
		}
	}
	return HtmlErr(404, fmt.Errorf("Post not found"))
}

// export const head = ({ config, item }) => {
//   return html`
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
//   `;
// };

// export const body = ({ item }) => {
//   return html`
//     <app-header></app-header>
//     <main class="w-full h-full">
//       <div class="w-full flex flex-1 flex-row justify-center">
//         <div class="flex flex-row flex-1 items-center max-w-5xl text-lg font-source p-4 mt-4">
//           <div class="w-full flex flex-1 flex-row justify-center">
//             <div class="flex flex-row flex-1 items-center max-w-5xl text-lg font-source p-4 mt-4">
//               <div>
//                 <div class="flex pb-4 border-black border-b">
//                   <div class="flex-1 text-2xl font-bold">${item.title}</div>
//                   <div class="mr-4">${item.uploadedOn}</div>
//                 </div>
//                 ${item.description.map((text) => {
//                   if (text) {
//                     if (text.code) {
//                       // bg-codebg font-monospace text-sm rounded-md py-1 px-4 my-3 mr-4
//                     }
//                     if (text.img) {
//                       return html`<div class="my-6"><img src="${text.img}" alt="${text.img}" /></div>`;
//                     }
//                     return html` <p class="text-black font-source mt-2">${text}</p> `;
//                   }
//                 })}
//               </div>
//             </div>
//           </div>
//         </div>
//       </div>
//     </main>
//   `;
// };
