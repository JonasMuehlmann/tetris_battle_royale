CREATE TABLE Users (
    id SERIAL PRIMARY KEY NOT NULL,
    username TEXT UNIQUE NOT NULL,
    pw_hash BYTEA NOT NULL,
    salt BYTEA NOT NULL
);

CREATE TABLE Sessions (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INTEGER UNIQUE NOT NULL,
    creation_time TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES Users(id)
);

