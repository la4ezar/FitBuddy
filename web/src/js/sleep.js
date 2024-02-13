document.addEventListener('DOMContentLoaded', function () {
    let currentDate = new Date();
    const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
    const graphqlEndpoint = 'http://localhost:8080/graphql';

    document.getElementById('currentDate').textContent = currentDate.toLocaleDateString();

    document.querySelector('.prev-date').addEventListener('click', function () {
        currentDate.setDate(currentDate.getDate() - 1);
        updateCurrentDate();

        fetchSleeps(currentDate);
    });

    document.querySelector('.next-date').addEventListener('click', function () {
        currentDate.setDate(currentDate.getDate() + 1);
        updateCurrentDate();

        fetchSleeps(currentDate);
    });

    document.getElementById('sleep-form').addEventListener('submit', function (event) {
        event.preventDefault();
        trackSleep();
    });

    function updateCurrentDate() {
        document.getElementById('currentDate').textContent = currentDate.toLocaleDateString();
    }

    function displaySleeps(sleeps) {
        const sleepListContainer = document.querySelector('.sleep-list');

        sleepListContainer.innerHTML = '';
        if (sleeps.length === 0) {
            const noSleepData = document.createElement('p');
            noSleepData.textContent = 'No sleep data available';
            sleepListContainer.appendChild(noSleepData);
        } else {
            sleeps.forEach(sleep => {
                const sleepItem = document.createElement('div');
                sleepItem.classList.add('sleep-item');

                const sleepTimeLabel = document.createElement('p');
                sleepTimeLabel.textContent = 'Sleep Time:';
                const sleepTime = document.createElement('p');
                sleepTime.textContent = sleep.SleepTime;

                const wakeTimeLabel = document.createElement('p');
                wakeTimeLabel.textContent = 'Wake Time:';
                const wakeTime = document.createElement('p');
                wakeTime.textContent = sleep.WakeTime;

                sleepItem.appendChild(sleepTimeLabel);
                sleepItem.appendChild(sleepTime);
                sleepItem.appendChild(wakeTimeLabel);
                sleepItem.appendChild(wakeTime);

                sleepListContainer.appendChild(sleepItem);
            });
        }
    }

    function trackSleep() {
        const sleepTime = parseCustomTimeString(document.getElementById('sleep-time').value);
        const wakeTime = parseCustomTimeString(document.getElementById('wake-time').value);

        const gqlMutation = `
            mutation {
                createSleep(userEmail: "${emailCookie}", sleepTime: "${sleepTime}", wakeTime: "${wakeTime}", date: "${currentDate.toISOString()}") {
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
                if (data.data.createSleep) {
                    fetchSleeps(currentDate);
                } else {
                    alert('Failed to track sleep. Please try again.');
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
                alert(`An error occurred. ${error.message}`);
            });
    }


    function fetchSleeps(date) {
        const gqlQuery = `
            query {
                getSleepByEmailAndDate(userEmail: "${emailCookie}", date: "${date.toISOString()}") {
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
            body: JSON.stringify({ query: gqlQuery }),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                const sleeps = data.data.getSleepByEmailAndDate;
                console.log(sleeps)
                displaySleeps(sleeps);
            })
            .catch(error => {
                console.error('Error fetching sleeps:', error);
                alert(`An error occurred while fetching sleeps. ${error.message}`);
            });
    }

    function parseCustomTimeString(customTimeString) {
        const [hours, minutes] = customTimeString.split(':');
        console.log(hours, minutes)
        let tempCurrDay = new Date();
        tempCurrDay.setHours(hours);
        tempCurrDay.setMinutes(minutes);

        console.log(tempCurrDay)
        return tempCurrDay.toISOString();
    }

    fetchSleeps(currentDate);
});
