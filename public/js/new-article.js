const easyMDE = new EasyMDE({
  autosave: {
    delay: 1000,
    enabled: true,
    text: "Last Saved ",
    uniqueId: "new-article",
  },
  element: document.getElementById("markdown"),
  forceSync: true,
  maxHeight: "500px",
});
