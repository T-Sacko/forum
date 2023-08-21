const signupLink = document.getElementById('sign-in');
const registerLink = document.getElementById('register_btn')
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
const SignOut = document.getElementById("sign_out")
const commentBtns = document.getElementsByClassName("btns")
const commentContainer = document.getElementsByClassName("comment")
const errorContainer = document.getElementById('err');

function showComments(commentId) {

    commentId = commentId.replace("comment-btn", "comment-section")
    let commentSection = document.getElementById(commentId)
    commentSection.innerHTML = '';
    let postId = commentId.replace("comment-section-", "")
    if (commentSection.style.display === "none" || commentSection.style.display === "") {
        commentSection.style.display = "block"
    } else {
        commentSection.style.display = "none"
        return
    }

    fetch (`/get-comments?postID=${postId}`, {
        method: 'GET'
    }).then(response => response.json())
    .then(data => {
    for (const comment of data) {
        let CommentDiv = document.createElement('div');
        CommentDiv.id = comment.id;
        CommentDiv.className = 'comments';
    
        let h4 = document.createElement("h4");
        h4.textContent = comment.username;
        CommentDiv.appendChild(h4);
    
        let p = document.createElement("p");
        p.textContent = comment.comment;
        CommentDiv.appendChild(p);
    
        // Create the likes container for each comment
        let likesContainer = document.createElement('div');
        likesContainer.className = 'likes-container';
        likesContainer.id = `${comment.id}-likes`; // Set the likes container id
    
        // Create the like button
        let likeButton = document.createElement('button');
        likeButton.className = 'like-button';
        likeButton.id = `${comment.id}-like`; // Set the like button id
        likeButton.setAttribute('onclick', 'likeSender(this.id)');
        likeButton.innerHTML = '<i class="fa fa-thumbs-o-up"></i>';
    
        // Create the like count span
        let likeCountSpan = document.createElement('span');
        likeCountSpan.className = 'like-count';
        likeCountSpan.id = `${comment.id}-1`; // Set the like count span id
        likeCountSpan.textContent = '0'; // Initialize with 0 likes
    
        likeButton.appendChild(likeCountSpan);
    
        // Create the dislike button (similar to the like button)
        let dislikeButton = document.createElement('button');
        dislikeButton.className = 'dislike-button';
        dislikeButton.id = `${comment.id}-dislike`; // Set the dislike button id
        dislikeButton.setAttribute('onclick', 'likeSender(this.id)');
        dislikeButton.innerHTML = '<i class="fa fa-thumbs-o-down"></i>';
    
        let dislikeCountSpan = document.createElement('span');
        dislikeCountSpan.className = 'dislike-count';
        dislikeCountSpan.id = `${comment.id}-2`; // Set the dislike count span id
        dislikeCountSpan.textContent = '0'; // Initialize with 0 dislikes
    
        dislikeButton.appendChild(dislikeCountSpan);
    
        // Append like and dislike buttons to the likes container
        likesContainer.appendChild(likeButton);
        likesContainer.appendChild(dislikeButton);
    
        // Append the likes container to the CommentDiv
        CommentDiv.appendChild(likesContainer);
    
        // Append the entire CommentDiv to the commentSection
        commentSection.appendChild(CommentDiv);
    }
    checkSession()
    .then(loggedIn => {
        if (loggedIn) {
            getLikedPosts(true, 0)
        } else{
            getLikedPosts(true, postId)
        }
    })
    console.log("This is the comment likes")
    })
};

function closePopup() {
    popupContainer.style.display = "none";
    modalContainer.style.display = "none";
    errorContainer.style.display = "none";
}

    signupLink.addEventListener('click', function() {
        if (signupLink.innerText === 'Sign Out') {
            const close = document.getElementById('close')
            popupContainer.style.display = 'flex'
            popupDiv.style.height = 'auto'
            close.style.display = 'none'
            SignOut.style.display = 'block'
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
    
    registerLink.addEventListener('click', function() {
        login.style.display = 'none'
        signup.style.display = 'block'
    });
    
    loginLink.addEventListener('click', function() {
        signup.style.display = 'none'
        login.style.display = 'block'
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


async function checkSession() {
    console.log('checkSession function being called');
    return fetch('/check-session', {
    method: 'POST',
      credentials: 'include' // Include cookies for session tracking
    })
    .then(response => {
    if (response.ok) {
        return true;
    } else {
        return false;
    }
    })
}

function signIn(){
    checkSession()
    .then(loggedIn => {
        if (loggedIn) {
            let signOutcode = `<a class="sign-in" id="sign-in"><i class="fa fa-sign-out"></i>Sign Out</a>`;
            console.log(document.getElementById("sign-in"));
            document.getElementById("sign-in").innerHTML = signOutcode;
            accountLink.style.display = 'block';
            createPostButton.style.display = 'block';
            for (let i = 0; i < commentContainer.length; i++) {
                commentContainer[i].style.display = "block";
            }
            getLikedPosts(false);
        } else {
            console.log("There has been an error in signing in!")
        }
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

function comment(postId) {
    const comment = document.getElementById(postId + "-comment").value
    if (comment == '') {
        return
    }

    const requestBody = {
        postId: postId,
        comment: comment
    };

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

function toggleButton(Button) {
    console.log(Button.id)
    if (Button.id.includes("dislike")){
        let icon = Button.querySelector('i');
        Button.classList.toggle('active')
        icon.classList.toggle('fa-thumbs-down');
    } else if (Button.id.includes("like")) {
        let icon = Button.querySelector('i');
        Button.classList.toggle('active')
        icon.classList.toggle('fa-thumbs-up');
    }
}

function handleLikeAction(postId, commentId, action) {
    fetch("/like-post", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ postId: postId, commentId: commentId, action: action })
    });
}

function increment(postId, id) {
    const dislikeCountElement = document.getElementById(postId + "-" + id);
    const dislikeCountText = dislikeCountElement.textContent;
    console.log(dislikeCountText)
    const dislikeCountValue = parseInt(dislikeCountText);
    const incrementedCount = dislikeCountValue + 1;
    dislikeCountElement.textContent = incrementedCount.toString();
}

function decrement(postId, id) {
    const dislikeCountElement = document.getElementById(postId + "-" + id);
    const dislikeCountText = dislikeCountElement.textContent;
    console.log(dislikeCountText)
    const dislikeCountValue = parseInt(dislikeCountText);
    const decrementedCount = dislikeCountValue - 1;
    dislikeCountElement.textContent = decrementedCount.toString();
}


function likeSender(likeID) {
    let match = likeID.match(/\d+/);
    let button = document.getElementById(likeID)
    let ID = match[0]
    let countIDNumber = button.querySelector('span').id.slice(-1);
    checkSession()
    .then(loggedIn => {
        if (loggedIn) {
            if (button.parentNode.parentNode.classList == "interaction-container") {
                if (button.classList.contains("active")) {
                    toggleButton(button)
                    decrement(ID, countIDNumber)
                    if (countIDNumber == 2) {
                        handleLikeAction(ID, "0", "removeDislike")
                        console.log("this is the id for removing dislike", ID)
                    } else {
                        handleLikeAction(ID, "0", "unlike")
                        console.log("this is the id for removing like", ID)
                    }
                } else {
                    increment(ID, countIDNumber)
                    if (countIDNumber == 2) {
                        let oppositeButton = document.getElementById(ID + "-like")
                        if (oppositeButton.classList.contains("active")) {
                            handleLikeAction(ID, "0", "unlike")
                            toggleButton(oppositeButton)
                            decrement(ID, "1")
            } 
                handleLikeAction(ID, "0", "dislike")
                toggleButton(button)
                console.log("this is the id for disliking", ID)
                    } else {
                        let oppositeButton = document.getElementById(ID + "-dislike")
                        if (oppositeButton.classList.contains("active")) {
                            handleLikeAction(ID, "0", "removeDislike")
                            toggleButton(oppositeButton)
                            decrement(ID, "2")
                        } 
                handleLikeAction(ID, "0", "like")
                toggleButton(button)
                console.log("this is the id for liking", ID)
                    }
                } 
            } else if (button.parentNode.parentNode.classList == "comments"){
                let postId = button.parentNode.parentNode.parentNode.id.match(/\d+/)[0]
                if (button.classList.contains("active")) {
                    toggleButton(button)
                    decrement(ID, countIDNumber)
                    if (countIDNumber == 2) {
                        handleLikeAction(postId, ID, "removeDislike")
                        console.log("this is the id for removing dislike", ID)
                    } else {
                        handleLikeAction(postId, ID, "unlike")
                        console.log("this is the id for removing like", ID)
                    }
                } else {
                    increment(ID, countIDNumber)
                    if (countIDNumber == 2) {
                    let oppositeButton = document.getElementById(ID + "-like")
                        if (oppositeButton.classList.contains("active")) {
                    handleLikeAction(postId, ID, "unlike")
                    toggleButton(oppositeButton)
                    decrement(ID, "1")
            } 
                handleLikeAction(postId, ID, "dislike")
                toggleButton(button)
                console.log("this is the id for disliking", ID)
                    } else {
                        let oppositeButton = document.getElementById(ID + "-dislike")
                        if (oppositeButton.classList.contains("active")) {
                        handleLikeAction(postId, ID, "removeDislike")
                        toggleButton(oppositeButton)
                    decrement(ID, "2")
            } 
            handleLikeAction(postId, ID, "like")
            toggleButton(button)
            console.log("this is the id for liking", ID)
                    }
        }
    }

    } else {
        errorContainer.style.display = 'flex';
        return;
    }
    });
}

function getLikedPosts(CommentLikes, postId) {
    if (CommentLikes) {
        fetch (`/get-post-likes?postID=${postId}`, {
            method: "GET"
        })
        .then(response => response.json())
        .then(data => {
        let Likes = data.likes
        let Dislikes = data.dislikes
        for (let like of Likes) {
            let button = document.getElementById(like.commentId + "-like")
            let amountOfLikes = document.getElementById(like.commentId + "-1")
            if (like.commentId === 0) {
                continue
            }
            amountOfLikes.textContent = like.value
            if (postId === 0) {
                toggleButton(button)
            }
        }

        for (let dislike of Dislikes) {
            let button = document.getElementById(dislike.commentId + "-dislike")
            let amountOfDislikes = document.getElementById(dislike.commentId + "-2")
            if (dislike.commentId === 0) {
                continue
            }
            amountOfDislikes.textContent = Math.abs(dislike.value)
            if (postId === 0) {
                toggleButton(button)
            }
        }
        })
        .catch(error => {
        console.error("Error with the posts toggleLiked data is:", error);
        });

    } else {
    fetch("/get-post-likes", {
        method: "GET"
    })
        .then(response => response.json())
        .then(data => {
        let Likes = data.likes
        let Dislikes = data.dislikes
        console.log(data, 'this is the data');
        for (let like of Likes) {
            let button = document.getElementById(like.postId + "-like")
            toggleButton(button)
        }

        for (let dislike of Dislikes) {
            let button = document.getElementById(dislike.postId + "-dislike")
            toggleButton(button)
        }
        })
        .catch(error => {
            console.error("Error with the posts toggleLiked data is:", error);
        // Handle the error
        });
    }
}

async function  deleteCookie() {
    console.log("deleteCookie function being called")
    return fetch('/del-cookie', {
        method: 'POST',
    }).then(response=> {
        if (response.ok) {
            return true
        } else {
            console.log("failed to execute deletion")
            return false
    }})
    .catch(error => {
        console.error('Error occurred while deleting user session:', error);
        return false
    });
};      

function signOut() {
    console.log("Signing Out")
    deleteCookie()
    .then(cookieDeleted => {
        if (cookieDeleted) {
            let signIncode = `<a class="sign-in" id="sign-in">Sign In<i class="fa fa-sign-in" id="sign-i"></i></a>`
            let close = document.getElementById("close")
            document.getElementById("sign-in").innerHTML = signIncode
            accountLink.style.display = 'none'
            createPostButton.style.display = 'none'
            popupContainer.style.display = 'none';
            SignOut.style.display = "none"
            close.style.display = "block"
            for (let i = 0; i < commentContainer.length; i++) {
                commentContainer[i].style.display = "block";
            }
            location.reload()
        } else {
            console.log("Could not delete cookie")
        }
    })
}
