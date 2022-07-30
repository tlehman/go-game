CREATE TABLE users(
    id         int primary key,
    created_at datetime
    handle     text,
    roles      text
);

CREATE TABLE games(
    id            int primary key,
    created_at    datetime,
    started_at    datetime,
    ended_at      datetime,
    black_user_id int,
    white_user_id int
);

