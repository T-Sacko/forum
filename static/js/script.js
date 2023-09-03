
// get liked posts data
async function checkSession(){
 const resp = await fetch('/session')
 return await resp.json()
}
let userStatus = false; // default value

checkSession().then(data => {
    userStatus = data.status;
    console.log("User logged in:", userStatus);
});


// fetch("/get-post-likes", {
//   method: "GET"
// })
//   .then(response => response.json())
//   .then(data => {
//     // Handle the like data
//     console.log(data, 'this is the data');

//     for (const likeData of data) {
//       const postId = likeData.postId;
//       const value = likeData.value;
//       const likeButton = document.getElementById(`${postId}-like`);
//       if (value == 1) {
//         if (likeButton) {
//           toggleLiked(likeButton)
//           console.log("toggleLiked it up still")
//         }
//       } else {
//         const dislikeButton = document.getElementById(`${postId}-dislike`);
//         toggleDisliked(dislikeButton)

//       }

//     }
//   })
//   .catch(error => {
//     console.error("Error with the posts toggleLiked data is:", error);
//     // Handle the error
//   });

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

// document.getElementById('dropdownToggle').addEventListener('click', function() {
//   var menu = document.querySelector('.dropdown-menu');

//   if (menu.style.opacity === "0" || menu.style.opacity === "") {
//     menu.style.opacity = "1";
//     menu.style.visibility = "visible";
//   } else {Fsi
//     menu.style.opacity = "0";
//     menu.style.visibility = "hidden";
//   }
// });



const likes = document.querySelectorAll('.likes')

likes.forEach(likeButton => {
  const likeStr = likeButton.getAttribute('id');
  const postId = likeStr.split('-')[0]; // This will split the likeId string at the '-' character
  const dislikeButton = document.getElementById(`${postId}-dislike`)


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

//////////////////////////////////////////-commentRetrieval
function createCommentElement(comment) {
  const commentDiv = document.createElement('div');
  commentDiv.classList.add('comment');

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

const posts = document.querySelectorAll('.post')


posts.forEach(post => {
  let commentsFetched = false;
  const commentBtn = post.querySelector('.comment-btn')
  commentBtn.addEventListener('click', async () => {
    const commentSection = post.querySelector('.comments')
    commentSection.innerHTML=''
    if (getComputedStyle(commentSection).display !== 'none') {
      commentSection.style.display = 'none';
      return; // Exit the function if the comment section is being hidden
    } else {
      commentSection.style.display = 'block';
     
    }

    if (commentsFetched) ;
    const postID = post.id.split('-')[1]
    const res = await fetch(`/get-comments?postID=${postID}`)
    const data = await res.json()
    if (data == null){
      return
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

    })
    commentsFetched = true; 
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
fetch(`/like-comment?id=${commentID}&action=${action}`)
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


document.addEventListener("DOMContentLoaded", function() {
  const commentInputs = document.querySelectorAll('.comment-input');

  commentInputs.forEach(input => {
    input.addEventListener('input', function() {
        this.style.height = 'auto';           // Reset height
        this.style.height = (this.scrollHeight) + 'px';  // Set to scrollHeight
    });
  });

  const commentForms = document.querySelectorAll(".comment-form")

  commentForms.forEach(commentForm => {
    
    console.log("form in a it")
          commentForm.addEventListener('submit', (e) => {
      e.preventDefault();
      let formData = new FormData(commentForm);

      const textarea = commentForm.querySelector('.comment-input');
      const commentErr = commentForm.querySelector('.comment-err')
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
          const id = comment.id
          const commentsDiv = document.getElementById(`comments-${id}`)
          if (!commentsDiv) {
            console.log("what are u thinking mate")
          }
          const commentDiv = document.createElement('div')
          commentDiv.className = 'comment'

          commentDiv.innerHTML = `
            <h4 class="username"><i class="fa fa-user"></i> ${comment.username}</h4>
            <p class="content">${comment.content}</p>
            <div class="comment-footer">
                <div class="comment-like-buttons">
                    <span id="comment${comment.id}-like" class="comment-like-button comment-button">
                        <i class="comment-icon fa fa-thumbs-o-up"></i> ${comment.likes}
                    </span>
                    <span id="comment${comment.id}-dislike" class="comment-dislike-button comment-button">
                        <i class="comment-icon fa fa-thumbs-o-down"></i> ${comment.dislikes}
                    </span>
                </div>
            </div>
          `;

          commentsDiv.prepend(commentDiv)

          // const usernameH4 = document.createElement('h4')
          // usernameH4.className = 'username';
          // usernameH4.innerHTML= `<i class=" fa fa-user"> ${comment.username}`

          // const contentp = document.createElement('p')
          // contentp.className='content'
          // contentp.textContent= comment.content
          
          // const footerDiv = document.createElement('div')
          // footerDiv.className = 'comment-footer'

          // const likeButtonsDiv = document.createElement('div')
          // likeButtonsDiv.className = 'comment-like-buttons'

          // const likeSpan = document.createElement('span')
          // likeSpan.id = `${comment.id}-like`
          // likeSpan.className = 'comment-like-button comment-button'
          // likeSpan.innerHTML = `<i class = `


      })
      .catch(error => {
          console.log("Error:", error);
      });
    })

  })
})

if (!userStatus){
  console.log(userStatus, "wagwan wha u know bout debugging man, its ber long una, truss me, long like long like...")

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

}