document.addEventListener('DOMContentLoaded', function () {
    const logoutButton = document.querySelector('.log-out-button');

    logoutButton.addEventListener('click', function () {
        const graphqlEndpoint = 'http://localhost:8080/graphql';

        const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
        if (emailCookie) {
            console.log('Email:', emailCookie);
        } else {
            console.log('Email cookie not found.');
        }

        const gqlMutation = `
            mutation {
                logoutUser(email: "${emailCookie}") {
                    ID
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
                console.log('GraphQL Response:', data);
                if (data.errors) {
                    alert('Registration failed. Please try again.');
                } else {
                    document.cookie = `email=; SameSite=Lax; path=/;`;
                    window.location.href = './login.html'; // Update with the actual path
                }
            })
            .catch(error => {
                alert(`An error occurred. ${error.message}`);
            });
    });
});