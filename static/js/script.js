
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




var usernameAvailable = false
var emailAvailable = false
//////////////////////////////////////////////////////////////////
function checkUsername(username) {
  return new Promise((resolve, reject) => {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/check-username?username=" + encodeURIComponent(username), true);
    xhr.onreadystatechange = function () {
      if (xhr.readyState === XMLHttpRequest.DONE) {
        if (xhr.status === 200) {
          var response = JSON.parse(xhr.responseText);
          var errorElement = document.querySelector(".username-error");
          if (response.available) {
            usernameAvailable = true
            errorElement.style.display='none'; // Clear any previous error message
            resolve(true); // Username is available
          } else {
            usernameAvailable = false
            errorElement.setAttribute("style", "display:block;");
            errorElement.textContent = "Error: Username Already Exists";
            resolve(false); // Username is not available
          }
        } else {
          console.error("Error:", xhr.status);
          reject(new Error("An error occurred while checking username availability."));
        }
      }
    };
    xhr.send();
  });
}


function checkEmail(email) {
  return new Promise((resolve, reject) => {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/check-email?email=" + encodeURIComponent(email), true);
    xhr.onreadystatechange = function () {
      if (xhr.readyState === XMLHttpRequest.DONE) {
        if (xhr.status === 200) {
          var response = JSON.parse(xhr.responseText);
          var errorElement = document.querySelector(".email-error");
          if (response.available) {
            emailAvailable = true
            errorElement.style.display='none'; // Clear any previous error message
            resolve(true); // Email is available
          } else {
            emailAvailable = false
            errorElement.setAttribute("style", "display:block;");
            errorElement.textContent = "Error: Email Already In Use";
            resolve(false); // Email is not available
          }
        } else {
          console.error("Error:", xhr.status);
          reject(new Error("An error occurred while checking email availability."));
        }
      }
    };
    xhr.send();
  });
}


// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// ------------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// login authentication

document.getElementById('loginForm').addEventListener('submit', function (event) {
  event.preventDefault(); // Prevent the form from submitting normally

  const formData = new FormData(event.target); // Get form data
  const email = formData.get('email');
  const password = formData.get('password');

  // Create the payload to send in the AJAX request
  const payload = {
    email: email,
    password: password
  };



  // Make the AJAX request to the server
  fetch('/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(payload)
  })
    .then(response => response.json())
    .then(data => {
      // Handle the response from the server
      if (data.loggedIn) {
        // If the login is successful, redirect the user to a new page or perform other actions
        console.log('Login successful!');
        window.location.href = '/'; // Redirect to the dashboard page after successful login
      } else {
        // If the login fails, display an error message
        console.log('Login failed:', data);
        const errorElement = document.querySelector('.login-email-error');
        errorElement.setAttribute('style', 'display:block;');
        errorElement.textContent = 'Error: Invalid email or password';
      }
    })
    .catch(error => {
      console.error('Error:', error);
    });
});

// --------------------------------------------------------------------------

//signup authentication

document.getElementById('signupForm').addEventListener('submit', (event) => {
  event.preventDefault()
  console.log('trying to signup una',usernameAvailable, 'thats username and this', emailAvailable,' is email')
 if (usernameAvailable&&emailAvailable){
  event.target.submit()
 
  }  if (!usernameAvailable){
  const usernameError = document.getElementById('username-error')
  usernameError.style.display='block'
  usernameError.textContent='Username is already taken'
 } if (!emailAvailable){
  const emailError = document.getElementById('email-error')
  emailError.style.display='block'
  emailError.textContent='Email is already taken'
 }

  

})

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// ------------------------------------------------------------------------------
// -----------------------------------------------------------------------------


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

const loginButton = document.getElementById('loginButton')

const signupButton = document.getElementById('signupButton')

const signupForm = document.getElementById('signupForm')

const loginForm = document.getElementById('loginForm')

const overlay = document.getElementById('overlay')

const signupCloseButton = document.getElementById('signInCloseButton')
const loginCloseButton = document.getElementById('loginCloseButton')

signIn.addEventListener('click', () => {
  loginForm.style.display = 'block'
  overlay.style.display = 'block'
  signupForm.style.display = 'none'
})

overlay.addEventListener('click', () => {
  signupForm.style.display = 'none';
  overlay.style.display = 'none';
  loginForm.style.display = 'none'
});

signupCloseButton.addEventListener('click', () => {
  signupForm.style.display = 'none';
  overlay.style.display = 'none';
});

loginCloseButton.addEventListener('click', () => {
  loginForm.style.display = 'none';
  overlay.style.display = 'none';
});

loginButton.addEventListener('click', () => {
  signupForm.style.display = 'none'
  loginForm.style.display = 'block'
})

signupButton.addEventListener('click', () => {
  signupForm.style.display = 'block'
  loginForm.style.display = 'none'
})
