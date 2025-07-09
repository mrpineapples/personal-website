import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = () => {
  return {
    meta: {
      title: "Michael Miranda | Contact",
      description: "Get in touch with Michael Miranda"
    }
  };
};
