import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = () => {
  return {
    meta: {
      title: "Michael Miranda | About",
      description: "Learn more about Michael Miranda"
    }
  };
};
