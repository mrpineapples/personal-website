document.addEventListener("submit", (e) => {
  if (!confirm("Are you sure you want to delete this article?")) {
    e.preventDefault();
    return;
  }

  const { articleId } = e.target.dataset;
  localStorage.removeItem(`smde_edit-article-${articleId}`);
});
