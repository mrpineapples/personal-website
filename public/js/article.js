const copyButtonLabel = "Copy Code";
const blocks = document.querySelectorAll(".prose pre");

blocks.forEach((block) => {
  const lang = block.parentElement.getAttribute("data-lang");
  const header = document.createElement("div");
  header.className = "font-inter flex items-baseline mb-4 select-none";
  const languageHeading = document.createElement("p");
  languageHeading.className =
    "text-base m-0 px-4 pb-1 border-b border-[#ff79c6] md:text-lg";
  languageHeading.innerText = lang || "Text";
  header.append(languageHeading);
  block.tabIndex = "-1";
  block.prepend(header);

  if (navigator.clipboard) {
    const button = document.createElement("button");
    button.className = "text-xs ml-auto p-2 border rounded-lg cursor-pointer";
    button.innerText = copyButtonLabel;
    button.addEventListener("click", async () => {
      await copyCode(block, button);
    });

    header.append(button);
  }
});

const copyCode = async (block, button) => {
  const code = block.querySelector("code");
  const text = code.innerText.replace(/^\s*[\r\n]/gm, "");
  await navigator.clipboard.writeText(text);
  button.innerText = "Copied!";

  setTimeout(() => {
    button.innerText = copyButtonLabel;
  }, 750);
};

const scrollProgress = () => {
  document.body.classList.add("overscroll-y-none");
  return {
    init() {
      window.addEventListener("scroll", () => {
        const scrollTop =
          document.body.scrollTop || document.documentElement.scrollTop;
        const height =
          document.documentElement.scrollHeight -
          document.querySelector("footer").clientHeight -
          document.documentElement.clientHeight;
        const scrollPercentage = (scrollTop / height) * 100;
        this.percent = Math.min(Math.max(0, scrollPercentage), 100);
      });
    },
    percent: 0,
  };
};
