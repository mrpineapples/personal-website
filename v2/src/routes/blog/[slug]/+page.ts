import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ data }) => {
  const { post } = data;
  const content = post.isIndexFile
    ? await import(`../../../posts/${post.slug}/index.md`)
    : await import(`../../../posts/${post.slug}.md`);

  return {
    ...data,
    content: content.default,
    meta: {
      title: post.title,
      description: post.description
    }
  };
};
