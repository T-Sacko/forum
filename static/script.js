// script.js
const signupLink = document.getElementById('sign-in');
const registerLink =document.getElementById('register_btn')
const loginLink = document.getElementById('login_btn')
const popupContainer = document.getElementById('popup-container');
const login = document.getElementById('login')
const signup = document.getElementById('signup')
const createPostButton = document.getElementById("create-post-button");
const modalContainer = document.getElementById("modal-container");
const closeButton = document.getElementById("close-button");
const errorElement = document.getElementById("username-error");

createPostButton.onclick = function () {
  console.log("clicked create post button")
  modalContainer.style.display = "flex";
};

closeButton.onclick = function () {
  modalContainer.style.display = "none";
};

window.onclick = function(event) {
  if (event.target == modalContainer) {
    modalContainer.style.display = "none";
  }
}

signupLink.addEventListener('click', function() {
    if (popupContainer.style.display === 'flex') {
        popupContainer.style.display = 'none';
        return
    }
    if (signup.style.display === 'block') {
        signup.style.display = 'none'
    }
    popupContainer.style.display = 'flex';
    login.style.display = 'block'
});

function closePopup() {
    popupContainer.style.display = "none";
}

registerLink.addEventListener('click', function() {
    login.style.display = 'none'
    signup.style.display = 'block'
});

loginLink.addEventListener('click', function() {
    signup.style.display = 'none'
    login.style.display = 'block'
});

  function checkSession() {
  let xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/create-post");
  xhr.onreadystatechange = function () {
    if (xhr.readyState == 4 && xhr.status == 200) {
      let response = JSON.parse(xhr.responseText);
      if (response.loggedIn) {
        // User is logged in
          let signOutcode = `<a class="sign-in" id="sign-in"><i class="fa fa-sign-out"></i>Sign Out</a>`
          console.log(document.getElementById("sign-in"))
          document.getElementById("sign-in").innerHTML = signOutcode
        // Continue with posting the form data or other actions
        document.getElementById("create-post-form").submit();
      } else {
        // User is not logged in
        const errorMessage = document.getElementById("error-message");
        errorMessage.style.color = "red"
        errorMessage.textContent = "You are not logged in";
        errorMessage.style.display = "block";
        // Hide the error message after 3 seconds
      }
    } else {
      // Handle the request error
      console.log("Request failed with status: " + xhr.status);
    }
  };
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


function checkUsername(username) {
  if (username === '') {
    errorElement.style.display = "none";
  }
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/check-username?username=" + encodeURIComponent(username), true);
  xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE) {
      if (xhr.status === 200) {
        var response = JSON.parse(xhr.responseText);
        if (response.available) {
          errorElement.textContent = "";
          errorElement.style.display = "none"; // Clear any previous error message
        } else {
          errorElement.style.display = "block";
          errorElement.textContent = "Error:Username Already Exists";
        }
      } else {
        console.error("Error:", xhr.status);
      }
    }
  };
  xhr.send();
}
function checkEmail(email) {
  if (email === '') {
    errorElement.style.display = "none";
  }
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/check-email?email=" + encodeURIComponent(email), true);
  xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE) {
      if (xhr.status === 200) {
        var response = JSON.parse(xhr.responseText);
        if (response.available) {
          errorElement.textContent = "";
          errorElement.style.display = "none";
          // Clear any previous error message
        } else {
          errorElement.style.display = "block";
          errorElement.textContent = "Error:Email Already In Use";
        } 
      } else {
        console.error("Error:", xhr.status);
      }
    }
  };
  xhr.send();
}
//NavBar
function incrementLikes(postID) {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/likes?email=" + encodeURIComponent(postID), true);
    var likeCountElement = document.getElementById("likeCount");
    var likeCount = parseInt(likeCountElement.innerHTML);
    likeCount++;
    likeCountElement.innerHTML = likeCount;
}
function incrementDislikes(postID) {
    var dislikeCountElement = document.getElementById("dislikeCount");
    var dislikeCount = parseInt(dislikeCountElement.innerHTML);
    dislikeCount++;
    dislikeCountElement.innerHTML = dislikeCount;
}
function showComment(){
    var commentArea = document.getElementById("comment-area");
    if (commentArea.classList.contains("hide")) {
        commentArea.classList.remove("hide");
    } else {
        commentArea.classList.add("hide");
    }
    
}
