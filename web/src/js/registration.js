document.addEventListener('DOMContentLoaded', function () {
    const registrationForm = document.getElementById('registration-form');

    registrationForm.addEventListener('submit', function (event) {
        event.preventDefault();

        const email = document.getElementById('reg-email').value;
        const password = document.getElementById('reg-password').value;
        const confirmPassword = document.getElementById('confirm-password').value;

        if (password !== confirmPassword) {
            alert('Passwords do not match.');
            return;
        }

        // Replace 'your-graphql-endpoint' with your actual GraphQL endpoint
        const graphqlEndpoint = 'http://localhost:8080/graphql';

        const gqlMutation = `
            mutation {
                createUser(email: "${email}", password: "${password}") {
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