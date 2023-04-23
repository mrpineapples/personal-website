const menu = document.querySelector("nav button");
const navLinks = document.querySelector("#mobile-nav-links");
menu.addEventListener("click", () => {
  navLinks.classList.toggle("invisible");
  navLinks.classList.toggle("animate-fade-in");
  navLinks.ariaHidden = navLinks.classList.contains("invisible");
  document.body.classList.toggle("overflow-hidden");
});
