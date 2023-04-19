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

const deleteForm = document.querySelector("#delete-article-form");
const cancelBtn = document.querySelector("#cancel-edit");
deleteForm.addEventListener("submit", (e) => {
  if (!confirm("Are you sure you want to delete this article?")) {
    e.preventDefault();
    return;
  }
});
cancelBtn.addEventListener("click", () => {
  easyMDE.clearAutosavedValue();
});
