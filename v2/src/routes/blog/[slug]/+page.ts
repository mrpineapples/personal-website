export async function load({ data }: any) {
  // load the markdown file based on slug
  const content = data.post.isIndexFile
    ? await import(`../../../../posts/${data.post.slug}/index.md`)
    : await import(`../../../../posts/${data.post.slug}.md`);

  return {
    post: data.post,
    content: content.default
  };
}
