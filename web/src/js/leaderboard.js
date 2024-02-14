document.addEventListener('DOMContentLoaded', function () {
    const graphqlEndpoint = 'http://localhost:8080/graphql';
    const gqlQuery = `
        query {
            getLeaderboardUsers {
                ID
                UserEmail
                Score
            }
        }
    `;

    // Get reference to the leaderboard container
    const leaderboardContainer = document.querySelector('.leaderboard-container');

    // Fetch data from GraphQL endpoint
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
            // Populate the container with data
            displayLeaderboard(leaderboardContainer, data.data.getLeaderboardUsers);
        })
        .catch(error => {
            console.error('Error making GraphQL request:', error);
            alert(`An error occurred. ${error.message}`);
        });

    function displayLeaderboard(leaderboardContainer, leaderboardUsers) {
        // Clear the existing content
        leaderboardContainer.innerHTML = '';

        const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];

        // Check if the 'users' array is defined and not empty before iterating
        if (Array.isArray(leaderboardUsers) && leaderboardUsers.length > 0) {
            // Create a table and append it to the container
            const leaderboardTable = document.createElement('table');
            leaderboardTable.classList.add('leaderboard-table');

            // Create table headers
            const headerRow = leaderboardTable.createTHead().insertRow();
            const headerColumns = ['User', 'Score'];

            headerColumns.forEach(columnName => {
                const th = document.createElement('th');
                th.textContent = columnName;
                headerRow.appendChild(th);
            });

            // Display each user in a table row
            leaderboardUsers.forEach(user => {
                const row = leaderboardTable.insertRow();

                const userEmailCell = row.insertCell();
                userEmailCell.textContent = user.UserEmail;

                const scoreCell = row.insertCell();
                scoreCell.textContent = user.Score;

                // Highlight the row for the currently logged-in user
                if (user.UserEmail === emailCookie) {
                    row.classList.add('current-user');
                }
            });

            // Append the table to the container
            leaderboardContainer.appendChild(leaderboardTable);
        } else {
            // If there are no users, display a message
            const noUsersMessage = document.createElement('p');
            noUsersMessage.textContent = 'No users available.';
            leaderboardContainer.appendChild(noUsersMessage);
        }
    }
});
