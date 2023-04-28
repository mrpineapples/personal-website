const { id, markdown } = window._data;

const easyMDE = new EasyMDE({
  autosave: {
    delay: 1000,
    enabled: true,
    text: "Last Saved ",
    uniqueId: `edit-article-${id}`,
  },
  element: document.getElementById("markdown"),
  forceSync: true,
  maxHeight: "500px",
  initialValue: markdown,
});

const cancelBtn = document.querySelector("#cancel-edit");
cancelBtn.addEventListener("click", () => {
  easyMDE.clearAutosavedValue();
});
