document.addEventListener('DOMContentLoaded', function () {
    const coachesTableBody = document.querySelector('.coaches-container tbody');

    // Example data (replace this with actual coach data)
    const coachesData = [
        { image: '../images/coach1.png', name: 'Maria Ilieva', specialty: 'Fitness Training' },
        { image: '../images/coach2.png', name: 'Atanas Kolev', specialty: 'Nutrition' },
        // Add more coaches as needed
    ];

    // Loop through coachesData and create table rows
    coachesData.forEach(coach => {
        const row = document.createElement('tr');

        // Add columns
        const imageCell = document.createElement('td');
        const image = document.createElement('img');
        image.src = coach.image; // Assuming coach.image is the image path
        image.alt = 'Coach Image'; // Provide alt text for accessibility
        image.style.maxWidth = '200px'; // Set the maximum width
        image.style.maxHeight = '200px'; // Set the maximum height
        imageCell.appendChild(image);
        row.appendChild(imageCell);

        const nameCell = document.createElement('td');
        nameCell.textContent = coach.name;
        row.appendChild(nameCell);

        const specialtyCell = document.createElement('td');
        specialtyCell.textContent = coach.specialty;
        row.appendChild(specialtyCell);

        const bookNowCell = document.createElement('td');
        const bookNowButton = document.createElement('button');
        bookNowButton.textContent = 'Book Now';
        // Add any event listeners or functionality for the "Book Now" button here
        bookNowCell.appendChild(bookNowButton);

        row.appendChild(bookNowCell);

        // Add the row to the table
        coachesTableBody.appendChild(row);
    });
});
