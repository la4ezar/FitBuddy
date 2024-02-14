document.addEventListener('DOMContentLoaded', function () {
    const email = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
    if (!email) {
        return
    }
    const graphqlEndpoint = 'http://localhost:8080/graphql';
    const coachesTableBody = document.querySelector('.coaches-container tbody');

    function fetchCoaches() {
        const gqlQuery = `
            query {
                getAllCoaches {
                    ID
                    ImageUrl
                    Name
                    Specialty
                }
            }
        `;

        return fetch(graphqlEndpoint, {
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
            .then(data => data.data.getAllCoaches)
            .catch(error => {
                console.error('Error fetching coaches:', error);
                return [];
            });
    }

    function isBookedByCurrentUser(coach) {
        const isBookedQuery = `
            query {
                isCoachBookedByUser(coachName: "${coach.Name}", userEmail: "${email}")
            }
        `;

        return fetch(graphqlEndpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({query: isBookedQuery}),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => data.data.isCoachBookedByUser)
            .catch(error => {
                console.error('Error checking if coach is booked:', error);
                return false;
            });
    }

    function isBookedByAnyUser(coach) {
        const isBookedQuery = `
            query {
                isCoachBooked(coachName: "${coach.Name}")
            }
        `;

        return fetch(graphqlEndpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({query: isBookedQuery}),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => data.data.isCoachBooked)
            .catch(error => {
                console.error('Error checking if coach is booked:', error);
                return false;
            });
    }

    function displayCoaches(coaches) {
        coachesTableBody.innerHTML = '';

        coaches.forEach(coach => {
            const row = document.createElement('tr');

            const imageCell = document.createElement('td');
            const image = document.createElement('img');

            image.src = coach.ImageUrl.startsWith('http') ? coach.ImageUrl : `http://${coach.ImageUrl}`;
            image.alt = 'Coach Image';
            image.style.maxWidth = '200px';
            image.style.maxHeight = '200px';

            imageCell.appendChild(image);
            row.appendChild(imageCell);

            const nameCell = document.createElement('td');
            nameCell.textContent = coach.Name;
            row.appendChild(nameCell);

            const specialtyCell = document.createElement('td');
            specialtyCell.textContent = coach.Specialty;
            row.appendChild(specialtyCell);

            const bookNowCell = document.createElement('td');
            const bookNowButton = document.createElement('button');
            bookNowButton.textContent = 'Book Now';


            isBookedByCurrentUser(coach)
                .then(isBooked => {
                    if (isBooked) {
                        const unbookButton = document.createElement('button');
                        unbookButton.textContent = 'Unbook Now';
                        unbookButton.addEventListener('click', () => unbookCoach(coach.Name, email));
                        bookNowCell.appendChild(unbookButton);
                    } else {
                        isBookedByAnyUser(coach)
                            .then(isBooked => {
                                if (isBooked) {
                                    const bookedText = document.createElement('span');
                                    bookedText.textContent = 'Booked by another user';
                                    bookNowCell.appendChild(bookedText);
                                } else {
                                    const bookNowButton = document.createElement('button');
                                    bookNowButton.textContent = 'Book Now';
                                    bookNowButton.addEventListener('click', () => bookCoach(coach.Name, email));
                                    bookNowCell.appendChild(bookNowButton);
                                }
                            })
                            .catch(error => {
                                console.error('Error checking if coach is booked:', error);
                            });
                    }
                })
                .catch(error => {
                    console.error('Error checking if coach is booked:', error);
                    bookNowCell.appendChild(bookNowButton);
                });

            row.appendChild(bookNowCell);

            coachesTableBody.appendChild(row);
        });
    }

    function bookCoach(coachName) {
        const bookCoachMutation = `
            mutation {
                bookCoach(email: "${email}", coachName: "${coachName}")
            }
        `;

        fetch(graphqlEndpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({query: bookCoachMutation}),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                if (data.data.bookCoach) {
                    fetchCoaches()
                        .then(coaches => {
                            displayCoaches(coaches);
                        })
                        .catch(error => {
                            console.error('Error fetching and displaying coaches:', error);
                        });
                } else {
                    alert('Failed to book the coach. Please try again.');
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
                alert(`An error occurred. ${error.message}`);
            });
    }

    function unbookCoach(coachName) {
        const unbookCoachMutation = `
            mutation {
                unbookCoach(email: "${email}", coachName: "${coachName}")
            }
        `;

        fetch(graphqlEndpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({query: unbookCoachMutation}),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                if (data.data.unbookCoach) {
                    fetchCoaches()
                        .then(coaches => {
                            displayCoaches(coaches);
                        })
                        .catch(error => {
                            console.error('Error fetching and displaying coaches:', error);
                        });
                } else {
                    alert('Failed to unbook the coach. Please try again.');
                }
            })
            .catch(error => {
                console.error('Error making GraphQL request:', error);
                alert(`An error occurred. ${error.message}`);
            });
    }

    fetchCoaches()
        .then(coaches => {
            displayCoaches(coaches);
        })
        .catch(error => {
            console.error('Error fetching and displaying coaches:', error);
        });
});