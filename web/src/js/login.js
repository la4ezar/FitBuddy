document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.getElementById('login-form');

    loginForm.addEventListener('submit', function (event) {
        event.preventDefault();

        const email = document.getElementById('login-email').value;
        const password = document.getElementById('login-password').value;

        // Replace 'your-graphql-endpoint' with your actual GraphQL endpoint
        const graphqlEndpoint = 'http://localhost:8080/graphql';

        const gqlMutation = `
            mutation {
                loginUser(email: "${email}", password: "${password}") {
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
            body: JSON.stringify({query: gqlMutation}),
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
                    alert('Login failed. Please try again.');
                } else {
                    alert('Login successful!');
                    // Redirect to the login page
                    window.location.href = './dashboard.html'; // Update with the actual path
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
                alert(`An error occurred. ${error.message}`);
            });

        document.cookie = `email=${email}; SameSite=Lax; path=/;`;
    });
});