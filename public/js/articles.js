document.addEventListener("submit", (e) => {
  const { articleId } = e.target.dataset;
  localStorage.removeItem(`smde_edit-article-${articleId}`);
});
