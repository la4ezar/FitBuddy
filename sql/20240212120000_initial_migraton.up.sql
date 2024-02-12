GRANT ALL PRIVILEGES ON DATABASE fitbuddy TO postgres;

-- CREATE SCHEMA fitbuddy;
--
-- GRANT ALL PRIVILEGES ON SCHEMA fitbuddy TO postgres;

CREATE TABLE users
(
    id       uuid PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    email varchar(256) NOT NULL,
    password varchar(256) NOT NULL,
    logged   bool NOT NULL
);
