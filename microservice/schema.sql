CREATE TABLE users(
    id         UUID PRIMARY KEY,
    username   TEXT,
    pw_hash    BYTEA,
    salt       BYTEA,
    CONSTRAINT uniq_username UNIQUE(username)
);

CREATE TABLE sessions(
    id            UUID PRIMARY KEY,
    user_id       UUID,
    creation_time TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT    fk_user_sessions FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT    uniq_userid UNIQUE(user_id)
);

CREATE TABLE player_ratings(
    id       SERIAL PRIMARY KEY,
    mmr      INTEGER,
    k_factor INTEGER
);

CREATE TABLE player_statistics(
    id               SERIAL PRIMARY KEY,
    score            INTEGER,
    score_per_minute DECIMAL(2),
    wins             INTEGER,
    losses           INTEGER,
    winrate          DECIMAL(2),
    wins_as_top_10   INTEGER,
    wins_as_top_5    INTEGER,
    wins_as_top_3    INTEGER,
    wins_as_top_1    INTEGER
);

CREATE TABLE player_profiles(
    id                   SERIAL PRIMARY KEY,
    user_id              UUID,
    playtime             INTEGER,
    player_rating_id     INTEGER,
    player_statistics_id INTEGER,
    last_update          TIMESTAMP,
    CONSTRAINT           fk_player_rating        FOREIGN KEY (player_rating_id)     REFERENCES player_ratings(id),
    CONSTRAINT           fk_player_statistics    FOREIGN KEY (player_statistics_id) REFERENCES player_statistics(id),
    CONSTRAINT           fk_user_player_profiles FOREIGN KEY (user_id)              REFERENCES users(id)
);

CREATE TABLE match_records(
    id            UUID PRIMARY KEY,
    user_id       UUID,
    win           BOOLEAN,
    win_kind      INTEGER,
    score         INTEGER,
    length        INTEGER,
    start         TIMESTAMP,
    rating_change INTEGER,
    CONSTRAINT    fk_user_match_records FOREIGN KEY (user_id) REFERENCES users(id)
);
