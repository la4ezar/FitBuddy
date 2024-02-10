// Sample data - replace this with actual data from your server
const userData = {
    name: 'John Doe',
    fitnessStats: 'Sample Fitness Stats',
    recentWorkouts: [
        { id: 1, exercise: 'Running', date: '2022-02-01' },
        { id: 2, exercise: 'Weightlifting', date: '2022-02-03' },
    ],
    recentNutritionLogs: [
        { id: 1, description: 'Breakfast', date: '2022-02-02' },
        { id: 2, description: 'Lunch', date: '2022-02-04' },
    ],
};

document.addEventListener('DOMContentLoaded', function () {
    const welcomeMessage = document.getElementById('welcome-message');
    const fitnessStats = document.getElementById('fitness-stats');
    const workoutsList = document.getElementById('workouts-list');
    const nutritionLogsList = document.getElementById('nutrition-logs-list');

    // Populate welcome message
    welcomeMessage.innerText = `Welcome, ${userData.name}!`;

    // Populate fitness stats (replace with actual content)
    fitnessStats.innerHTML = '<h3>Fitness Stats</h3><p>' + userData.fitnessStats + '</p>';

    // Populate recent workouts
    userData.recentWorkouts.forEach((workout) => {
        const li = document.createElement('li');
        li.innerText = `${workout.exercise} - ${workout.date}`;
        workoutsList.appendChild(li);
    });

    // Populate recent nutrition logs
    userData.recentNutritionLogs.forEach((nutritionLog) => {
        const li = document.createElement('li');
        li.innerText = `${nutritionLog.description} - ${nutritionLog.date}`;
        nutritionLogsList.appendChild(li);
    });
});