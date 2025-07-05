import { mdsvex, escapeSvelte } from "mdsvex";
import { createHighlighter } from "shiki";
import adapter from "@sveltejs/adapter-auto";
import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

const langMap = {
  css: "CSS",
  go: "Go",
  html: "HTML",
  javascript: "JavaScript",
  js: "JavaScript",
  json: "JSON",
  jsx: "JavaScript",
  md: "Markdown",
  sql: "SQL",
  ts: "TypeScript",
  tsx: "TypeScript",
  typescript: "TypeScript"
};

const mdsvexOptions = {
  extensions: [".md"],
  highlight: {
    highlighter: async (code, lang = "text") => {
      const highlighter = await createHighlighter({
        themes: ["dracula"],
        langs: ["go", "typescript", "tsx"]
      });
      await highlighter.loadLanguage("go", "typescript", "tsx");
      const html = escapeSvelte(
        highlighter.codeToHtml(code, {
          lang,
          theme: "dracula",
          transformers: [
            {
              pre(node) {
                node.properties["data-lang"] = langMap[lang] || lang;
              }
            }
          ]
        })
      );
      return `{@html \`${html}\` }`;
    }
  }
};

/** @type {import('@sveltejs/kit').Config} */
const config = {
  // Consult https://svelte.dev/docs/kit/integrations
  // for more information about preprocessors
  preprocess: [vitePreprocess(), mdsvex(mdsvexOptions)],
  kit: {
    // adapter-auto only supports some environments, see https://svelte.dev/docs/kit/adapter-auto for a list.
    // If your environment is not supported, or you settled on a specific environment, switch out the adapter.
    // See https://svelte.dev/docs/kit/adapters for more information about adapters.
    adapter: adapter()
  },
  extensions: [".svelte", ".svx", ".md"]
};

export default config;
