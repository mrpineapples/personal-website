<script lang="ts">
  import Footer from "$lib/components/Footer.svelte";
  import Header from "$lib/components/Header.svelte";
  import "../app.css";

  import { page } from "$app/state";

  type HeadMeta = {
    description?: string;
    image?: string;
    title?: string;
  };

  console.log("page", page);
  const { children } = $props();
  const head = $state<HeadMeta>({});

  const setHead = () => {
    if (page.data?.meta?.title) {
      head.title = page.data?.meta?.title;
    } else {
      head.title = "Michael Miranda";
    }

    if (page.data?.meta?.description) {
      head.description = page.data?.meta?.description;
    }

    if (page.data?.meta?.image) {
      head.image = page.data?.meta?.image;
    } else {
      head.image = "https://michaelmiranda.dev/images/share.png";
    }
  };

  setHead();
  $effect(setHead);
</script>

<svelte:head>
  <title>{head.title}</title>
  <meta property="og:title" content={head.title} />
  <meta name="twitter:title" content={head.title} />
  {#if head.description}
    <meta name="description" content={head.description} />
    <meta property="og:description" content={head.description} />
    <meta name="twitter:description" content={head.description} />
  {/if}
  <meta property="og:image" content={head.image} />
  <meta name="twitter:image" content={head.image} />
</svelte:head>

<Header />
<main class="px-5 py-10">
  {@render children()}
</main>
<Footer />
