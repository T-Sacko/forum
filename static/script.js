// script.js
function checkUsername(username) {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/check-username?username=" + encodeURIComponent(username), true);
  xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE) {
      if (xhr.status === 200) {
        var response = JSON.parse(xhr.responseText);
        var errorElement = document.getElementById("username-error");
        if (response.available) {
          errorElement.textContent = ""; // Clear any previous error message
        } else {
          errorElement.setAttribute("style", "display:block;")
          errorElement.textContent = "Error:Username Already Exists";
        }
      } else {
        console.error("Error:", xhr.status);
      }
    }
  };
  xhr.send();
}


// document.getElementById('signupForm').addEventListener('submit', function(e) {
//     /*This function should check with Go as to whether or not the input is valid through the DB Check*/
//     e.preventDefault(); // Prevent the form from submitting normally

//     var userId = 123; // Replace with the actual user ID

//     // Build the URL with the user ID and set it as the form's action
//     this.action = 'https://example.com/user/' + userId;

//     // Submit the form
//     this.submit();
// });

// document.getElementById('loginForm').addEventListener('submit', function(e) {
//     /*This function should check with Go as to whether or not the input is valid through the DB Check*/
//     e.preventDefault(); // Prevent the form from submitting normally

//     var userId = 123; // The ID will come from the Go backend

//     // Build the URL with the user ID and set it as the form's action
//     this.action = 'https://example.com/user/' + userId;

//     // Submit the form
//     this.submit();
// });


//NavBar
function IconBar(){
  var iconBar = document.getElementById("iconBar");
  var navigation = document.getElementById("navigation");
  if (navigation.classList.contains("hide")) {
    iconBar.setAttribute("style", "display:none;");
    navigation.classList.remove("hide");
  } else {
    iconBar.setAttribute("style", "display:block;");
    navigation.classList.add("hide");
  }
}

function incrementLikes() {
    var likeCountElement = document.getElementById("likeCount");
    var likeCount = parseInt(likeCountElement.innerHTML);
    likeCount++;
    likeCountElement.innerHTML = likeCount;
}
function incrementDislikes() {
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
