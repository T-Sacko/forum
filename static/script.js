// script.js
const signupLink = document.getElementById('sign-in');
const registerLink =document.getElementById('register_btn')
const loginLink = document.getElementById('login_btn')
const popupContainer = document.getElementById('popup-container');
const popupDiv = document.getElementById('container-div')
const login = document.getElementById('login')
const signup = document.getElementById('signup')
const createPostButton = document.getElementById("create-post-button");
const modalContainer = document.getElementById("modal-container");
const closeButton = document.getElementById("close-button");
const errorElement = document.getElementById("error");
const accountLink = document.getElementById("account-link")
const signOut = document.getElementById("sign_out")
const commentBtn = document.getElementById("comment-btn")


function comment() {
  const postId = document.getElementsByClassName("posts").id
  const comment = document.querySelector(".comment-box").value;

  console.log("Here is the id",comment.id)

  const requestBody = {
    postId: postId,
    comment: comment
  };
  console.log(postId)

  console.log(requestBody)
  fetch('/comment', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(requestBody)
  }).then(response=> {
    if (response.ok) {
      console.log("postId and Comment submitted")
    } else {
      console.log("failed to submit comment")
    }
      location.reload();

    }).catch(error => {
      console.error('Error occurred while deleting user session:', error);
    });
}

createPostButton.onclick = function () {
  console.log("clicked create post button")
  modalContainer.style.display = "flex";
};

window.onclick = function(event) {
  if (event.target == modalContainer) {
    modalContainer.style.display = "none";
  }
}

signupLink.addEventListener('click', function() {
  if (signupLink.innerText === 'Sign Out') {
      const close = document.getElementById('close')
      popupContainer.style.display = 'flex'
      popupDiv.style.height = 'auto'
      close.style.display = 'none'
      signOut.style.display = 'block'
      return
  }
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


function SignOut() {
  console.log("function being called")
    fetch('/del-cookie', {
        method: 'POST',
  }).then(response=> {
    if (response.ok) {
      let signIncode = `<a class="sign-in" id="sign-in">Sign In<i class="fa fa-sign-in" id="sign-i"></i></a>`
      let close = document.getElementById("close")
      document.getElementById("sign-in").innerHTML = signIncode
      accountLink.style.display = 'none'
      createPostButton.style.display = 'none'
      popupContainer.style.display = 'none';
      signOut.style.display = "none"
      close.style.display = "block"
    } else {
      console.log("failed to execute deletion")
    }})
    .catch(error => {
      console.error('Error occurred while deleting user session:', error);
    });
};      


function closePopup() {
    popupContainer.style.display = "none";
    modalContainer.style.display = "none";
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
  console.log('checkSession function being called')
  fetch('/check-session', {
    method: 'POST',
    credentials: 'include' // Include cookies for session tracking
  })
  .then(response => {
    if (response.ok) {
      let signOutcode = `<a class="sign-in" id="sign-in"><i class="fa fa-sign-out"></i>Sign Out</a>`
      console.log(document.getElementById("sign-i"))
      document.getElementById("sign-in").innerHTML = signOutcode
      accountLink.style.display = 'block'
      createPostButton.style.display = 'block'
      //commentSesh.style.display = 'block'
      console.log('User session is valid');
    } else {
      console.log('User session is invalid or expired');
    }
  })
  .catch(error => {
    // Handle any network errors
    console.error('Error occurred while checking user session:', error);
    // Show error message to the user or retry the request
  });
}


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


function showComment(){
    var commentArea = document.getElementById("comment-area");
    if (commentArea.classList.contains("hide")) {
        commentArea.classList.remove("hide");
    } else {
        commentArea.classList.add("hide");
    }
    
}
