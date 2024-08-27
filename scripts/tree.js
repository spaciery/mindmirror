const folders = document.querySelectorAll(".folder");

for (const folder of folders) {
  folder.onclick = () => {
    folder.classList.toggle("active");
    const target_list = folder.nextElementSibling;
    target_list.style.maxHeight = target_list.style.maxHeight ? null : "100vh";
  };
}
