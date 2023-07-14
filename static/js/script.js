fetch("/get-post-likes", {
  method: "GET"
})
  .then(response => response.json())
  .then(data => {
    // Handle the like data
    console.log(data, 'this is the data');

    for (const likeData of data) {
      const postId = likeData.postId;
      const value = likeData.value;
      if (value === 1) {
        const likeButton = document.getElementById(`${postId}-like`);
        if (likeButton) {
          toggleLiked(likeButton)
          console.log("toggleLiked it up still")
        }
      } else {
        const dislikeButton = document.getElementById(`${postId}-dislike`);
        disliked(dislikeButton)

      }

    }
  })
  .catch(error => {
    console.error("Error with the posts toggleLiked data is:", error);
    // Handle the error
  });

function toggleLiked(likeButton) {
  likeButton.classList.toggle('liked')
  likeButton.classList.toggle('fa-thumbs-o-up')
  likeButton.classList.toggle('fa-thumbs-up');
}

function unlike(likeButton) {
  likeButton.classList.toggle('liked')
  likeButton.classList.toggle('fa-thumbs-o-up')
  likeButton.classList.toggle('fa-thumbs-up');
}

function disliked(dislikeButton) {
  dislikeButton.classList.toggle('disliked')
  dislikeButton.classList.toggle('fa-thumbs-o-down')
  dislikeButton.classList.toggle('fa-thumbs-down')
}


const likes = document.querySelectorAll('.likes')

likes.forEach(likeButton => {
  console.log("yhh")
  var str = likeButton.getAttribute('id');
  var postId = str.split('-')[0]; // This will split the likeId string at the '-' character

  console.log(postId, "mans up inna da ting uno"); // Outputs: [part1, part2, ...]


  likeButton.addEventListener('click', () => {
    // toggle like button 

    toggleLiked(likeButton)
    if (likeButton.classList.contains('liked')) {
      // send unlike req to server
      unlikePost(postId)
    } else {
      //send like req
      console.log("unliked to liked")
      likePost(postId)
    }


  })
})

function unlikePost()

function likePost(postId) {

  fetch('/like-post', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ postId: postId })
  })
}



//////////////////////////////////////////////////////////////////
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
function checkEmail(email) {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/check-email?email=" + encodeURIComponent(email), true);
  xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE) {
      if (xhr.status === 200) {
        var response = JSON.parse(xhr.responseText);
        var errorElement = document.getElementById("username-error");
        if (response.available) {
          errorElement.textContent = ""; // Clear any previous error message
        } else {
          errorElement.setAttribute("style", "display:block;")
          errorElement.textContent = "Error:Email Already In Use";
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
function IconBar() {
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
function showComment() {
  var commentArea = document.getElementById("comment-area");
  if (commentArea.classList.contains("hide")) {
    commentArea.classList.remove("hide");
  } else {
    commentArea.classList.add("hide");
  }

}
