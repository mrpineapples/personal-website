import { recentPosts } from "$lib/data/posts";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {
  return {
    recentPosts
  };
};
