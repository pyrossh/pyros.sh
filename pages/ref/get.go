package ref

import (
	"context"

	"github.com/pyros2097/gromer"
	. "github.com/pyros2097/gromer/handlebars"

	"pyros.sh/utils"
)

var _ = Css(`
	.ref-container {
		font-family: system-ui;
	}

	.ref-container li {
		list-style-type: disc;
		list-style-position: outside;
	}

	.ref-container h2 {
		font-weight: 500;
		font-size: 1.4rem;
		line-height: 1.6rem;
		margin: 12px 0;
	}

	.ref-container h3 {
		font-weight: 500;
		font-size: 1.2rem;
    line-height: 1.4rem;
		margin: 0;
		margin: 4px 0;
	}

	.ref-container h4 {
		font-weight: 500;
		font-size: 1rem;
    line-height: 1.2rem;
		margin: 0;
		margin: 8px 0;
	}
`)

func GET(c context.Context) (HtmlContent, int, error) {
	url := gromer.GetUrl(c)
	isPrint := url.Query().Get("print") == "true"
	return Html(`
		{{#Page url=url title="Reference" description="Reference" keywords="pyros.sh,pyrossh,reference"}}
			{{#if isPrint}}
				{{#RefHeader}}{{/RefHeader}}
			{{else}}
				{{#Header}}{{/Header}}
			{{/if}}
			{{#Layout}}
				<div class="ref-container">
					<div>
						Programming is my passion. I've been coding from a very young age and still do so. I have created and contributed to many open source projects. My
						main interest lies in frontend, backend, and devops. I build frameworks and automate common tasks to make it easier to develop applications. I do
						occasionally work on creating 2d and 3d games for fun.
					</div>
					<div class="flex flex-1 flex-col">
						<h2>Experience</h2>
						<div class="mb-4">
							<h3>Equal Experts</h3>
							<h4>Software Developer, Oct 2018 - present</h4>
							<p class="text-black mt-2">
								Making Software. Better. Equal Experts is a network of talented, experienced software consultants, specialising in agile delivery.
							</p>
							<div>
								These are the client projects I worked on,
								<ul class="mt-2 ml-6 sm:ml-10">
									<li>
										John Lewis
										<ul class="mt-2 ml-6 sm:ml-10">
											<li>Built the new Product Description Page (PDP) for Simple Carpets using nextjs in home interiors</li>
											<li>Helped build, improve, and run AB experiments in the customer services and my orders pages in customer hub</li>
										</ul>
									</li>
									<li>
										iOWNA
										<ul class="mt-2 ml-6 sm:ml-10">
											<li>Built a content editor for curators to create content on our platform</li>
											<li>Built the iOWNA crossplatform app in react-native that works in android, ios, web. One codebase</li>
											<li>Implemented admin specific functionalities in the web app</li>
											<li>Implemented an in-app OTA update method using expo and semantic versionin</li>
											<li>Built the iOWNA marketing website</li>
											<li>Built the Health Questionnaire clinician/patient feedback form feature with push notifications</li>
										</ul>
									</li>
									<li>
										Zeta
										<ul class="mt-2 ml-6 sm:ml-10">
											<li>Added multiple download formats (pdf,html,xlsx) for reports in the Reporting Center</li>
											<li>Created a Groovy Script to automate creating ACL's in sandbox authorization framework</li>
											<li>Integrated Camunda Workflow Engine into Zeta services in Operations Center</li>
											<li>Built a few workflows on Camunda and integrated it with some internal services</li>
											<li>Built a custom k8s CRD using the operator framework to automate creating ACL's in sandbox using custom resources</li>
										</ul>
									</li>
									<li>
										Lifebox
										<ul class="mt-2 ml-6 sm:ml-10">
											<li>Added search functionality to the hospital procedure codes</li>
											<li>Converted the Health Questionnaire from json to native go code making it easier to make changes to it</li>
											<li>Integrated with SAP using HL7 message format to create patients, episodes, schedule/cancel procedures</li>
											<li>Built the Patient Document repository for uploading, downloading patient files</li>
											<li>Built the patients online reports view using the data from SAP</li>
										</ul>
									</li>
								</ul>
							</div>
						</div>
						<div class="mb-4">
							<h3>Numberz</h3>
							<h4>Full Stack Developer, Sept 2016 - Oct 2018</h4>
							<p class="text-black mt-2">
								Numberz integrates banking with your day-2-day business work-flows freeing up a lot of effort, time and heart-burn, to help your business
								grow! You???ll never have to spend time syncing information between different systems. Save and record receipts &amp; spends with a single tap,
								automatically categorise transactions for book-keeping and easily share information with your accountant.
							</p>
							<ul class="mt-2 ml-6 sm:ml-10">
								<li>Implemented new features and fixed bugz on the numberz android app</li>
								<li>Integrated numberz with Tally ERP</li>
								<li>Implemented CI and Devops within out system to ease our deployment process using docker and docker-compose</li>
								<li>Created a slack bot called deploybot which was used to deploy our microservices to our servers</li>
								<li>Started tech talks on Thursdays to discuss cutting edge technologies</li>
								<li>Architected and built the Notification Service using rabbitmq and nodejs</li>
								<li>Built the Multi-User and Multi-Company Features</li>
								<li>Implemented and lead a team on integration with the GSTN System so that customers could file their GST Taxes within our product</li>
								<li>Architected and built the numberz Identity Management Service(IMS)</li>
								<li>Architected our new product called the Numberz Accounts Receivable (AR)</li>
								<li>Laid down the architecture for the frontend and backend</li>
								<li>
									Built these core features for our new AR Product - Auto Reminders, Manual Reminders, User Management, Cash Discounting, CFO Dashboard
									Metrics
								</li>
							</ul>
						</div>
						<div class="mb-4">
							<h3>Playlyfe</h3>
							<h4>Full Stack Developer, Apr 2014 - Sept 2016</h4>
							<p class="text-black mt-2">
								Playlyfe is an online Gamification Platform which empowers anyone to design and implement a gamified system.This platform allows a normal user
								to convert a gamified system's design into a fully functional web-application complete with teams, real-time notifications, leaderboards and
								many more features. Made Critical decisions on what tech stack to switch to for our gamification product Catalyst.
							</p>
							<ul class="mt-2 ml-6 sm:ml-10">
								<li>Built the Playlyfe Hybrid Mobile app using Phonegap for Android and iOS</li>
								<li>Built the leaderboard system using Mongodb, Redis and Couchbase as the datastore</li>
								<li>Migration of data from v1 to v2 of the API</li>
								<li>Built SDKs for the Playlyfe v2 API in C#, Java, Python, Ruby, PHP, Nodejs, Gos</li>
								<li>Maintained the job server and wrote most of the background jobs</li>
								<li>Added documentation to most of the REST API and created the Developer Console which was Like Swagger UI</li>
								<li>Integrated our REST API in Leanosphere's LMS</li>
								<li>Integrated the REST API in Knolskape's product Aktivlearn LMS</li>
								<li>Integrated the REST API in Linkstreet's LMS</li>
								<li>Integrated the REST API in Moodle LMS by creating an admin plugin to add Gamification features</li>
								<li>Rebuilt the core platform in golang for performance and developement speed</li>
								<li>Built the GraphQL API for our cutting edge product Catalyst</li>
								<li>Convinced the boss to create our own scripting language plscript for our rule engine.</li>
								<li>Build a REST API for our product catalyst for a POC with Vodafone</li>
								<li>Built the Quiz, Review, Match, Tournament features in Catalyst for a major project with HP</li>
								<li>Added gamification features to Saudi Matches a popular game tracking app</li>
							</ul>
						</div>
					</div>
				</div>
			{{/Layout}}
		{{/Page}}
		`).
		Prop("url", utils.GetUrl(c)).
		Prop("isPrint", isPrint).
		Render()
}
