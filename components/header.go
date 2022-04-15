package components

import (
	. "github.com/pyros2097/gromer/handlebars"
)

var _ = Css(`
	header > nav {
		display: flex;
		width: 100%;
		flex: 1 1 0%;
		flex-direction: row;
		justify-content: center;
		background-color: var(--black);
		font-family: monospace;
	}
	
	header > nav .logo {
		display: flex;
		padding-left: 0.25rem;
		padding-right: 1rem;
		font-size: 1.25rem;
		line-height: 1.75rem;
		color: var(--yellow);
	}
	
	@media (min-width: 640px) {
		header > nav .logo {
			font-size: 1.5rem;
			line-height: 2rem;
		}
	}
	
	header > nav .nav-link {
		margin-left: 0.25rem;
		margin-right: 0.25rem;
		padding-left: 0.5rem;
		padding-right: 0.5rem;
	}
	
	header > nav .nav-link:hover {
		background-color: rgb(55 65 81);
	}
`)

func Header() *Template {
	return Html(`
		<header>
      <nav>
          <div class="flex flex-row flex-1 items-center max-w-5xl text-white p-3">
            <a class="logo" href="/"> pyros.sh </a>
            <div class="flex flex-row flex-1 items-center text-base sm:text-lg">
              <a class="nav-link" href="/work"> work </a>
              <div class="text-lg">|</div>
              <a class="nav-link" href="/ref"> ref </a>
              <div class="text-lg">|</div>
              <a class="nav-link" href="/blog"> blog </a>
          </div>
      </nav>
    </header>
	`)
}
