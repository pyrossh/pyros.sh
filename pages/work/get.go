package work

import (
	"context"

	. "github.com/pyros2097/gromer/handlebars"

	"pyros.sh/utils"
)

func GET(c context.Context) (HtmlContent, int, error) {
	return Html(`
		{{#Page url=url title="Work" description="Work" keywords="pyros.sh,pyrossh,work"}}
			{{#Header}}{{/Header}}
			{{#Layout}}
				<div>
					<div>
						I'm currently a tech lead at Equal Experts. I have around 7 years of work experience. I've mostly worked with startups and product based
						companies. I have gained a lot of domain specific knowledge in healthcare and finance during these startup years.
					</div>
					<div>These are some of the products I've worked on,</div>
					<div class="flex flex-row">
						<div class="flex flex-1 flex-col">
							<div>
								<div class="text-2xl font-bold mt-8">iOWNA</div>
								<div>
									iOWNA is a digital app that provides clinicians with a library of trusted guidance in a patient friendly format that they have at their
									fingertips to distribute to their patients to improve outcomes and enable people live longer healthier lives.
								</div>
							</div>
							<div>
								<div class="text-2xl font-bold mt-8">LifeBox</div>
								<div>
									LifeBox is a smart ePOA and health record system which supports pre-surgical patient assessment, hospital decision-making and treatment.
								</div>
							</div>
							<div>
								<div class="text-2xl font-bold mt-8">Numberz AR</div>
								<div>An accounting app to help companies get their invoices paid faster and reconcile with the bank.</div>
							</div>
							<div>
								<div class="text-2xl font-bold mt-8">Numberz CFM</div>
								<div>An invoicing and payment app to help SME businesses handle their cash flow properly and avail loans.</div>
							</div>
							<div>
								<div class="text-2xl font-bold mt-8">Catalyst</div>
								<div>
									An employee engagement app built using the playlyfe v3 gamification api. Used by the likes of HP, Accenture to improve their call center
									performance metrics.
								</div>
							</div>
							<div>
								<div class="text-2xl font-bold mt-8">Playlyfe</div>
								<div>A gamification as a service platform where developers could gamify their apps and websites easily.</div>
							</div>
						</div>
					</div>
				</div>
			{{/Layout}}
		{{/Page}}
		`).
		Prop("url", utils.GetUrl(c)).
		Render()
}
