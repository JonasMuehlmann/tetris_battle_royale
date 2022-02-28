CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username TEXT,
    pw_hash BYTEA,
    salt BYTEA,
    CONSTRAINT uniq_username UNIQUE(username)
);

CREATE TABLE sessions(
id SERIAL PRIMARY KEY,
user_id INTEGER, 
creation_time TIMESTAMP WITHOUT TIME ZONE,
CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
CONSTRAINT uniq_userid UNIQUE(user_id)
);
