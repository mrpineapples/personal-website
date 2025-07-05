<script lang="ts">
  import Footer from "$lib/components/Footer.svelte";
  import Header from "$lib/components/Header.svelte";
  import "../app.css";

  let { children } = $props();
</script>

<Header />
<main class="px-5 py-10">
  {@render children()}
</main>
<Footer />
