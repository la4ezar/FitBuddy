document.addEventListener('DOMContentLoaded', function () {
    const forumContainer = document.querySelector('.post-display');

    fetchAllPosts();

    const postForm = document.getElementById('post-form');
    postForm.addEventListener('submit', function (event) {
        event.preventDefault();

        const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
        if (emailCookie) {
            console.log('Email:', emailCookie);
        } else {
            console.log('Email cookie not found.');
        }
        const title = document.getElementById('post-title').value;
        const content = document.getElementById('post-content').value;

        const graphqlEndpoint = 'http://localhost:8080/graphql';

        const gqlMutation = `
            mutation {
                createPost(email: "${emailCookie}", title: "${title}", content: "${content}") {
                    ID
                    UserEmail
                    Title
                    Content
                    CreatedAt
                }
            }
        `;

        // Make the GraphQL request to create a new post
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
        // Replace 'your-graphql-endpoint' with your actual GraphQL endpoint
        const graphqlEndpoint = 'http://localhost:8080/graphql';

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

        // Make the GraphQL request to fetch all posts
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

        // Clear the existing posts
        postsListContainer.innerHTML = '';

        // Check if the 'posts' array is defined and not empty before iterating
        if (Array.isArray(posts) && posts.length > 0) {
            // Display each post
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

                postsListContainer.appendChild(postElement);
            });
        } else {
            // If there are no posts, display a message
            const noPostsMessage = document.createElement('p');
            noPostsMessage.textContent = 'No posts available.';
            postsListContainer.appendChild(noPostsMessage);
        }
    }
});
