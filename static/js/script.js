
// get liked posts data
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
      const likeButton = document.getElementById(`${postId}-like`);
      if (value == 1) {
        if (likeButton) {
          toggleLiked(likeButton)
          console.log("toggleLiked it up still")
        }
      } else {
        const dislikeButton = document.getElementById(`${postId}-dislike`);
        toggleDisliked(dislikeButton)

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


function toggleDisliked(dislikeButton) {
  dislikeButton.classList.toggle('disliked')
  dislikeButton.classList.toggle('fa-thumbs-o-down')
  dislikeButton.classList.toggle('fa-thumbs-down')
}

function increment(postId, id) {
  const dislikeCountElement = document.getElementById(`${postId}-${id}`);
  const dislikeCountText = dislikeCountElement.textContent;
  const dislikeCountValue = parseInt(dislikeCountText);
  const incrementedCount = dislikeCountValue + 1;
  dislikeCountElement.textContent = incrementedCount.toString();
}

function decrement(postId, id) {
  const dislikeCountElement = document.getElementById(`${postId}-${id}`);
  const dislikeCountText = dislikeCountElement.textContent;
  const dislikeCountValue = parseInt(dislikeCountText);
  const decrementedCount = dislikeCountValue - 1;
  dislikeCountElement.textContent = decrementedCount.toString();
}



const likes = document.querySelectorAll('.likes')

likes.forEach(likeButton => {
  console.log("a post")
  const likeStr = likeButton.getAttribute('id');
  const postId = likeStr.split('-')[0]; // This will split the likeId string at the '-' character
  const dislikeButton = document.getElementById(`${postId}-dislike`)

  console.log(postId, "mans up inna da ting uno"); // Outputs: [part1, part2, ...]


  likeButton.addEventListener('click', () => {
    // toggle like button 
    toggleLiked(likeButton)



    if (likeButton.classList.contains('liked')) {
      // send like req to server
      console.log("we liking suttin")
      increment(postId, 1)
      if (dislikeButton.classList.contains('disliked')) {
        toggleDisliked(dislikeButton)
        decrement(postId, 2)
        // send req to remove dislike
        handleLikeAction(postId, "removeDislike")
      }
      handleLikeAction(postId, "like")
    } else {
      //send unlike req
      console.log("we unliking suttin")
      handleLikeAction(postId, "unlike")
      decrement(postId, 1)
    }



  })

  dislikeButton.addEventListener('click', () => {
    // toggle like button 
    toggleDisliked(dislikeButton)


    if (dislikeButton.classList.contains('disliked')) {
      // send dislike req to server
      console.log("we disliking suttin")
      increment(postId, 2)
      if (likeButton.classList.contains('liked')) {
        toggleLiked(likeButton)
        decrement(postId, 1)
        // send req to remove like
        handleLikeAction(postId, "unlike")
      }
      handleLikeAction(postId, "dislike")
    } else {
      //send remove dislike req
      decrement(postId, 2)
      console.log("we removing dislike")
      handleLikeAction(postId, "removeDislike")
    }


  })

})



function handleLikeAction(postId, action) {
  fetch('/handle-like-action', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ postId: postId, action: action })
  });
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

const signIn = document.getElementById('sign-in')

const signInModal = document.getElementById('signInModal')

const overlay = document.getElementById('overlay')

const signInCloseButton = document.getElementById('signInCloseButton')

signIn.addEventListener('click', () => {
  signInModal.style.display = 'block'
  overlay.style.display = 'block'
})

overlay.addEventListener('click', () => {
  signInModal.style.display = 'none';
  overlay.style.display = 'none';
});

signInCloseButton.addEventListener('click', () => {
  signInModal.style.display = 'none';
  overlay.style.display = 'none';
});
