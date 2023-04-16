/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.html", "./public/js/*"],
  theme: {
    extend: {
      fontFamily: {
        inter: ["Inter", "sans-serif"],
        fira: ["Fira Code", "monospace"],
      },
      typography() {
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
                color: "#ff79c6",
              },
            },
          },
        };
      },
    },
  },
  plugins: [require("@tailwindcss/typography"), require("@tailwindcss/forms")],
};
