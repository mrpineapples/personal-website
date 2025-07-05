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

  $effect(() => {
    const copyButtonLabel = "Copy Code";
    const blocks = document.querySelectorAll(".prose pre");

    blocks.forEach((block) => {
      const lang = block.getAttribute("data-lang");
      const header = document.createElement("div");
      header.className = "font-inter flex items-baseline mb-4 select-none";
      const languageHeading = document.createElement("p");
      languageHeading.className =
        "text-base !m-0 px-4 pb-1 border-b border-[#ff79c6] md:text-lg";
      languageHeading.innerText = lang || "Text";
      header.append(languageHeading);
      block.tabIndex = "-1";
      block.prepend(header);

      if (navigator.clipboard) {
        const button = document.createElement("button");
        button.className =
          "text-xs ml-auto p-2 border rounded-lg cursor-pointer";
        button.innerText = copyButtonLabel;
        button.addEventListener("click", async () => {
          await copyCode(block, button);
        });

        header.append(button);
      }
    });

    const copyCode = async (block, button) => {
      const code = block.querySelector("code");
      const text = code.innerText.replace(/^\s*[\r\n]/gm, "");
      await navigator.clipboard.writeText(text);
      button.innerText = "Copied!";

      setTimeout(() => {
        button.innerText = copyButtonLabel;
      }, 750);
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
