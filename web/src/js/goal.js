document.addEventListener('DOMContentLoaded', function () {
    const goalForm = document.getElementById('goal-form');
    const goalsListContainer = document.querySelector('.goals-list');

    const graphqlEndpoint = 'http://localhost:8080/graphql';
    const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];

    goalForm.addEventListener('submit', function (event) {
        event.preventDefault();

        const name = document.getElementById('name').value;
        const description = document.getElementById('description').value;
        const startDate = parseCustomDateString(document.getElementById('start-date').value);
        console.log(startDate)
        const endDate = parseCustomDateString(document.getElementById('end-date').value);


        createGoal(name, description, startDate, endDate, emailCookie)
            .then(() => {
                fetchAndDisplayGoals(emailCookie);
            })
            .catch(error => {
                console.error('Error creating goal:', error);
                alert('Failed to create goal. Please try again.');
            });
    });

    const createGoal = (name, description, startDate, endDate, email) => {
        const createGoalMutation = `
            mutation {
                createGoal(name: "${name}", description: "${description}", startDate: "${startDate}", endDate: "${endDate}", email: "${email}") {
                    ID
                }
            }
        `;

        return fetch(graphqlEndpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ query: createGoalMutation }),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                if (!data.data.createGoal.ID) {
                    throw new Error('Failed to create goal.');
                }
            });
    };

    const fetchAndDisplayGoals = (email) => {
        const getGoalsQuery = `
            query {
                getGoals(email: "${email}") {
                    ID
                    Name
                    Description
                    StartDate
                    EndDate
                }
            }
        `;

        fetch(graphqlEndpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ query: getGoalsQuery }),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                console.log(data.data)
                displayGoals(data.data.getGoals);
            })
            .catch(error => {
                console.error('Error fetching goals:', error);
            });
    };

    const displayGoals = (goals) => {
        goalsListContainer.innerHTML = '';

        goals.forEach(goal => {
            const goalCard = document.createElement('div');
            goalCard.classList.add('goal-card');

            const nameElement = document.createElement('h3');
            nameElement.textContent = `Name: ${goal.Name}`;
            goalCard.appendChild(nameElement);

            const descriptionElement = document.createElement('p');
            descriptionElement.textContent = `Description: ${goal.Description}`;
            goalCard.appendChild(descriptionElement);

            const startDateElement = document.createElement('p');
            let goalStartDate = `${goal.StartDate}`.split(" ")[0]
            startDateElement.textContent = `Start Date: ${goalStartDate}`;
            goalCard.appendChild(startDateElement);

            const endDateElement = document.createElement('p');
            let goalEndDate = `${goal.EndDate}`.split(" ")[0]
            endDateElement.textContent = `End Date: ${goalEndDate}`;
            goalCard.appendChild(endDateElement);

            goalsListContainer.appendChild(goalCard);
        });
    };

    function parseCustomDateString(customDateString) {
        const [year, month, day] = customDateString.split('-');

        const parsedDate = new Date(`${year}-${month.padStart(2, '0')}-${day.padStart(2, '0')}T00:00:00Z`);

        return parsedDate.toISOString();
    }

    fetchAndDisplayGoals(emailCookie);
});
