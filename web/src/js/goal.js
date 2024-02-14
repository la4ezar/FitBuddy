document.addEventListener('DOMContentLoaded', function () {
    const email = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
    if (!email) {
        return
    }

    const goalForm = document.getElementById('goal-form');
    const goalsListContainer = document.querySelector('.goals-list');

    const graphqlEndpoint = 'http://localhost:8080/graphql';

    goalForm.addEventListener('submit', function (event) {
        event.preventDefault();

        const name = document.getElementById('name').value;
        const description = document.getElementById('description').value;
        const startDate = parseCustomDateString(document.getElementById('start-date').value);
        const endDate = parseCustomDateString(document.getElementById('end-date').value);


        createGoal(name, description, startDate, endDate)
            .then(() => {
                fetchAndDisplayGoals();
            })
            .catch(error => {
                console.error('Error creating goal:', error);
                alert('Failed to create goal. Please try again.');
            });
    });

    const createGoal = (name, description, startDate, endDate) => {
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

    const fetchAndDisplayGoals = () => {
        const getGoalsQuery = `
            query {
                getGoals(email: "${email}") {
                    ID
                    Name
                    Description
                    StartDate
                    EndDate
                    Completed
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
            const goalElement = document.createElement('div');
            goalElement.classList.add('goal-card');
            goalElement.classList.add(`goal-${goal.ID}`);

            const nameElement = document.createElement('h3');
            nameElement.textContent = `Name: ${goal.Name}`;
            goalElement.appendChild(nameElement);

            const descriptionElement = document.createElement('p');
            descriptionElement.textContent = `Description: ${goal.Description}`;
            goalElement.appendChild(descriptionElement);

            const startDateElement = document.createElement('p');
            let goalStartDate = `${goal.StartDate}`.split(" ")[0]
            startDateElement.textContent = `Start Date: ${goalStartDate}`;
            goalElement.appendChild(startDateElement);

            const endDateElement = document.createElement('p');
            let goalEndDate = `${goal.EndDate}`.split(" ")[0]
            endDateElement.textContent = `End Date: ${goalEndDate}`;
            goalElement.appendChild(endDateElement);

            if (!goal.Completed) {
                const completeButton = document.createElement('button');
                completeButton.textContent = '✓';
                completeButton.className = 'complete-goal-button';
                completeButton.classList.add(`complete-button-${goal.ID}`);
                goalElement.appendChild(completeButton)
                completeButton.addEventListener('click', () => {
                    completeGoal(goal.ID)
                });
            } else {
                goalElement.classList.add('completed-goal');
            }

            const deleteButton = document.createElement('button');
            deleteButton.textContent = 'X';
            deleteButton.className = 'delete-goal-button';
            goalElement.appendChild(deleteButton);

            deleteButton.addEventListener('click', function (event) {
                deleteGoal(goal.ID);
            });

            goalsListContainer.appendChild(goalElement);
        });
    };

    function deleteGoal(goalID) {
        const gqlMutation = `
            mutation {
                deleteGoal(goalID: "${goalID}")
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
                    fetchAndDisplayGoals()
                } else {
                    console.error('Failed to delete goal.');
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
            });
    }

    function completeGoal(goalID) {
        const gqlMutation = `
            mutation {
                completeGoal(userEmail: "${email}", goalID: "${goalID}")
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
                console.log(data);
                if (data.data && data.data.completeGoal) {
                    const completedButton = document.querySelector(`.complete-button-${goalID}`);
                    const parent = document.querySelector(`.goal-${goalID}`);
                    if (completedButton) {
                        completedButton.remove();
                        parent.classList.add('completed-goal');
                    }
                } else {
                    console.error('Failed to complete goal.');
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
            });
    }

    function parseCustomDateString(customDateString) {
        const [year, month, day] = customDateString.split('-');

        const parsedDate = new Date(`${year}-${month.padStart(2, '0')}-${day.padStart(2, '0')}T00:00:00Z`);

        return parsedDate.toISOString();
    }

    fetchAndDisplayGoals();
});
