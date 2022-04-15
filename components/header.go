package components

import (
	. "github.com/pyros2097/gromer/handlebars"
)

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
