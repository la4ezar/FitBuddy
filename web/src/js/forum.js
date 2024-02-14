document.addEventListener('DOMContentLoaded', function () {
    const email = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
    if (!email) {
        return
    }
    const forumContainer = document.querySelector('.post-display');

    const graphqlEndpoint = 'http://localhost:8080/graphql';

    fetchAllPosts();

    const postForm = document.getElementById('post-form');
    postForm.addEventListener('submit', function (event) {
        event.preventDefault();

        const title = document.getElementById('post-title').value;
        const content = document.getElementById('post-content').value;

        const gqlMutation = `
            mutation {
                createPost(email: "${email}", title: "${title}", content: "${content}") {
                    ID
                    UserEmail
                    Title
                    Content
                    CreatedAt
                }
            }
        `;

        fetch(graphqlEndpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ query: gqlMutation }),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                if (data.errors) {
                    alert('Creating workout failed. Please try again.');
                } else {
                    fetchAllPosts();
                }
                console.log(data)
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
                alert(`An error occurred. ${error.message}`);
            });
    });

    function fetchAllPosts() {
        const gqlQuery = `
            query {
                getAllPosts {
                    ID
                    UserEmail
                    Title
                    Content
                    CreatedAt
                }
            }
        `;

        fetch(graphqlEndpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ query: gqlQuery }),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                if (data.errors) {
                    alert('Getting all posts failed. Please try again.');
                } else {
                    displayPosts(data.data.getAllPosts);
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
                alert(`An error occurred. ${error.message}`);
            });
    }

    function displayPosts(posts) {
        const postsListContainer = document.querySelector('.posts-list');

        postsListContainer.innerHTML = '';

        if (Array.isArray(posts) && posts.length > 0) {
            posts.forEach(post => {
                const postElement = document.createElement('div');
                postElement.className = 'post';

                const titleElement = document.createElement('h2');
                titleElement.textContent = post.Title;

                const contentElement = document.createElement('p');
                contentElement.textContent = post.Content;

                const userElement = document.createElement('p');
                userElement.textContent = `By ${post.UserEmail} on ${new Date(post.CreatedAt).toLocaleString()}`;

                postElement.appendChild(titleElement);
                postElement.appendChild(contentElement);
                postElement.appendChild(userElement);

                if (email === post.UserEmail) {
                    const deleteButton = document.createElement('button');
                    deleteButton.textContent = 'X';
                    deleteButton.className = 'delete-post-button';
                    postElement.appendChild(deleteButton);

                    deleteButton.addEventListener('click', function (event) {
                       deletePost(post.ID);
                    });
                }

                postsListContainer.appendChild(postElement);
            });
        } else {
            const noPostsMessage = document.createElement('p');
            noPostsMessage.textContent = 'No posts available.';
            postsListContainer.appendChild(noPostsMessage);
        }
    }

    function deletePost(postID) {
        const gqlMutation = `
            mutation {
                deletePost(postID: "${postID}")
            }
        `;

        fetch(graphqlEndpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ query: gqlMutation }),
        })
            .then(response => response.json())
            .then(data => {
                console.log(data)
                if (data.data) {
                    fetchAllPosts()
                } else {
                    console.error('Failed to delete post.');
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
            });
    }
});
