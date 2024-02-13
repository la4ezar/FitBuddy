GRANT ALL PRIVILEGES ON DATABASE fitbuddy TO postgres;

-- CREATE SCHEMA fitbuddy;
--
-- GRANT ALL PRIVILEGES ON SCHEMA fitbuddy TO postgres;

CREATE TABLE users
(
    id       uuid PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    email    varchar(256) NOT NULL,
    password varchar(256) NOT NULL,
    logged   bool         NOT NULL
);

CREATE TABLE exercises
(
    id   uuid PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    name varchar(256) NOT NULL
);

INSERT INTO exercises (id, name)
VALUES ('00000000-0000-0000-0000-000000000001', 'Bench press'),
       ('00000000-0000-0000-0000-000000000002', 'Overhead press'),
       ('00000000-0000-0000-0000-000000000003', 'Deadlift'),
       ('00000000-0000-0000-0000-000000000004', 'Squat'),
       ('00000000-0000-0000-0000-000000000005', 'Resistance band'),
       ('00000000-0000-0000-0000-000000000006', 'Biceps curl'),
       ('00000000-0000-0000-0000-000000000007', 'Triceps extension'),
       ('00000000-0000-0000-0000-000000000008', 'Pull up'),
       ('00000000-0000-0000-0000-000000000009', 'Dips'),
       ('00000000-0000-0000-0000-000000000010', 'Push ups'),
       ('00000000-0000-0000-0000-000000000011', 'Sit ups');

CREATE TABLE workouts
(
    id          uuid PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    user_id     uuid      NOT NULL,
    exercise_id uuid      NOT NULL,
    sets        int       NOT NULL,
    reps        int       NOT NULL,
    weight      float     NOT NULL,
    logged_at   timestamp NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) on DELETE CASCADE,
    FOREIGN KEY (exercise_id) REFERENCES exercises (id) on DELETE CASCADE
);

CREATE TABLE meals
(
    id       uuid PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    name     varchar(256) NOT NULL,
    calories int          NOT NULL
);

INSERT INTO meals (id, name, calories)
VALUES ('00000000-0000-0000-0000-000000000001', 'Egg', 54),
       ('00000000-0000-0000-0000-000000000002', 'Beef', 212),
       ('00000000-0000-0000-0000-000000000003', 'Pork', 189),
       ('00000000-0000-0000-0000-000000000004', 'Chicken', 101),
       ('00000000-0000-0000-0000-000000000005', 'Potato', 88),
       ('00000000-0000-0000-0000-000000000006', 'Rice', 86),
       ('00000000-0000-0000-0000-000000000007', 'Tomato', 17),
       ('00000000-0000-0000-0000-000000000008', 'Cucumber', 12),
       ('00000000-0000-0000-0000-000000000009', 'Soup', 354),
       ('00000000-0000-0000-0000-000000000010', 'Cake', 678);

CREATE TABLE nutrition
(
    id        uuid PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    user_id   uuid      NOT NULL,
    meal_id   uuid      NOT NULL,
    logged_at timestamp NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) on DELETE CASCADE,
    FOREIGN KEY (meal_id) REFERENCES meals (id) on DELETE CASCADE
);

CREATE TABLE sleep
(
    id         uuid PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    user_id    uuid      NOT NULL,
    sleep_time timestamp NOT NULL,
    wake_time  timestamp NOT NULL,
    logged_at  timestamp NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) on DELETE CASCADE
);

CREATE TABLE goals
(
    id          uuid PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    user_id     uuid         NOT NULL,
    name        varchar(256) NOT NULL,
    description varchar(256) NOT NULL,
    start_date  timestamp    NOT NULL,
    end_date    timestamp    NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) on DELETE CASCADE
);

CREATE TABLE coaches
(
    id        uuid PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    image_url varchar(256) NOT NULL,
    name      varchar(256) NOT NULL,
    specialty varchar(256) NOT NULL
);

INSERT INTO coaches (id, image_url, name, specialty)
VALUES ('00000000-0000-0000-0000-000000000001',
        'https://media.istockphoto.com/id/856797530/photo/portrait-of-a-beautiful-woman-at-the-gym.jpg?s=612x612&w=0&k=20&c=0wMa1MYxt6HHamjd66d5__XGAKbJFDFQyu9LCloRsYU=',
        'Maria Ilieva', 'Fitness Coach'),
       ('00000000-0000-0000-0000-000000000002',
        'https://media.istockphoto.com/id/1072395722/photo/fitness-trainer-at-gym.jpg?s=612x612&w=0&k=20&c=3VBLCgbxG3bGNRp9Sc3tN_7G-g_DxXhGk9rhuZo-jkE=',
        'Atanas Kolev', 'Nutrition Coach');

CREATE TABLE bookings
(
    id       uuid PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    user_id  uuid NOT NULL,
    coach_id uuid NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) on DELETE CASCADE,
    FOREIGN KEY (coach_id) REFERENCES coaches (id) on DELETE CASCADE
);

CREATE TABLE leaderboard
(
    id      uuid PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    user_id uuid NOT NULL,
    score   int  NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) on DELETE CASCADE
);

CREATE TABLE posts
(
    id         uuid PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    user_id    uuid         NOT NULL,
    title      varchar(256) NOT NULL,
    content    varchar(256) NOT NULL,
    created_at timestamp    NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) on DELETE CASCADE
);