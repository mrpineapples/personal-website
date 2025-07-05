<script lang="ts">
  let isOpen = false;

  // Toggle the menu and handle body overflow
  const toggleMenu = () => {
    isOpen = !isOpen;
    if (isOpen) {
      document.body.classList.add("overflow-hidden");
    } else {
      document.body.classList.remove("overflow-hidden");
    }
  };
</script>

<nav class="sticky top-0 z-10 w-full bg-slate-800 px-5 text-slate-200">
  <div class="mx-auto flex h-12 max-w-4xl items-center justify-between">
    <div class="font-extrabold">
      <a class="text-white hover:text-pink-400" href="/">Michael Miranda</a>
    </div>
    <div class="hidden space-x-8 md:flex">
      <ul class="list-reset flex flex-1 items-center gap-x-5 md:flex-none">
        <li>
          <a class="inline-block hover:text-pink-400" href="/blog">Blog</a>
        </li>
        <li>
          <a class="inline-block hover:text-pink-400" href="/about">About</a>
        </li>
        <li>
          <a class="inline-block hover:text-pink-400" href="/contact">Contact</a
          >
        </li>
      </ul>
    </div>
    <button
      class="relative z-10 h-10 w-10 focus:outline-none md:hidden"
      on:click={toggleMenu}
    >
      <span class="sr-only">Open main menu</span>
      <div class="absolute left-1/2 top-1/2 w-5">
        <span
          aria-hidden="true"
          class="absolute block h-0.5 w-5 transform bg-current transition duration-500 ease-in-out"
          class:rotate-45={isOpen}
          class:-translate-y-1.5={!isOpen}
        ></span>
        <span
          aria-hidden="true"
          class="absolute block h-0.5 w-5 transform bg-current transition duration-500 ease-in-out"
          class:opacity-0={isOpen}
        ></span>
        <span
          aria-hidden="true"
          class="absolute block h-0.5 w-5 transform bg-current transition duration-500 ease-in-out"
          class:-rotate-45={isOpen}
          class:translate-y-1.5={!isOpen}
        ></span>
      </div>
    </button>
    {#if isOpen}
      <div class="absolute left-0 top-0 h-[100vh] w-full bg-slate-800">
        <ul
          class="list-reset my-auto flex h-full flex-col items-center justify-center gap-y-5"
        >
          <li>
            <a href="/blog">Blog</a>
          </li>
          <li>
            <a href="/about">About</a>
          </li>
          <li>
            <a href="/contact">Contact</a>
          </li>
        </ul>
      </div>
    {/if}
  </div>
</nav>
