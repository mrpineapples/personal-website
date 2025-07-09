import { posts } from "$lib/data/posts";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {
  return {
    posts,
    meta: {
      title: "Michael Miranda | Blog",
      description: "Read the latest posts from Michael Miranda"
    }
  };
};
