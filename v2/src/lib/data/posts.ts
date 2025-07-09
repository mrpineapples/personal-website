import { browser } from "$app/environment";
import { format } from "date-fns";
import readingTime from "reading-time";

if (browser) {
  throw new Error("posts can only be imported server-side");
}

type MDsveXPost = {
  metadata: any;
  default: any;
};

type Post = {
  date: string;
  title: string;
  description: string;
  isIndexFile: boolean;
  readingTime: {
    text: string;
    minutes: number;
    time: number;
    words: number;
  };
  slug: string;
  next?: Post;
  previous?: Post;
};

const addTimezoneOffset = (date: Date) => {
  const offsetInMilliseconds = new Date().getTimezoneOffset() * 60 * 1000;
  return new Date(new Date(date).getTime() + offsetInMilliseconds);
};

const rawFiles = import.meta.glob<string>("/src/posts/**/*.md", {
  eager: true,
  query: "?raw",
  import: "default"
});

export const posts: Post[] = Object.entries<MDsveXPost>(
  import.meta.glob("/src/posts/**/*.md", { eager: true })
)
  .map(([filepath, post]) => {
    const rawText = rawFiles[filepath];
    return {
      ...post.metadata,
      isIndexFile: filepath.endsWith("/index.md"),
      date: post.metadata.date
        ? format(
            addTimezoneOffset(new Date(post.metadata.date)),
            "MMM dd, yyyy"
          )
        : undefined,
      readingTime: readingTime(rawText),
      slug: filepath
        .replace(/(\/index)?\.md/, "")
        .split("/")
        .pop()
    };
  })
  .sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime())
  .map((post, index, allPosts) => ({
    ...post,
    next: allPosts[index - 1],
    previous: allPosts[index + 1]
  }));

export const recentPosts = posts.slice(0, 3);
