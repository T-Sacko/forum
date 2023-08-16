var createPostButton = document.getElementById("create-post-button");




  var modalContainer = document.getElementById("modal-container");
  var closeButton = modalContainer.querySelector("#close-button");

  createPostButton.onclick = function () {
    console.log("clicked create post button")
    modalContainer.style.display = "flex";
  };

  closeButton.onclick = function () {
    modalContainer.style.display = "none";
  };

  window.onclick = function (event) {
    if (event.target == modalContainer) {
      modalContainer.style.display = "none";
    }
  }



  // Add event listener to the submit button
  var postForm = modalContainer.querySelector("#create-post-form");
  postForm.addEventListener("submit", async (event) => {
    event.preventDefault(); // Prevent the form from submitting
    resp = await fetch('/session')
    data = await resp.json()
    if (data.status) {
      postForm.submit()
      console.log("Cliked submit button")
    }
    // Check if the user is logged in
  });


