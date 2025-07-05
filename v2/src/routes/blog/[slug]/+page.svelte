<script>
  let { data } = $props();

  let percent = $state(0);

  // Scroll progress logic
  $effect(() => {
    document.body.classList.add("overscroll-y-none");

    const updateScrollProgress = () => {
      const scrollTop =
        document.body.scrollTop || document.documentElement.scrollTop;
      const footerHeight = document.querySelector("footer")?.clientHeight || 0;
      const scrollHeight = document.documentElement.scrollHeight - footerHeight;
      const height = scrollHeight - document.documentElement.clientHeight;
      const scrollPercentage = (scrollTop / height) * 100;
      percent = Math.min(Math.max(0, scrollPercentage), 100);
    };

    window.addEventListener("scroll", updateScrollProgress);
    updateScrollProgress();

    return () => {
      window.removeEventListener("scroll", updateScrollProgress);
      document.body.classList.remove("overscroll-y-none");
    };
  });
</script>

<article class="mx-auto max-w-2xl overscroll-none">
  <div class="fixed inset-x-0 top-12">
    <div class="h-1 bg-pink-400" style={`width: ${percent}%`}></div>
  </div>

  <div class="mb-10">
    <h1
      class="mb-6 text-4xl font-semibold leading-none text-slate-900 md:text-5xl dark:text-white"
    >
      {data.post.title}
    </h1>

    <p class="text-xl leading-snug text-slate-400">
      {data.post.description}
    </p>

    <!-- <div class="mt-6 flex items-baseline text-sm text-slate-400">
      <time class="block" datetime={data.post.createdAt.String()}>
        {data.post.createdAt.Format("Jan 2, 2006")}
      </time>
      <p class="ml-4">{utils.GetReadingTime(article.Markdown)}</p>
    </div> -->
  </div>

  <div
    class="prose prose-slate prose-blue dark:prose-invert prose-a:no-underline prose-pre:bg-[#282a36] prose-pre:text-[#f8f8f2] prose-pre:pb-0 md:prose-code:max-h-[500px] prose-code:font-fira max-w-none md:text-lg"
  >
    <data.content />
  </div>
</article>
