document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.getElementById('login-form');

    loginForm.addEventListener('submit', function (event) {
        event.preventDefault();

        const email = document.getElementById('login-email').value;
        const password = document.getElementById('login-password').value;

        const graphqlEndpoint = 'http://localhost:8080/graphql';

        const gqlMutation = `
            mutation {
                loginUser(email: "${email}", password: "${password}") {
                    ID
                }
            }
        `;

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
                if (data.errors) {
                    alert('Login failed. Please try again.');
                } else {
                    window.location.href = './dashboard.html';
                }
            })
            .catch(error => {
                alert(`An error occurred. ${error.message}`);
            });

        document.cookie = `email=${email}; SameSite=Lax; path=/;`;
    });
});