document.addEventListener('DOMContentLoaded', function () {
    // Get the current date
    let currentDate = new Date();

    // Display the current date
    document.getElementById('currentDate').textContent = currentDate.toLocaleDateString();

    // Add event listener for the Previous Date button
    document.querySelector('.prev-date').addEventListener('click', function () {
        currentDate.setDate(currentDate.getDate() - 1);
        updateCurrentDate();
        fetchAllNutritions(currentDate.toISOString());
    });

    // Add event listener for the Next Date button
    document.querySelector('.next-date').addEventListener('click', function () {
        currentDate.setDate(currentDate.getDate() + 1);
        updateCurrentDate();
        fetchAllNutritions(currentDate.toISOString());
    });

    // Function to update the displayed current date
    function updateCurrentDate() {
        document.getElementById('currentDate').textContent = currentDate.toLocaleDateString();
    }

    const graphqlEndpoint = 'http://localhost:8080/graphql';
    const gqlQuery = `
            query {
                getAllMeals() {
                    ID
                    Name
                }
            }
        `;
    // Get reference to the meal input and datalist
    const mealDatalist = document.getElementById('mealList');

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
            console.log(data.data.getAllMeals);
            // Populate the datalist with exercise options
            data.data.getAllMeals.forEach(meal => {
                const option = document.createElement('option');
                option.value = meal.Name; // Replace 'name' with the actual property of your exercise object
                mealDatalist.appendChild(option);
            });
        })
        .catch(error => {
            console.error('Error making GraphQL request:', error);
            alert(`An error occurred. ${error.message}`);
        });

    fetchAllNutritions(currentDate.toISOString())

    // Add event listener for the workout form submission
    document.getElementById('nutrition-form').addEventListener('submit', async function (event) {
        event.preventDefault();

        const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
        if (emailCookie) {
            console.log('Email:', emailCookie);
        } else {
            console.log('Email cookie not found.');
        }

        const mealInput = document.getElementById('meal');
        const servingSizeInput = document.getElementById('serving-size');
        const numberOfServingsInput = document.getElementById('number-of-servings');

        const meal = mealInput.value;
        const servingSize = parseInt(servingSizeInput.value, 10);
        const numberOfServings = parseInt(numberOfServingsInput.value, 10);
        const date = currentDate.toISOString();

        const graphqlEndpoint = 'http://localhost:8080/graphql';

        // GraphQL mutation to create a workout
        const gqlMutation = `
            mutation {
                createNutrition(email: "${emailCookie}", meal: "${meal}", date: "${date}", servingSize: ${servingSize}, numberOfServings: ${numberOfServings}) {
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
        const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
        if (emailCookie) {
            console.log('Email:', emailCookie);
        } else {
            console.log('Email cookie not found.');
        }
        // Replace 'your-graphql-endpoint' with your actual GraphQL endpoint
        const graphqlEndpoint = 'http://localhost:8080/graphql';

        const gqlQuery = `
            query {
                getAllNutritionsByEmailAndDate(email: "${emailCookie}", date: "${date}") {
                    ID
                    UserEmail
                    MealName
                    Grams
                    Calories
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

        // Clear the existing posts
        nutritionsListContainer.innerHTML = '';

        // Check if the 'posts' array is defined and not empty before iterating
        if (Array.isArray(nutritions) && nutritions.length > 0) {
            // Display each post
            // Assuming workoutsListContainer is the container where you want to append the table
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
            footerCell.colSpan = headerColumns.length; // Span all columns except the first one
            footerCell.textContent = 'Total Calories: ' + calculateTotalCalories(nutritions);

            // Append the table to the container
            nutritionsListContainer.appendChild(nutritionsTable);

        } else {
            // If there are no posts, display a message
            const noNutritionsMessage = document.createElement('p');
            noNutritionsMessage.textContent = 'No nutritions available.';
            nutritionsListContainer.appendChild(noNutritionsMessage);
        }
    }
});

// Function to calculate the total calories
function calculateTotalCalories(nutritions) {
    const totalCalories = nutritions.reduce((sum, nutrition) => {
        return sum + (nutrition.Grams / 100 * nutrition.Calories);
    }, 0);

    return totalCalories.toFixed(0);
}