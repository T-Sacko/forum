var createPostButton = document.getElementById("create-post-button");

var isLoggedIn = document.cookie.indexOf("session=") !== -1;



var modalContainer = document.getElementById("modal-container");
var closeButton = document.getElementById("close-button");

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

// create post function ----------------------

function checkSession() {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/create-post");
  console.log("OPened", xhr.status)

  xhr.onreadystatechange = function () {
    if (xhr.readyState == 4) {
      console.log("response received")
      var response = JSON.parse(xhr.responseText);
      if (response.loggedIn) {
        // User is logged in
        console.log("User is authorized to post");
        // Continue with posting the form data or other actions
        document.getElementById("create-post-form").submit();
      } else {
        console.log("User is not authorized to post");
        var errorMessage = document.getElementById("error-message");
        errorMessage.style.color = "red"
        errorMessage.textContent = "You are not logged in";
        errorMessage.style.display = "block";
        // Handle the request error
        console.log("Request failed with status: " + xhr.status);

      }

    };
  }
  xhr.send();
}
  // Add event listener to the submit button
  var submitButton = document.getElementById("submit-button");
  submitButton.addEventListener("click", function (event) {
    event.preventDefault(); // Prevent the form from submitting
    console.log("Cliked submit button")
    // Check if the user is logged in
    checkSession();
  });

// const likes = document.querySelectorAll('.likes')

// likes.forEach(like => {
//   console.log("yhh")
//   const postId = like.getAttribute('id')
//   like.addEventListener('click', () => {
//     console.log(postId)
//     likePost(postId)
//   })
// })

