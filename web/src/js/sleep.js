document.addEventListener('DOMContentLoaded', function () {
    const email = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
    if (!email) {
        return
    }

    let currentDate = new Date();
    const graphqlEndpoint = 'http://localhost:8080/graphql';

    document.getElementById('currentDate').textContent = currentDate.toLocaleDateString();

    document.querySelector('.prev-date').addEventListener('click', function () {
        currentDate.setDate(currentDate.getDate() - 1);
        updateCurrentDate();

        fetchSleepLogs(currentDate);
    });

    document.querySelector('.next-date').addEventListener('click', function () {
        currentDate.setDate(currentDate.getDate() + 1);
        updateCurrentDate();

        fetchSleepLogs(currentDate);
    });

    document.getElementById('sleep-form').addEventListener('submit', function (event) {
        event.preventDefault();
        trackSleepLog();
    });

    function updateCurrentDate() {
        document.getElementById('currentDate').textContent = currentDate.toLocaleDateString();
    }

    function displaySleepLogs(sleepLogs) {
        const sleepListContainer = document.querySelector('.sleep-list');

        sleepListContainer.innerHTML = '';
        if (sleepLogs.length === 0) {
            const noSleepData = document.createElement('p');
            noSleepData.textContent = 'No sleep data available';
            sleepListContainer.appendChild(noSleepData);
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

                const deleteButton = document.createElement('button');
                deleteButton.textContent = 'X';
                deleteButton.className = 'delete-sleep-button';
                sleepItem.appendChild(deleteButton);
                console.log(sleepLog.ID)
                deleteButton.addEventListener('click', function (event) {
                    deleteSleep(sleepLog.ID);
                });

                sleepListContainer.appendChild(sleepItem);
            });
        }
    }

    function trackSleepLog() {
        const sleepTime = parseCustomTimeString(document.getElementById('sleep-time').value);
        const wakeTime = parseCustomTimeString(document.getElementById('wake-time').value);

        let newDate = new Date(currentDate);
        newDate.setHours(+newDate.getHours() + 2);

        const gqlMutation = `
            mutation {
                createSleepLog(userEmail: "${email}", sleepLogTime: "${sleepTime}", wakeTime: "${wakeTime}", date: "${newDate.toISOString()}") {
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
                if (data.data.createSleepLog) {
                    fetchSleepLogs(currentDate);
                } else {
                    alert('Failed to track sleep. Please try again.');
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
                alert(`An error occurred. ${error.message}`);
            });
    }


    function fetchSleepLogs(date) {
        let newDate = new Date(date);
        newDate.setHours(+newDate.getHours() + 2);

        const gqlQuery = `
            query {
                getSleepLogByEmailAndDate(userEmail: "${email}", date: "${newDate.toISOString()}") {
                    ID
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
                console.error('Error fetching sleeps:', error);
                alert(`An error occurred while fetching sleeps. ${error.message}`);
            });
    }

    function deleteSleep(sleepLogID) {
        const gqlMutation = `
            mutation {
                deleteSleepLog(sleepLogID: "${sleepLogID}")
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
                    fetchSleepLogs(currentDate);
                } else {
                    console.error('Failed to delete sleep.');
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
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

    fetchSleepLogs(currentDate);
});
