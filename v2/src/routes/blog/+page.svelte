<script lang="ts">
  let { articles = $bindable([]), isAdmin = $bindable(false) } = $props<{
    articles?: any[];
    isAdmin?: boolean;
  }>();

  const formatDate = (dateString: string): string => {
    const date = new Date(dateString);
    return date.toLocaleDateString("en-US", {
      year: "numeric",
      month: "short",
      day: "numeric"
    });
  };

  const getReadingTime = (markdown: string) => {
    const words = markdown.split(/\s+/).length;
    const minutes = Math.ceil(words / 200);
    return `${minutes} min read`;
  };
</script>

<div class="mx-auto max-w-4xl md:min-h-screen">
  <h1
    class="mb-5 text-4xl font-semibold text-slate-900 md:text-5xl dark:text-white"
  >
    All Posts 📝
  </h1>

  <ul>
    {#if articles.length > 0}
      {#each articles as article}
        <li
          class="flex flex-wrap gap-x-8 gap-y-4 border-t border-slate-300 py-8 leading-normal md:flex-nowrap"
        >
          <time
            class="w-full shrink-0 text-slate-500 md:w-28 dark:text-slate-400"
            datetime={article.CreatedAt}
          >
            {formatDate(article.CreatedAt)}
          </time>

          <div class="flex flex-col">
            <h2
              class="mb-4 text-2xl font-medium text-slate-900 dark:text-white"
            >
              {article.Title}
            </h2>

            {#if article.Description}
              <p class="mb-4 text-slate-500 dark:text-slate-400">
                {article.Description}
              </p>
            {/if}

            <a class="max-w-max text-blue-500" href={`/blog/${article.Slug}`}>
              Read More ({getReadingTime(article.Markdown)})
              <span aria-hidden="true">→</span>
            </a>
          </div>
        </li>
      {/each}
    {:else}
      <li class="mb-8 text-2xl leading-normal">
        No articles exist, how sad 😢
      </li>
    {/if}
  </ul>
</div>
