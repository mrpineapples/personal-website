import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ data }) => {
  const content = data.post.isIndexFile
    ? await import(`../../../../posts/${data.post.slug}/index.md`)
    : await import(`../../../../posts/${data.post.slug}.md`);

  return {
    post: data.post,
    content: content.default
  };
};
