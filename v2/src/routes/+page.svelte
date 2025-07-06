<script lang="ts">
  const { data } = $props();
  const { recentPosts } = data;
</script>

<div class="mx-auto max-w-lg md:max-w-4xl">
  <div class="mb-20 flex flex-col md:flex-row">
    <img
      class="mb-10 h-32 w-32 self-center rounded-full md:mb-0 md:mr-10 md:h-64 md:w-64"
      src="/images/headshot.webp"
      alt="Headshot of Michael Miranda"
    />
    <div>
      <h1
        class="mb-10 text-center text-4xl font-semibold leading-tight text-slate-900 md:text-left md:text-5xl dark:text-white"
      >
        Hey! I'm Michael <span class="animate-wave inline-block">👋🏼</span>
      </h1>
      <div class="flex flex-col gap-y-4">
        <p class="text-lg md:text-xl">
          I'm a software engineer and here you'll find some of my projects,
          experiences, and maybe even a few hot takes on the tech industry.
        </p>
        <p class="text-lg md:text-xl">
          Find me on Twitter
          <a class="text-blue-500" href="https://twitter.com/mrpineapples24">
            &#64;mrpineapples24
          </a>
        </p>
      </div>
    </div>
  </div>
  <div class="mb-12 flex flex-col">
    <h1
      class="mb-6 w-max border-b-2 border-pink-300 pb-2 text-3xl font-semibold leading-tight text-slate-900 md:text-4xl dark:border-pink-400 dark:text-white"
    >
      Technologies
    </h1>
    <p class="text-lg">
      I enjoy messing around with all kinds of technologies, here are a few of
      my favorites:
    </p>
    <ul
      class="mt-10 grid grid-cols-2 gap-y-5 md:grid-cols-3 md:gap-x-4 md:gap-y-8"
    >
      <li>
        <a
          class="flex items-center gap-4 duration-200 ease-in-out hover:scale-105"
          href="https://developer.mozilla.org/en-US/docs/Web/JavaScript"
        >
          <div
            class="flex h-14 w-14 items-center justify-center rounded-full bg-slate-800 dark:bg-slate-600"
          >
            <img
              class="w-1/2"
              src="/images/logos/javascript.svg"
              alt="JavaScript logo"
            />
          </div>
          JavaScript
        </a>
      </li>
      <li>
        <a
          class="flex items-center gap-4 duration-200 ease-in-out hover:scale-105"
          href="https://nodejs.org"
        >
          <div
            class="flex h-14 w-14 items-center justify-center rounded-full bg-slate-800 dark:bg-slate-600"
          >
            <img
              class="w-1/2"
              src="/images/logos/node.svg"
              alt="Node JS logo"
            />
          </div>
          Node
        </a>
      </li>
      <li>
        <a
          class="flex items-center gap-4 duration-200 ease-in-out hover:scale-105"
          href="https://www.typescriptlang.org"
        >
          <div
            class="flex h-14 w-14 items-center justify-center rounded-full bg-slate-800 dark:bg-slate-600"
          >
            <img
              class="w-1/2"
              src="/images/logos/typescript.svg"
              alt="TypeScript logo"
            />
          </div>
          TypeScript
        </a>
      </li>
      <li>
        <a
          class="flex items-center gap-4 duration-200 ease-in-out hover:scale-105"
          href="https://go.dev"
        >
          <div
            class="flex h-14 w-14 items-center justify-center rounded-full bg-slate-800 dark:bg-slate-600"
          >
            <img
              class="w-1/2"
              src="/images/logos/go.svg"
              alt="Go programming language logo"
            />
          </div>
          Go
        </a>
      </li>
      <li>
        <a
          class="flex items-center gap-4 duration-200 ease-in-out hover:scale-105"
          href="https://react.dev"
        >
          <div
            class="flex h-14 w-14 items-center justify-center rounded-full bg-slate-800 dark:bg-slate-600"
          >
            <img
              class="w-1/2"
              src="/images/logos/react.svg"
              alt="React JS logo"
            />
          </div>
          React
        </a>
      </li>
      <li>
        <a
          class="flex items-center gap-4 duration-200 ease-in-out hover:scale-105"
          href="https://tailwindcss.com"
        >
          <div
            class="flex h-14 w-14 items-center justify-center rounded-full bg-slate-800 dark:bg-slate-600"
          >
            <img
              class="w-1/2"
              src="/images/logos/tailwind.svg"
              alt="Tailwind CSS logo"
            />
          </div>
          Tailwind
        </a>
      </li>
      <li>
        <a
          class="flex items-center gap-4 duration-200 ease-in-out hover:scale-105"
          href="https://www.postgresql.org"
        >
          <div
            class="flex h-14 w-14 items-center justify-center rounded-full bg-slate-800 dark:bg-slate-600"
          >
            <img
              class="w-1/2"
              src="/images/logos/postgres.svg"
              alt="Postgres logo"
            />
          </div>
          Postgres
        </a>
      </li>
    </ul>
  </div>
  <div class="flex flex-col md:text-lg">
    <h1
      class="mb-6 w-max border-b-2 border-pink-300 pb-2 text-3xl font-semibold leading-tight text-slate-900 md:text-4xl dark:border-pink-400 dark:text-white"
    >
      Latest Writing
    </h1>
    <ul class="mb-6">
      {#each recentPosts as post}
        <li>
          <span class="mr-1" aria-hidden="true"> &#128279; </span>
          <a href={`/blog/${post.slug}`}>
            {post.title}
          </a>
        </li>
      {/each}
    </ul>
    <p>
      Check out the rest of my <a class="text-blue-500" href="/blog">
        blog posts
      </a>.
    </p>
  </div>
</div>
