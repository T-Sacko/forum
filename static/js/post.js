var createPostButton = document.getElementById("create-post-button");
var modalContainer = document.getElementById("modal-container");
var closeButton = document.getElementById("close-button");

createPostButton.addEventListener("click", function() {
  modalContainer.style.display = "block";
});

closeButton.addEventListener("click", function() {
  modalContainer.style.display = "none";
});
