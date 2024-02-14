document.addEventListener('DOMContentLoaded', function () {
    const email = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
    if (!email) {
        return
    }

    let currentDate = new Date();

    // Display the current date
    document.getElementById('currentDate').textContent = currentDate.toLocaleDateString();

    // Add event listener for the Previous Date button
    document.querySelector('.prev-date').addEventListener('click', function () {
        currentDate.setDate(currentDate.getDate() - 1);
        updateCurrentDate();
        fetchAllWorkouts(currentDate.toISOString());
    });

    // Add event listener for the Next Date button
    document.querySelector('.next-date').addEventListener('click', function () {
        currentDate.setDate(currentDate.getDate() + 1);
        updateCurrentDate();
        fetchAllWorkouts(currentDate.toISOString());
    });

    // Function to update the displayed current date
    function updateCurrentDate() {
        document.getElementById('currentDate').textContent = currentDate.toLocaleDateString();
    }

    const graphqlEndpoint = 'http://localhost:8080/graphql';
    const gqlQuery = `
            query {
                getAllExercises() {
                    ID
                    Name
                }
            }
        `;
    // Get reference to the exercise input and datalist
    const exerciseDatalist = document.getElementById('exerciseList');

    // Fetch the exercise list from the backend
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
            console.log(data.data.getAllExercises);
            // Populate the datalist with exercise options
            data.data.getAllExercises.forEach(exercise => {
                const option = document.createElement('option');
                option.value = exercise.Name; // Replace 'name' with the actual property of your exercise object
                exerciseDatalist.appendChild(option);
            });
        })
        .catch(error => {
            console.error('Error making GraphQL request:', error);
            alert(`An error occurred. ${error.message}`);
        });

    fetchAllWorkouts(currentDate.toISOString())

    // Add event listener for the workout form submission
    document.getElementById('workout-form').addEventListener('submit', async function (event) {
        event.preventDefault();

        const exerciseInput = document.getElementById('exercise');
        const setsInput = document.getElementById('sets');
        const repsInput = document.getElementById('reps');
        const weightInput = document.getElementById('weight');

        const exercise = exerciseInput.value;
        const sets = parseInt(setsInput.value, 10);
        const reps = parseInt(repsInput.value, 10);
        const weight = parseFloat(weightInput.value);
        const date = currentDate.toISOString();

        const graphqlEndpoint = 'http://localhost:8080/graphql';

        // GraphQL mutation to create a workout
        const gqlMutation = `
            mutation {
                createWorkout(email: "${email}", exercise: "${exercise}", date: "${date}", sets: ${sets}, reps: ${reps}, weight: ${weight}) {
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
                    alert('Creating workout failed. Please try again.');
                } else {
                    fetchAllWorkouts(date);
                }
            })
            .catch(error => {
                alert(`An error occurred. ${error.message}`);
            });
    });

    function fetchAllWorkouts(date) {
        const graphqlEndpoint = 'http://localhost:8080/graphql';

        const gqlQuery = `
            query {
                getAllWorkoutsByEmailAndDate(email: "${email}", date: "${date}") {
                    ID
                    UserEmail
                    ExerciseName
                    Sets
                    Reps
                    Weight
                    Date
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
                    alert('Getting all workouts failed. Please try again.');
                } else {
                     displayWorkouts(data.data.getAllWorkoutsByEmailAndDate);
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
                alert(`An error occurred. ${error.message}`);
            });
    }

    function displayWorkouts(workouts) {
        const workoutsListContainer = document.querySelector('.workouts-list');

        // Clear the existing posts
        workoutsListContainer.innerHTML = '';

        // Check if the 'posts' array is defined and not empty before iterating
        if (Array.isArray(workouts) && workouts.length > 0) {
            // Display each post
            // Assuming workoutsListContainer is the container where you want to append the table
            const workoutsTable = document.createElement('table');
            workoutsTable.className = 'workouts-table';

            const headerRow = workoutsTable.createTHead().insertRow();
            const headerColumns = ['Exercise', 'Reps', 'Sets', 'Weight', 'Time'];

            headerColumns.forEach(columnName => {
                const headerCell = document.createElement('th');
                headerCell.textContent = columnName;
                headerRow.appendChild(headerCell);
            });

            workoutsListContainer.appendChild(workoutsTable);

            workouts.forEach(workout => {
                const row = workoutsTable.insertRow();

                const exerciseCell = row.insertCell();
                exerciseCell.textContent = workout.ExerciseName;

                const repsCell = row.insertCell();
                repsCell.textContent = workout.Reps;

                const setsCell = row.insertCell();
                setsCell.textContent = workout.Sets;

                const weightCell = row.insertCell();
                weightCell.textContent = workout.Weight;

                const dateCell = row.insertCell();
                dateCell.textContent = new Date(workout.Date).toLocaleTimeString();
            });


        } else {
            // If there are no posts, display a message
            const noWorkoutsMessage = document.createElement('p');
            noWorkoutsMessage.textContent = 'No workouts available.';
            workoutsListContainer.appendChild(noWorkoutsMessage);
        }
    }
});