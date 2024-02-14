document.addEventListener('DOMContentLoaded', function () {
    const graphqlEndpoint = 'http://localhost:8080/graphql';
    const email = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
    if (!email) {
        return
    }

    let currentDate = new Date();
    currentDate.setHours(+currentDate.getHours() + 2);

    document.getElementById('currentDate').textContent = currentDate.toLocaleDateString();

    document.querySelector('.prev-date').addEventListener('click', function () {
        currentDate.setDate(currentDate.getDate() - 1);
        updateCurrentDate();
        fetchAllNutritions(currentDate.toISOString());
    });

    document.querySelector('.next-date').addEventListener('click', function () {
        currentDate.setDate(currentDate.getDate() + 1);
        updateCurrentDate();
        fetchAllNutritions(currentDate.toISOString());
    });

    function updateCurrentDate() {
        document.getElementById('currentDate').textContent = currentDate.toLocaleDateString();
    }

    const gqlQuery = `
            query {
                getAllMeals() {
                    ID
                    Name
                }
            }
        `;
    const mealDatalist = document.getElementById('mealList');

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
            console.log(data.data.getAllMeals);
            data.data.getAllMeals.forEach(meal => {
                const option = document.createElement('option');
                option.value = meal.Name;
                mealDatalist.appendChild(option);
            });
        })
        .catch(error => {
            console.error('Error making GraphQL request:', error);
            alert(`An error occurred. ${error.message}`);
        });

    fetchAllNutritions(currentDate.toISOString())

    document.getElementById('nutrition-form').addEventListener('submit', async function (event) {
        event.preventDefault();

        const mealInput = document.getElementById('meal');
        const servingSizeInput = document.getElementById('serving-size');
        const numberOfServingsInput = document.getElementById('number-of-servings');

        const meal = mealInput.value;
        const servingSize = parseInt(servingSizeInput.value, 10);
        const numberOfServings = parseInt(numberOfServingsInput.value, 10);
        const date = currentDate.toISOString();

        const gqlMutation = `
            mutation {
                createNutrition(email: "${email}", meal: "${meal}", date: "${date}", servingSize: ${servingSize}, numberOfServings: ${numberOfServings}) {
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
                    fetchAllNutritions(date);
                }
            })
            .catch(error => {
                alert(`An error occurred. ${error.message}`);
            });
    });

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
                caloriesCell.textContent = (nutrition.Grams / 100 * nutrition.Calories).toFixed(0);

                const dateCell = row.insertCell();
                dateCell.textContent = new Date(nutrition.Date).toLocaleTimeString();
            });
            const footerRow = nutritionsTable.createTFoot().insertRow();
            const footerCell = footerRow.insertCell();
            footerCell.colSpan = headerColumns.length;
            footerCell.textContent = 'Total Calories: ' + calculateTotalCalories(nutritions);

            nutritionsListContainer.appendChild(nutritionsTable);
        } else {
            const noNutritionsMessage = document.createElement('p');
            noNutritionsMessage.textContent = 'No nutritions available.';
            nutritionsListContainer.appendChild(noNutritionsMessage);
        }
    }
});

function calculateTotalCalories(nutritions) {
    const totalCalories = nutritions.reduce((sum, nutrition) => {
        return sum + (nutrition.Grams / 100 * nutrition.Calories);
    }, 0);

    return totalCalories.toFixed(0);
}