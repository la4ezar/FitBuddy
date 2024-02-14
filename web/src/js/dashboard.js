document.addEventListener('DOMContentLoaded', function () {
    const graphqlEndpoint = 'http://localhost:8080/graphql';
    const email = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
    if (!email) {
        return
    }
    document.getElementById('email').textContent = email;

    let currentDate = new Date()
    currentDate.setHours(+currentDate.getHours()+2)

    fetchAllWorkouts(currentDate.toISOString())
    fetchAllNutritions(currentDate.toISOString())
    fetchSleepLogs(currentDate.toISOString())

    function fetchAllWorkouts(date) {
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

        fetch(graphqlEndpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({query: gqlQuery}),
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

        workoutsListContainer.innerHTML = '';

        if (Array.isArray(workouts) && workouts.length > 0) {
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
            const workoutsTitle = document.querySelector('.workouts-title');
            workoutsTitle.innerHTML = 'No workouts today'

            const noPostsMessage = document.createElement('p');
            noPostsMessage.textContent = '\"The only bad workout is the one that didn\'t happen.\"';
            workoutsListContainer.appendChild(noPostsMessage);
        }
    }

    function displaySleepLogs(sleepLogs) {
        const sleepListContainer = document.querySelector('.sleep-list');

        sleepListContainer.innerHTML = '';
        if (sleepLogs.length === 0) {
            const sleepTitle = document.querySelector('.sleep-title')
            sleepTitle.innerHTML = '';

        } else {
            sleepLogs.forEach(sleepLog => {
                const sleepItem = document.createElement('div');
                sleepItem.classList.add('sleep-item');

                const sleepTimeLabel = document.createElement('p');
                sleepTimeLabel.textContent = 'Sleep Time:';
                const sleepTime = document.createElement('p');
                sleepTime.textContent = sleepLog.SleepTime;

                const wakeTimeLabel = document.createElement('p');
                wakeTimeLabel.textContent = 'Wake Time:';
                const wakeTime = document.createElement('p');
                wakeTime.textContent = sleepLog.WakeTime;

                sleepItem.appendChild(sleepTimeLabel);
                sleepItem.appendChild(sleepTime);
                sleepItem.appendChild(wakeTimeLabel);
                sleepItem.appendChild(wakeTime);

                sleepListContainer.appendChild(sleepItem);
            });
        }
    }

    function fetchSleepLogs(date) {
        let newDate = new Date(date);
        newDate.setHours(+newDate.getHours() + 2);

        const gqlQuery = `
            query {
                getSleepLogByEmailAndDate(userEmail: "${email}", date: "${newDate.toISOString()}") {
                    SleepTime
                    WakeTime
                }
            }
        `;

        fetch(graphqlEndpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({query: gqlQuery}),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                const sleeps = data.data.getSleepLogByEmailAndDate;
                console.log(sleeps)
                displaySleepLogs(sleeps);
            })
            .catch(error => {
                console.error('Error fetching sleep logs:', error);
                alert(`An error occurred while fetching sleep logs. ${error.message}`);
            });
    }

    function fetchAllNutritions(date) {
        const gqlQuery = `
            query {
                getAllNutritionsByEmailAndDate(email: "${email}", date: "${date}") {
                    ID
                    UserEmail
                    MealName
                    Grams
                    Calories
                    Date
                }
            }
        `;

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
                    alert('Getting all nutritions failed. Please try again.');
                } else {
                    displayNutritions(data.data.getAllNutritionsByEmailAndDate);
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
                alert(`An error occurred. ${error.message}`);
            });
    }

    function displayNutritions(nutritions) {
        const nutritionsListContainer = document.querySelector('.nutritions-list');

        nutritionsListContainer.innerHTML = '';

        if (Array.isArray(nutritions) && nutritions.length > 0) {
            const nutritionsTable = document.createElement('table');
            nutritionsTable.className = 'nutritions-table';

            const headerRow = nutritionsTable.createTHead().insertRow();
            const headerColumns = ['Meal', 'Grams', 'Calories', 'Time'];

            headerColumns.forEach(columnName => {
                const headerCell = document.createElement('th');
                headerCell.textContent = columnName;
                headerRow.appendChild(headerCell);
            });

            nutritionsListContainer.appendChild(nutritionsTable);

            nutritions.forEach(nutrition => {
                const row = nutritionsTable.insertRow();

                const mealCell = row.insertCell();
                mealCell.textContent = nutrition.MealName;

                const gramsCell = row.insertCell();
                gramsCell.textContent = nutrition.Grams;

                const caloriesCell = row.insertCell();
                caloriesCell.textContent = (nutrition.Grams/100 * nutrition.Calories).toFixed(0);

                const dateCell = row.insertCell();
                dateCell.textContent = new Date(nutrition.Date).toLocaleTimeString();
            });
            const footerRow = nutritionsTable.createTFoot().insertRow();
            const footerCell = footerRow.insertCell();
            footerCell.colSpan = headerColumns.length;
            footerCell.textContent = 'Total Calories: ' + calculateTotalCalories(nutritions);

            nutritionsListContainer.appendChild(nutritionsTable);

        } else {
            const nutritionTitle = document.querySelector('.nutrition-title');
            nutritionTitle.innerHTML = 'No tracked nutrition today'

            const noNutritionMessage = document.createElement('p');
            noNutritionMessage.textContent = '\"Tracking your daily nutrition – a simple yet powerful tool to optimize performance and embrace a healthier lifestyle.\"';

            nutritionsListContainer.appendChild(noNutritionMessage);
        }
    }

    function calculateTotalCalories(nutritions) {
        const totalCalories = nutritions.reduce((sum, nutrition) => {
            return sum + (nutrition.Grams / 100 * nutrition.Calories);
        }, 0);

        return totalCalories.toFixed(0);
    }

    function parseCustomTimeString(customTimeString) {
        const [hours, minutes] = customTimeString.split(':');
        console.log(hours, minutes)
        let tempCurrDay = new Date();
        tempCurrDay.setHours(+hours + 2);
        tempCurrDay.setMinutes(minutes);
        tempCurrDay.setSeconds(0);

        return tempCurrDay.toISOString();
    }
});