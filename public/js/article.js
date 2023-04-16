const copyButtonLabel = "Copy Code";
const blocks = document.querySelectorAll(".prose pre");

blocks.forEach((block) => {
  const lang = block.parentElement.getAttribute("data-lang");
  const header = document.createElement("div");
  header.className = "font-inter flex items-baseline mb-4 select-none";
  const languageHeading = document.createElement("p");
  languageHeading.className = "text-lg m-0 px-4 pb-1 border-b border-[#ff79c6]";
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
