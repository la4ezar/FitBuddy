document.addEventListener('DOMContentLoaded', function () {
    // Get the current date
    let currentDate = new Date();

    // Display the current date
    document.getElementById('currentDate').textContent = currentDate.toLocaleDateString();

    // Add event listener for the Previous Date button
    document.querySelector('.prev-date').addEventListener('click', function () {
        currentDate.setDate(currentDate.getDate() - 1);
        updateCurrentDate();
    });

    // Add event listener for the Next Date button
    document.querySelector('.next-date').addEventListener('click', function () {
        currentDate.setDate(currentDate.getDate() + 1);
        updateCurrentDate();
    });

    // Function to update the displayed current date
    function updateCurrentDate() {
        document.getElementById('currentDate').textContent = currentDate.toLocaleDateString();
    }
});
