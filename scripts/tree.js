document.addEventListener("DOMContentLoaded", function () {
  var folders = document.querySelectorAll(".folder");
  folders.forEach(function (folder) {
    folder.addEventListener("click", function () {
      this.parentElement.querySelector("ul").classList.toggle("hidden");
      this.textContent = this.textContent.startsWith("▶")
        ? "▼ " + this.textContent.slice(2)
        : "▶ " + this.textContent.slice(2);
    });
  });
});
