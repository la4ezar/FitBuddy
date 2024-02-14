## Overview:
FitBuddy is your dedicated companion for achieving health and fitness goals. Gain access to exercise tracking, customized workout plans, nutrition monitoring, and professional coaching expertise. Set and monitor fitness goals while finding motivation on your journey to a healthier lifestyle with FitBuddy.

## Functionality:

- **Account Management:**
    - **Create an Account:** Users must create an account to unlock the full capabilities of the FitBuddy app.
    - **Log In:** Access the app by logging into your account.
    - **Log Out:** Log out of your account by clicking the "Log out" button in the top right corner.

- **Dashboard:**
    - View a personalized dashboard upon logging in, displaying daily workout, nutrition, and sleep summaries.
    - Navigate to the dashboard anytime by clicking the "FitBuddy" logo in the top left corner.

- **Workout Tracker:**
    - Access the workout tracker by clicking the "Workout" tab.
    - Navigate through different days using "Previous Date" and "Next Date" buttons.
    - Enter workout details (exercise, sets, reps, weight) and track the workout.
    - View the "All Workouts" table showcasing entered data.

- **Meal Tracker:**
    - Access the meal tracker through the "Nutrition" tab.
    - Navigate through days using date navigation buttons.
    - Enter meals (Meal, Serving size, Number of servings) and track them.
    - View the "All Nutritions" table displaying entered data, including calories and grams consumed.

- **Sleep Tracker:**
    - Access the sleep tracker via the "Sleep" tab.
    - Navigate through days using date navigation buttons.
    - Enter sleep details (Sleep time, Wake time) and track your sleep.
    - View "All Sleep For Today" notes displaying entered data.

- **Goal Tracker:**
    - Access the goal tracker through the "Goals" tab.
    - Enter goals (Name, Description, Start date, End date) and track them.
    - View "All Goals" notes showcasing entered goals.
    - Remove goals using the "X" button or mark them as completed using the checkmark button.

- **Coaches:**
    - Access the coaches page through the "Coaches" tab.
    - View coach details (picture, name, specialty, availability).
    - Book a coach by clicking "Book now" (availability permitting).
    - Unbook coaches, with restrictions on coaches booked by other users.
    - Ability to book more than one coach if needed.

- **Leaderboard:**
    - Access the leaderboard through the "Leaderboard" tab.
    - View all registered users and their scores.
    - Score increases upon completing goals (1 point per goal).

- **Forum:**
    - Access the forum through the "Forum" tab.
    - Create posts by entering "Title" and "Content" fields.
    - Add posts to the board by clicking "Add Post."
    - View your posts and others' contributions to the post board.

## Architecture

- **Style**
  - Client-Server architecture with a microservices approach.
  - Frontend: HTML, CSS, JS making GraphQL calls.
  - Backend: Go (Golang) handling GraphQL requests and interacting with PostgreSQL database.

- **Components and Modules**
  - **Frontend:**
    - Utilizes GraphQL queries to communicate with the backend.
    - Dynamic user interface components for different functionalities.

  - **Backend:**
    - Root Resolver: Handles incoming GraphQL requests and routes them to the corresponding resolvers.
    - Resolvers: Process specific GraphQL queries and mutations.
    - Service Layer: Implements business logic and interacts with the repository layer.
    - Repository Layer: Manages database interactions using PostgreSQL.

- **Error Handling and Logging**
  - Comprehensive error handling in each layer.
  - Centralized logging for monitoring and debugging.

## Realization

- **Mutations**
    - User
      - createUser(email: String!, password: String!) User - Creates a user and leaderboard record for the user
      - loginUser(email: String!, password: String!) User - Log in a user
      - logoutUser(email: String!) User - Log out a user
    - Post
      - createPost(title: String!, content: String!, email: String!) Post - Creates a post
      - deletePost(postID: ID!) Boolean! - Deletes a post
    - Coach
      - bookCoach(email: ID!, coachName: ID!) Boolean! - Books a coach
      - unbookCoach(email: ID!, coachName: ID!) Boolean! - Unbooks a coach
    - Goal 
      - createGoal(name: String!, description: String!, startDate: String!, endDate: String!, email: String!) Goal - Creates a goal
      - completeGoal(userEmail: String!, goalID: ID!) Boolean! - Completes a goal
      - deleteGoal(goalID: ID!) Boolean! - Deletes a goal
    - Sleep 
      - createSleepLog(userEmail: String!, sleepLogTime: String!, wakeTime: String!, date: String!) SleepLog - Creates a sleep log
      - deleteSleepLog(sleepLogID: ID!) Boolean! - Deletes a sleep log
    - Nutrition  
      - createNutrition(email: String!, meal: String!, date: String!, servingSize: Int!, numberOfServings: Int!): Nutrition
      - deleteNutrition(nutritionID: ID!) Boolean! - Deletes a nutrition
    - Workout 
      - createWorkout(email: String!, exercise: String!, date: String!, sets: Int!, reps: Int!, weight: Float!): Workout
      - deleteWorkout(workoutID: ID!) Boolean! - Deletes a workout
- **Queries**
    - Coach
      - getAllCoaches [Coach!]! - Gets all coaches
      - isCoachBookedByUser(coachName: String!, userEmail: String!) Boolean! - Checks if a coach is booked by user
      - isCoachBooked(coachName: String!) Boolean! - Checks if coach is booked
    - Goal
      - getGoals(email: String!) [Goal!]! - Gets goals
    - Post
      - getAllPosts: [Post!]! - Gets all posts 
    - Sleep
      - getSleepLogByEmailAndDate(userEmail: String!, date: String!) [SleepLog!]! - Gets sleep log by email and date 
    - Exercise
      - getAllExercises: [Exercise!]! - Gets all exercises
    - Workout
      - getAllWorkoutsByEmailAndDate(email: String!, date: String!) [Workout!]! - Gets all workouts by email and date
    - Nutrition
      - getAllNutritionsByEmailAndDate(email: String!, date: String!) [Nutrition!]! - Gets all nutritions by email and date
    - Meal
      - getAllMeals: [Meal!]! - Gets all meals
    - Leaderboard
      - getLeaderboardUsers: [LeaderboardUser!]! - Gets the leaderboard

## Data models
- **User** 
```bash
type User {
  ID: ID!
  Email: String!
}
  ```
- **Workout**
```bash
type Workout {
  ID: ID!
  UserEmail: String!
  ExerciseName: String!
  Sets: Int!
  Reps: Int!
  Weight: Float!
  Date: String!
}
```
- **Exercise**
```bash
type Exercise {
  ID: ID!
  Name: String!
}
```
- **Meal**
```bash
type Meal {
  ID: ID!
  Name: String!
}
```
- **Nutrition**
```bash
type Nutrition {
  ID: ID!
  UserEmail: String!
  MealName: String!
  Grams: Int!
  Calories: Int!
  Date: String!
}
```
- **Post**
```bash
type Post {
  ID: ID!
  UserEmail: String!
  Title: String!
  Content: String!
  CreatedAt: String!
}
```
- **Coach**
```bash
type Coach {
  ID: ID!
  ImageUrl: String!
  Name: String!
  Specialty: String!
}
```
- **Goal**
```bash
type Goal {
  ID: ID!
  Name: String!
  Description: String!
  StartDate: String!
  EndDate: String!
  Completed: Boolean!
}
```
- **SleepLog**
```bash
type SleepLog {
  ID: ID!
  SleepTime: String!
  WakeTime: String!
}
```
- **LeaderboardUser**
```bash
type LeaderboardUser {
  ID: ID!
  UserEmail: String!
  Score: Int!
}
```

## Configurations
 - config/database_config.env - Contains configuration about the PosgreSQL database.

## Technologies and libraries
 - **Docker**
 - **Golang**
 - **PostgreSQL**
 - **Database package - “database/sql”**
 - **Viper package for working with environment configs - “github.com/spf13/viper”**
 - **Migrate - “github.com/golang-migrate/migrate/v4”**