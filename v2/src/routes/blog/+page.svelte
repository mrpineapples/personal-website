<script lang="ts">
  export let articles: any[] = []; // Array of article objects passed as a prop
  export let isAdmin = false; // Admin flag passed as a prop

  // Function to format date (similar to Go's "Jan 2, 2006" format)
  const formatDate = (dateString: string): string => {
    const date = new Date(dateString);
    return date.toLocaleDateString("en-US", {
      year: "numeric",
      month: "short",
      day: "numeric"
    });
  };

  // Function to calculate reading time (replace with your actual implementation)
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

            {#if isAdmin}
              <div class="mt-4 flex">
                <a
                  class="mr-4 hover:text-blue-500"
                  href={`/admin/blog/${article.Slug}/edit`}
                >
                  Edit
                </a>

                <form
                  action={`/articles/${article.ID}?_method=DELETE`}
                  data-article-id={article.ID}
                  method="post"
                  on:submit|preventDefault={(e) => {
                    if (
                      confirm("Are you sure you want to delete this article?")
                    ) {
                      localStorage.removeItem(
                        `smde_edit-article-${e.target.dataset.articleId}`
                      );
                      e.target.submit();
                    }
                  }}
                >
                  <button class="text-red-500" type="submit"> Delete </button>
                </form>
              </div>
            {/if}
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
