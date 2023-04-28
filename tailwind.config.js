/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.html", "./public/js/*"],
  theme: {
    extend: {
      fontFamily: {
        inter: ["Inter", "sans-serif"],
        fira: ["Fira Code", "monospace"],
      },
      typography(theme) {
        return {
          DEFAULT: {
            css: {
              "code::before": {
                content: "none",
              },
              "code::after": {
                content: "none",
              },
              code: {
                color: theme("colors.pink.400"),
              },
            },
          },
        };
      },
    },
  },
  plugins: [require("@tailwindcss/typography"), require("@tailwindcss/forms")],
};
