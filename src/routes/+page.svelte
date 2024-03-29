<script>
	import Slide from '$lib/components/Slide.svelte';
	import stats from '$lib/assets/logos/stats.png';
	import brave from '$lib/assets/logos/brave.png';
	import code from '$lib/assets/logos/code.png';
	import helix from '$lib/assets/logos/helix.png';
	import fish from '$lib/assets/logos/fish.png';
	import iterm from '$lib/assets/logos/iterm.png';
	import slurp from '$lib/assets/logos/slurp.png';
	import curve from '$lib/assets/logos/curve.png';
</script>

<svelte:head>
	<title>pyros.sh</title>
	<meta
		name="description"
		content="A fullstack developer from Bangalore who likes to create frameworks and programming languages."
	/>
</svelte:head>

<div class="mx-auto">
	<div>
		<h1 class="text-3xl font-bold mb-4">Hello!</h1>
		<p>
			I'm a full-stack engineer from Bengaluru, India named
			<a
				class="underline"
				href="https://www.linkedin.com/in/pyrossh"
				target="_blank"
				rel="noopener noreferrer"
			>
				<strong>Peter John</strong>
			</a>. I have the good fortune to write code for a living.
		</p>
		<p>
			At the moment, Equal Experts employs me. While I've experimented with functional programming
			languages like Haskell, Idris, and Koka, I prefer to use statically typed imperative
			programming languages like Rust, Go, and Java. I've generally worked for new businesses and
			corporations that focus on products. Throughout these early years of launch, I have learned a
			great deal about the healthcare and financial industries.
		</p>
		<p>
			I enjoy working on open-source projects, and during the last decade, I've created a number of
			valuable projects that are extensively utilized by people worldwide.
		</p>
		<div>
			<div class="flex items-center mt-8 mb-4">
				<div class="i-ant-design:project-outlined text-pink-800 mr-2" />
				<h2>Projects</h2>
			</div>
			<div class="grid grid-cols-1 gap-4 sm:grid-cols-2 sm:gap-8">
				<Slide title="rust-embed" link="https://github.com/pyrossh/rust-embed">
					<p>
						A rust proc-macro which loads files into the rust binary at compile time during release
						and loads the file from the fs during dev
					</p>
				</Slide>
				<Slide title="pyros.sh" link="https://github.com/pyrossh/pyros.sh">
					<p>
						木 Personal website of pyrossh. Built with
						<span class="children:underline">
							<a href="https://kit.svelte.dev/" target="_blank" rel="noopener noreferrer"
								>sveltekit</a
							>,
							<a href="https://unocss.dev/" target="_blank" rel="noopener noreferrer">unocss</a>,
							<a href="https://iconify.design/" target="_blank" rel="noopener noreferrer">iconify</a
							>,
							<a href="https://vitejs.dev/" target="_blank" rel="noopener noreferrer">vite</a>.
						</span>
					</p>
				</Slide>
				<Slide title="pacos.sh" link="https://github.com/pyrossh/pacos.sh">
					<p>👾 A statically typed, functional programming language inspired by rust, koka.</p>
					<p>WIP to support compilation to WASM.</p>
				</Slide>
				<Slide title="onlybible.app" link="https://github.com/pyrossh/onlybible.app">
					<p>
						The only bible app you will ever need. No ads, No in-app purchases, No distractions.
					</p>
					<div class="flex mt-4">
						<a
							class="mr-4"
							href="https://play.google.com/store/apps/details?id=sh.pyros.only_bible_app"
							rel="noopener noreferrer"
							target="blank"
						>
							<!-- <GooglePlayIcon /> -->
						</a>
						<a
							href="https://apps.apple.com/us/app/only-bible-app/id6467606465"
							rel="noopener noreferrer"
							target="blank"
						>
							<!-- <AppStoreIcon /> -->
						</a>
					</div>
				</Slide>
				<section>
					<div class="flex items-center">
						<div class="i-eos-icons:hardware-circuit text-green-800 mr-2" />
						<h2>Hardware</h2>
					</div>
					<ul class="grid gap-2 grid-cols-1 text-center mt-4 children:(bg-gray-200 p-1)">
						<li>M2 Macbook Air</li>
						<li>Raspberry Pi 4B</li>
						<li>Raspberry Pi Zero 2W</li>
					</ul>
				</section>
				<section>
					<div class="flex items-center">
						<div class="i-eos-icons:software-outlined text-blue-800 mr-2" />
						<h2>Software</h2>
					</div>
					<ul class="grid gap-2 grid-cols-3 text-center mt-4 children:(bg-gray-200 p-2 pb-4)">
						<li>
							<a
								class="soft-link"
								href="https://github.com/exelban/stats"
								target="_blank"
								rel="noopener noreferrer"
							>
								Stats
								<img class="soft-icon" src={stats} alt="Stats" />
							</a>
						</li>
						<li>
							<a
								class="soft-link"
								href="https://github.com/brave/brave-browser"
								target="_blank"
								rel="noopener noreferrer"
							>
								Brave
								<img class="soft-icon" src={brave} alt="Brave" />
							</a>
						</li>
						<li>
							<a
								class="soft-link"
								href="https://github.com/microsoft/vscode"
								target="_blank"
								rel="noopener noreferrer"
							>
								VS Code
								<img class="soft-icon" src={code} alt="VsCode" />
							</a>
						</li>
						<li>
							<a
								class="soft-link"
								href="https://github.com/helix-editor/helix"
								target="_blank"
								rel="noopener noreferrer"
							>
								Helix
								<img class="soft-icon" src={helix} alt="Helix" />
							</a>
						</li>
						<li>
							<a
								class="soft-link"
								href="https://github.com/fish-shell/fish-shell"
								target="_blank"
								rel="noopener noreferrer"
							>
								Fish
								<img class="soft-icon" src={fish} alt="Fish" />
							</a>
						</li>
						<li>
							<a
								class="soft-link"
								href="https://github.com/gnachman/iTerm2"
								target="_blank"
								rel="noopener noreferrer"
							>
								iTerm2
								<img class="soft-icon" src={iterm} alt="iterm2" />
							</a>
						</li>
						<li>
							<a
								class="soft-link"
								href="https://colorslurp.com"
								target="_blank"
								rel="noopener noreferrer"
							>
								Slurp
								<img class="soft-icon" src={slurp} alt="slurp" />
							</a>
						</li>
						<li>
							<a
								class="soft-link"
								href="https://www.vectornator.io"
								target="_blank"
								rel="noopener noreferrer"
							>
								Curve
								<img class="soft-icon" src={curve} alt="curve" />
							</a>
						</li>
					</ul>
				</section>
				<section>
					<div class="flex items-center">
						<div class="i-mdi:file-tree-outline text-amber-900 mr-2" />
						<h2>Interests</h2>
					</div>
					<ul class="grid gap-2 grid-cols-3 text-center mt-4 children:(bg-slate-100 p-1)">
						<li>HTML</li>
						<li>CSS</li>
						<li>Tailwind</li>
						<li>Javascript</li>
						<li>React</li>
						<li>Flutter</li>
						<li>Go</li>
						<li>Rust</li>
						<li>Kotlin</li>
						<li>Firestore</li>
						<li>Kubernetes</li>
						<li>Istio</li>
					</ul>
				</section>
				<section>
					<div class="flex items-center">
						<div class="i-mdi-contact-outline text-cyan-950 mr-2" />
						<h2>Contact</h2>
					</div>
					<ul
						class=" grid gap-2 grid-cols-1 text-left mt-4 children:(bg-slate-100 p-2 flex flex-col items-baseline sm:flex-row)"
					>
						<li>
							<strong class="mr-2">Email:</strong>
							<span>pyros2097@gmail.com</span>
						</li>
						<li>
							<strong class="mr-2">Github:</strong>
							<a class="text-blue-900" href="https://github.com/pyrossh">
								https://github.com/pyrossh
							</a>
						</li>
						<li>
							<strong class="mr-2">LinkedIn:</strong>
							<a class="text-blue-900" href="https://www.linkedin.com/in/peter-john-in">
								https://www.linkedin.com/in/pyrossh
							</a>
						</li>
					</ul>
				</section>
			</div>
		</div>
	</div>
</div>

<style>
	h2 {
		--at-apply: text-xl font-semibold;
	}

	section {
		--at-apply: flex flex-col mt-4 mx-1;
	}

	.soft-link {
		--at-apply: text-base text-blue-900 flex flex-col items-center;
	}

	.soft-icon {
		--at-apply: w-12 h-12;
	}
</style>
