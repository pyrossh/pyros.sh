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

	header > nav > div {
		display: flex;
		flex-direction: row;
		flex: 1;
		align-items: center;
		padding: var(--space-3);
		max-width: var(--5xl);
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

	header > nav .links-container {
		display: flex;
		flex-direction: row;
		flex: 1;
		align-items: center;
		font-size: 1rem;
		line-height: 1.5rem
	}

	@media (min-width: 640px) {
		header > nav .links-container {
			font-size: 1.125rem;
			line-height: 1.75rem
		}
	}
	
	header > nav .nav-link {
		color: white;
		margin-left: 0.25rem;
		margin-right: 0.25rem;
		padding-left: 0.5rem;
		padding-right: 0.5rem;
	}

	header > nav .nav-sep {
		color: white;
		font-size: 1.125rem;
		line-height: 1.75rem
	}
	
	header > nav .nav-link:hover {
		background-color: rgb(55 65 81);
	}
`)

func Header() *Template {
	return Html(`
		<header>
      <nav>
          <div>
            <a class="logo" href="/"> pyros.sh </a>
            <div class="links-container">
              <a class="nav-link" href="/work"> work </a>
              <div class="nav-sep">|</div>
              <a class="nav-link" href="/ref"> ref </a>
              <div class="nav-sep">|</div>
              <a class="nav-link" href="/blog"> blog </a>
          </div>
      </nav>
    </header>
	`)
}
