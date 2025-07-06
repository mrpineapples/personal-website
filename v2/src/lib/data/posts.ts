import path from "path";
import fs from "fs";
import { browser } from "$app/environment";
import { format } from "date-fns";
import readingTime from "reading-time";

if (browser) {
  throw new Error(`posts can only be imported server-side`);
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

const getMarkdownContent = (filepath: string): string => {
  const fullPath = path.resolve("posts", filepath.replace(/^\/posts\//, ""));
  return fs.readFileSync(fullPath, "utf-8");
};

const addTimezoneOffset = (date: Date) => {
  const offsetInMilliseconds = new Date().getTimezoneOffset() * 60 * 1000;
  return new Date(new Date(date).getTime() + offsetInMilliseconds);
};

export const posts: Post[] = Object.entries<MDsveXPost>(
  import.meta.glob("/posts/**/*.md", { eager: true })
)
  .map(([filepath, post]) => {
    console.log("post", post);
    return {
      ...post.metadata,
      isIndexFile: filepath.endsWith("/index.md"),
      date: post.metadata.date
        ? format(
            addTimezoneOffset(new Date(post.metadata.date)),
            "MMM dd, yyyy"
          )
        : undefined,
      readingTime: readingTime(getMarkdownContent(filepath)),
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
