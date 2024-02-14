# FitBuddy
In the realm of personal fitness and wellness, FitBuddy emerges as your dedicated companion for achieving your health and fitness aspirations.

## Features

- **Workout Tracking**: Access personalized workout routines and track your progress.
- **Nutrition Management**: Log your daily nutrition intake and monitor your dietary habits.
- **Sleep Monitoring**: Keep track of your sleep patterns to ensure optimal recovery.
- **Goal Setting**: Set fitness goals and track your achievements.
- **Leaderboard**: Compete with other users and see where you stand in the fitness rankings.
- **Forum**: Engage with the fitness community, share experiences, and get motivated.
- **Coaches**: Connect with fitness coaches for personalized guidance.
- **User Authentication**: Register, log in, and log out securely to manage your fitness journey.

### Prerequisites
- Golang
- Docker

### Installation
1. Clone the repository:
```bash
git clone https://github.com/la4ezar/FitBuddy.git
```
2. Start the Postgres Database
```bash
./hack/run_db.sh
```
3. Run the Migrator which applies all the DB migrations
```bash
go run cmd/migrator/main.go
```
4. Run the Fitbuddy server
```bash
go run cmd/fitbuddy/main.go
```
5. Connect to localhost http://localhost:63342/FitBuddy/web/views/registration.html