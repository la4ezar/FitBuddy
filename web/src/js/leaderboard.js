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

    const leaderboardContainer = document.querySelector('.leaderboard-container');

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
            displayLeaderboard(leaderboardContainer, data.data.getLeaderboardUsers);
        })
        .catch(error => {
            console.error('Error making GraphQL request:', error);
            alert(`An error occurred. ${error.message}`);
        });

    function displayLeaderboard(leaderboardContainer, leaderboardUsers) {
        leaderboardContainer.innerHTML = '';

        const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];

        if (Array.isArray(leaderboardUsers) && leaderboardUsers.length > 0) {
            const leaderboardTable = document.createElement('table');
            leaderboardTable.classList.add('leaderboard-table');

            const headerRow = leaderboardTable.createTHead().insertRow();
            const headerColumns = ['User', 'Score'];

            headerColumns.forEach(columnName => {
                const th = document.createElement('th');
                th.textContent = columnName;
                headerRow.appendChild(th);
            });

            leaderboardUsers.forEach(user => {
                const row = leaderboardTable.insertRow();

                const userEmailCell = row.insertCell();
                userEmailCell.textContent = user.UserEmail;

                const scoreCell = row.insertCell();
                scoreCell.textContent = user.Score;

                if (user.UserEmail === emailCookie) {
                    row.classList.add('current-user');
                }
            });

            leaderboardContainer.appendChild(leaderboardTable);
        } else {
            const noUsersMessage = document.createElement('p');
            noUsersMessage.textContent = 'No users available.';
            leaderboardContainer.appendChild(noUsersMessage);
        }
    }
});
