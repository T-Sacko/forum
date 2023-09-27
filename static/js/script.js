
// get liked posts data
async function checkSession(){
 const resp = await fetch('/session')
 return await resp.json()
}
let userStatus = false; // default value

checkSession().then(data => {
  userStatus = data.status;
  console.log("User logged in:", userStatus);
  if (!userStatus) {
    const overlay = document.getElementById('overlay')
    const signIn = document.getElementById('sign-in')
  
    const loginButton = document.getElementById('loginButton')
  
    const signupButton = document.getElementById('signupButton')
  
    const signupForm = document.getElementById('signupForm')
  
    const loginForm = document.getElementById('loginForm')
  
  
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
    
    
} else   {
  var createPostButton = document.getElementById("create-post-button");




  var   modalContainer = document.getElementById("modal-container");
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

  // const postContent = postForm.querySelector('.form-text')
  // postContent.addEventListener('input', () => {
  //   this.style.height = 'auto'
  //   this.style.height = this.scrollHeight + 'px'
  // })




  signOutBtn = document.querySelector('#sign-out-btn')
  console.log(signOutBtn)
  signOutBtn.addEventListener('click', () => {
    // Clear local session first
    let cookieValue = getCookie("session");
    clearLocalAuthToken();
  
    // Check for internet connectivity
    if (navigator.onLine) {
      // If online, send logout request
      fetch(`/signout?seshID=${cookieValue}`, {
          method: 'POST',
          // ... other necessary configurations
      })
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }else{
          window.location.href = '/';
        }
          
      })
      .catch(error => {
        console.log('Failed to log out from server:', error);
        alert('Logged out locally but there was an issue logging out from the server. You may need to logout again when connected to the internet.');
      });
    }else {
      alert('You are offline. You have been logged out locally but may need to logout again when connected to the internet.');
    }
  });
      
}
});
  
function clearLocalAuthToken() {
  document.cookie = "session=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
}

function getCookie(name) {
  const value = "; " + document.cookie;
  const parts = value.split("; " + name + "=");
  if (parts.length == 2) return parts.pop().split(";").shift();
}


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



function increment(button) {
  const likeCountElement = button.querySelector('.like-count')
  const likeCountText = likeCountElement.textContent;
  const likeCountValue = parseInt(likeCountText);
  const incrementedCount = likeCountValue + 1;
  likeCountElement.textContent = incrementedCount.toString();
}

function decrement(dislikeBtn) {
  const dislikeCountElement = dislikeBtn.querySelector('.like-count')
  const dislikeCountText = dislikeCountElement.textContent;
  const dislikeCountValue = parseInt(dislikeCountText);
  const decrementedCount = dislikeCountValue - 1;
  dislikeCountElement.textContent = decrementedCount.toString();
}


//////////////////////////////////////////-commentRetrieval
function createCommentElement(comment) {
  const commentDiv = document.createElement('div');
  commentDiv.classList.add('comment');
  commentDiv.setAttribute('id',`${comment.id}`)

  const usernameH4 = document.createElement('h4');
  usernameH4.classList.add('username');
  usernameH4.innerHTML = `<i class="fa fa-user"></i> ${comment.username}`;
  commentDiv.appendChild(usernameH4);

  const contentP = document.createElement('p');
  contentP.classList.add('content');
  contentP.textContent = comment.content;
  commentDiv.appendChild(contentP);

  const footerDiv = document.createElement('div');
  footerDiv.classList.add('comment-footer');

  const buttonsDiv = document.createElement('div');
  buttonsDiv.classList.add('comment-like-buttons');

  const likeSpan = document.createElement('span');
  likeSpan.id = `comment-${comment.id}-like`;
  likeSpan.classList.add('comment-like-button', 'comment-button');
  likeSpan.innerHTML = `<i class="comment-icon comment-like-icon fa fa-thumbs-o-up"></i><span class="comment-likes"> ${comment.likes}</span>`;
  buttonsDiv.appendChild(likeSpan);

  const dislikeSpan = document.createElement('span');
  dislikeSpan.id = `comment-${comment.id}-dislike`;
  dislikeSpan.classList.add('comment-dislike-button', 'comment-button');
  dislikeSpan.innerHTML = `<i class="comment-icon comment-dislike-icon fa fa-thumbs-o-down"></i><span class="comment-dislikes"> ${comment.dislikes}</span>`;
  buttonsDiv.appendChild(dislikeSpan);
  const dislikeIcon = dislikeSpan.querySelector('.comment-dislike-icon')
  const likeIcon = likeSpan.querySelector('.comment-like-icon')
  if (comment.user_like_status==-1){
    toggleDisliked(dislikeIcon)
  }else if (comment.user_like_status==1) {
    toggleLiked(likeIcon)
  }

  footerDiv.appendChild(buttonsDiv);
  commentDiv.appendChild(footerDiv);

  return commentDiv;
}

//----------------------------------------display-comment-section


function addCommentListeners(comment) {
  const commentLikeBtn = comment.querySelector(`#comment-${comment.id}-like`);
  console.log(comment, 'whaaa', comment.id)
  const commentDislikeBtn = comment.querySelector(`#comment-${comment.id}-dislike`);
  const dislikeIcon =  comment.querySelector('.comment-dislike-icon');
  const likeIcon = comment.querySelector('.comment-like-icon');
  const likeCount = comment.querySelector('.comment-likes');
  const dislikeCount = comment.querySelector('.comment-dislikes');
  
  commentLikeBtn.addEventListener('click', () => {
    handleLikeClick(comment, likeIcon, dislikeIcon, likeCount, dislikeCount);
  });


  commentDislikeBtn.addEventListener('click', () => {
    handleDislikeClick(comment, likeIcon, dislikeIcon, likeCount, dislikeCount);
  });
  
}

function handleLikeClick(comment, likeIcon, dislikeIcon, likeCount, dislikeCount) {
  toggle(likeIcon);
  if (likeIcon.classList.contains('liked')) {
      incComment(likeCount);
      if (dislikeIcon.classList.contains('disliked')) {
          toggle(dislikeIcon);
          decrComment(dislikeCount);
          likeReq(comment.id, 'removeDislike');
      }
      likeReq(comment.id, 'like');
  } else {
      decrComment(likeCount);
      likeReq(comment.id, 'removeLike');
  }
}

function handleDislikeClick(comment, likeIcon, dislikeIcon, likeCount, dislikeCount) {
  toggle(dislikeIcon);
  if (dislikeIcon.classList.contains('disliked')) {
      incComment(dislikeCount);
      if (likeIcon.classList.contains('liked')) {
          toggle(likeIcon);
          decrComment(likeCount);
          likeReq(comment.id, 'removeLike');
      }
      likeReq(comment.id, 'dislike');
  } else {
      decrComment(dislikeCount);
      likeReq(comment.id, 'removeDislike');
  }
}

const posts = document.querySelectorAll('.post')

function setUpLike(likeButton, dislikeButton, postId) {
  const likeIcon = likeButton.querySelector('.likes')
  const dislikeIcon = dislikeButton.querySelector('.dislikes')
  likeButton.addEventListener('click', () => {
    // toggle like button 
    toggleLiked(likeIcon)

    if (likeIcon.classList.contains('liked')) {
      // send like req to server
      console.log("we liking suttin")
      increment(likeButton)
      if (dislikeIcon.classList.contains('disliked')) {
        toggleDisliked(dislikeIcon)
        decrement(dislikeButton)
        // send req to remove dislike
        handleLikeAction(postId, "removeDislike")
      }
      handleLikeAction(postId, "like")
    } else {
      //send unlike req
      console.log("we unliking suttin")
      handleLikeAction(postId, "unlike")
      decrement(likeButton)
    }



  })

  dislikeButton.addEventListener('click', () => {
    // toggle like button 
    toggleDisliked(dislikeIcon)


    if (dislikeIcon.classList.contains('disliked')) {
      // send dislike req to server
      console.log("we disliking suttin")
      increment(dislikeButton)
      if (likeIcon.classList.contains('liked')) {
        toggleLiked(likeIcon)
        decrement(likeButton)
        // send req to remove like
        handleLikeAction(postId, "unlike")
      }
      handleLikeAction(postId, "dislike")
    } else {
      //send remove dislike req
      decrement(dislikeButton)
      console.log("we removing dislike")
      handleLikeAction(postId, "removeDislike")
    }


  })
}

function setContent(content, seeMore) {
  if (content.scrollHeight > content.clientHeight) {
    seeMore.style.display = 'inline'
  }else{
    return
  }

  var isMore = false
  seeMore.addEventListener('click', () => {
    if (isMore) {
      content.style.maxHeight = '197px'
      seeMore.textContent = 'see more...'
      isMore = false
      return
    }
    console.log(content)
    content.style.maxHeight = 'none';
    seeMore.textContent = 'see less'
    isMore = true
  })

}

posts.forEach(post => {
  const seeMore = post.querySelector('.see-more')
  const content = post.querySelector('.post-content')
  const postId = post.id.split('-')[1];
  const commentBtn = post.querySelector('.comment-btn')
  const likebtn = post.querySelector('.post-l')
  const dislikeBtn = post.querySelector('.post-d')
  setUpLike(likebtn,dislikeBtn,postId)
  const commentForm = post.querySelector('.comment-form')
  const commentsDiv = post.querySelector('.comments')

  setContent(content, seeMore)


  commentForm.addEventListener('submit', (e) => {
    e.preventDefault();
    let formData = new FormData(commentForm);

    const textarea = commentForm.querySelector('.comment-input');
    // const commentErr = commentForm.querySelector('.comment-err')
    if (textarea.value==''){
      // commentErr.textContent='input required'
      // setTimeout()
      return
    }
   textarea.value = '';


    fetch('/comment', {
      method: 'POST',
      body:   formData
    })
    .then(response => response.json()) 
    .then(comment => {
      console.log(comment);
      // const id = comment.id
      // if (!commentsDiv) {
      //   console.log("what are u thinking mate")
      //  }
      const commentDiv = createCommentElement(comment)
      addCommentListeners(commentDiv)

      commentsDiv.prepend(commentDiv)
      commentsDiv.style.display = 'block'

        


    })
    .catch(error => {
        console.log("Error:", error);
    });
  })


  
  commentBtn.addEventListener('click', async () => {
    const commentSection = post.querySelector('.comments')
    commentSection.innerHTML=''
   
    
    
    const postID = post.id.split('-')[1]
    const res = await fetch(`/get-comments?postID=${postID}`)
    const data = await res.json()
    if (data == null){
      console.log('issue here')
      return
    }
    if (getComputedStyle(commentSection).display !== 'none') {
      commentSection.style.display = 'none';
      return; // Exit the function if the comment section is being hidden
    } else {
      commentSection.style.display = 'block';
     
    }
    data.forEach(comment => {
      console.log(comment)
      const commentElement = createCommentElement(comment);
      commentSection.appendChild(commentElement);
      const commentLikeBtn = commentSection.querySelector(`#comment-${comment.id}-like`)
      const commentDislikeBtn = commentSection.querySelector(`#comment-${comment.id}-dislike`)
      const dislikeIcon =  commentDislikeBtn.querySelector('.comment-dislike-icon')
      const likeIcon = commentLikeBtn.querySelector('.comment-like-icon')
      const likeCount = commentLikeBtn.querySelector('.comment-likes')
      const dislikeCount = commentDislikeBtn.querySelector('.comment-dislikes')
      
      let yo = false
      checkSession().then(data => {
        
        yo = data.status
        console.log('yo =', yo)
        if (yo){
          commentLikeBtn.addEventListener('click', () => {
            toggle(likeIcon)
            if (likeIcon.classList.contains('liked')) {
              incComment(likeCount)
              if (dislikeIcon.classList.contains('disliked')){
                toggle(dislikeIcon)
                decrComment(dislikeCount)
                likeReq(comment.id,'removeDislike')
              }
              likeReq(comment.id,'like')
            }else{
              decrComment(likeCount)
              likeReq(comment.id,'removeLike')
            }
            
          })
  
          commentDislikeBtn.addEventListener('click', () => {
            toggle(dislikeIcon)
            if (dislikeIcon.classList.contains('disliked')) {
              incComment(dislikeCount)
              if (likeIcon.classList.contains('liked')) {
                likeReq(comment.id,'removeLike')
                toggle(likeIcon)
                decrComment(likeCount)
              }
              likeReq(comment.id,'dislike')
            }else{
              decrComment(dislikeCount)
              likeReq(comment.id,'removeDislike')
            }
        
          
          })
        }
      })
    })
    const yOffset = -24; // 1.5rem = 24px, assuming the user's default font-size is 16px
    const y = post.getBoundingClientRect().top + window.pageYOffset + yOffset;

    window.scrollTo({ top: y, behavior: 'smooth' });
  })
})

function toggle(element) {
    if (element.classList.contains('comment-like-icon')) {
        element.classList.toggle('liked');
        element.classList.toggle('fa-thumbs-o-up');
        element.classList.toggle('fa-thumbs-up');
    } else if (element.classList.contains('comment-dislike-icon')) {
        element.classList.toggle('disliked');
        element.classList.toggle('fa-thumbs-o-down');
        element.classList.toggle('fa-thumbs-down');
    }
}


function incComment(element) {
  let Count = parseInt(element.textContent.trim(), 10) + 1;

  element.textContent = ' ' + Count;

}

function decrComment(element) {
  let Count = parseInt(element.textContent.trim(), 10) - 1;
  element.textContent = ' ' + Count;

}



const likeReq = (commentID,action) => {
  console.log(action,'action is previous word')
  fetch(`/like-comment?id=${commentID}&action=${action}`, {method: 'POST'})
}

function toggleCount(element) {
  let currentCount = parseInt(element.textContent.trim(), 10); // Get the current count (ignoring spaces)

  if (element.classList.contains('toggled')) {
      currentCount -= 1; // If it's toggled, decrease the count
  } else {
      currentCount += 1; // If it's not toggled, increase the count
  }
  
  element.textContent = ' ' + currentCount; // Update the text content with the new count (with a space before the number)
  element.classList.toggle('toggled'); // Toggle the class to keep track of the current state
}
//--------------------------------------------------------



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
            errorElement.textContent = "Username is taken";
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
            errorElement.textContent = "Email Already In Use";
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

async function fetchData(url, options = {}){
  const resp = await fetch(url, options)

  if (!resp.ok) {
    throw new Error(`HTTP error! Status: ${resp.status}`);
  }
  const data = await resp.json()
  return data
  
}



document.getElementById('loginForm').addEventListener('submit', async (event) => {
  event.preventDefault(); // Prevent the form from submitting normally

  const formData = new FormData(event.target); // Get form data
  const email = formData.get('email');
  const password = formData.get('password');

  // Create the payload to send in the AJAX request
  const payload = {
    email: email,
    password: password
  };

  try {
    // Make the AJAX request to the server
    const response = await fetch('/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    });

    // Handle the response from the server
    if (response.ok) {
      // If the login is successful, redirect the user to a new page or perform other actions
      console.log('Login successful!');
      window.location.href = '/'; // Redirect to the dashboard page after successful login
    } else {
      // If the login fails, display an error message
      console.log('Login failed:555555555');
      const errorElement = document.querySelector('.login-email-error');
      errorElement.setAttribute('style', 'display:block;');
      errorElement.textContent = 'Error: Invalid email or password';
    }
  } catch (error) {
    console.error('Error:', error);
  }
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


  const commentInputs = document.querySelectorAll('.comment-input');

  commentInputs.forEach(input => {
    input.addEventListener('input', function() {
        this.style.height = 'auto';           // Reset height
        this.style.height = (this.scrollHeight) + 'px';  // Set to scrollHeight
    });
  });

 
 