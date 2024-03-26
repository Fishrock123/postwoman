CREATE TABLE request (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    url TEXT NOT NULL,
    method TEXT NOT NULL,
    headers TEXT,
    body TEXT,
    status TEXT NOT NULL,
    date TEXT NOT NULL,
    deleted BOOLEAN NOT NULL
);

CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    history INTEGER[],
    favorites INTEGER[],
    deleted BOOLEAN NOT NULL
);

ALTER TABLE request
ADD CONSTRAINT fk_request_user_id
FOREIGN KEY (user_id)
REFERENCES "user"(id);