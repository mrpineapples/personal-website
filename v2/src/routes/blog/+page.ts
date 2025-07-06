import type { PageLoad } from "./$types";

export const load: PageLoad = ({ data }) => {
  return {
    ...data,
    meta: {
      title: "Michael Miranda | Blog",
      description: "Read the latest posts from Michael Miranda"
    }
  };
};
