import type { PageLoad } from "./$types";

export const load: PageLoad = ({ data }) => {
  return {
    ...data,
    meta: {
      title: "Michael Miranda"
    }
  };
};
