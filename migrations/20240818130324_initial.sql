-- +goose Up
CREATE TABLE chats (
    id serial PRIMARY KEY NOT NULL,
    usernames text[]
);

-- +goose Down
DROP TABLE chats;
