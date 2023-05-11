'use strict';

document.addEventListener('DOMContentLoaded', () => {
  const postForm = document.getElementById('postForm');
  const postsContainer = document.getElementById('postsContainer');
  const errorMessage = document.getElementById('errorMessage'); // Get the error message element

  // Fetch all posts from the backend and display them on page load
  fetch('/posts')
    .then(response => response.json())
    .then(posts => {
      posts.forEach(post => {
        const postElement = createPostElement(post);
        postsContainer.appendChild(postElement);
      });
    });

  // Handle form submission
  const signupForm = document.querySelector('.signup form');
  postForm.addEventListener('submit', event => {
    event.preventDefault();
    const username = signupForm.querySelector('input[name="username"]').value;
    const email = signupForm.querySelector('input[name="email"]').value;
    const password = signupForm.querySelector('input[name="password"]').value;
    const postContent = document.getElementById('postContent').value;
    const postImage = document.getElementById('postImage').files[0];

    // Check username availability before submitting the form
    isUsernameAvailable(username)
      .then(isAvailable => {
        if (isAvailable) {
          // Proceed with form submission
          const formData = new FormData();
          formData.append('username', username);
          formData.append('email', email);
          formData.append('password', password);
          formData.append('content', postContent);
          formData.append('image', postImage);

          fetch('/create-post', {
            method: 'POST',
            body: formData
          })
            .then(response => response.json())
            .then(post => {
              const postElement = createPostElement(post);
              postsContainer.insertBefore(postElement, postsContainer.firstChild);
            });
        } else {
          // Display an error message indicating that the username is taken
          errorMessage.textContent = 'Username is already taken.';
          errorMessage.style.display = 'block';
        }
      })
      .catch(error => {
        console.error('Error checking username availability:', error);
      });

    postForm.reset();
  });

  // Function to check username availability
  function isUsernameAvailable(username) {
    return fetch(`/check-username?username=${username}`)
      .then(response => response.json())
      .then(response => response.available)
      .catch(error => {
        console.error('Error checking username availability:', error);
        return false;
      });
  }
  
  // Function to create a post element
  function createPostElement(post) {
    const postElement = document.createElement('div');
    postElement.classList.add('post');

    const contentElement = document.createElement('p');
    contentElement.textContent = post.content;
    postElement.appendChild(contentElement);

    if (post.image) {
      const imageElement = document.createElement('img');
      imageElement.src = post.image;
      postElement.appendChild(imageElement);
    }

    const likeButton = document.createElement('button');
    likeButton.textContent = 'Like';
    likeButton.addEventListener('click', () => {
      // Send request to like the post
      fetch(`/like/${post.id}`, { method: 'POST' })
        .then(response => response.json())
        .then(updatedPost => {
          // Update the post with the new like count
          const likesElement = postElement.querySelector('.likes');
          likesElement.textContent = `Likes: ${updatedPost.likes}`;
        });
    });
    postElement.appendChild(likeButton);

    const likesElement = document.createElement('span');
    likesElement.classList.add('likes');
    likesElement.textContent = `Likes: ${post.likes}`;
    postElement.appendChild(likesElement);

    return postElement;
  }
});
