document.addEventListener('DOMContentLoaded', function () {
    const logoutButton = document.querySelector('.log-out-button');

    logoutButton.addEventListener('click', function () {
        // Replace 'your-graphql-endpoint' with your actual GraphQL endpoint
        const graphqlEndpoint = 'http://localhost:8080/graphql';

        const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];

        if (emailCookie) {
            console.log('Email:', emailCookie);
        } else {
            console.log('Email cookie not found.');
        }

        // Replace 'your-logout-mutation' with your actual GraphQL mutation for logout
        const gqlMutation = `
            mutation {
                logoutUser(email: "${emailCookie}") {
                    ID
                }
            }
        `;

        // Make the GraphQL request
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
                console.log('GraphQL Response:', data);
                // Handle the response data
                // For example, you can check if the registration was successful
                if (data.errors) {
                    alert('Registration failed. Please try again.');
                } else {
                    alert('Registration successful!');
                    document.cookie = `email=; SameSite=Lax; path=/;`;
                    // Redirect to the login page
                    window.location.href = './login.html'; // Update with the actual path
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
                alert(`An error occurred. ${error.message}`);
            });
    });
});