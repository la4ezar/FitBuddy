const email = getCookie();

if (email) {
    document.getElementById('email').textContent = email;
} else {
    document.title = '401 Unauthorized';
    document.body.innerHTML = '<h1>401 Unauthorized</h1><p>You are not authorized to access this page.</p>';
}

function getCookie() {
    const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];

    if (emailCookie) {
        console.log('Email:', emailCookie);
    } else {
        console.log('Email cookie not found.');
    }

    return emailCookie
}

document.addEventListener('DOMContentLoaded', function () {
    const graphqlEndpoint = 'http://localhost:8080/graphql';

    let currentDate = new Date()
    currentDate.setHours(+currentDate.getHours()+2)

    fetchAllWorkouts(currentDate.toISOString())
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

        // Make the GraphQL request to fetch all posts
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
            const headerColumns = ['Exercise', 'Reps', 'Sets', 'Weight'];

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