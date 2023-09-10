CREATE TABLE users
(
    "id" SERIAL PRIMARY KEY,
    "name" varchar(255) not null,
    "username" varchar(255) not null unique,
    "password_hash" varchar(255) not null
);

CREATE TABLE todo_lists
(
    "id" SERIAL PRIMARY KEY,
    "title" varchar(255) not null,
    "description" varchar(255)
);

CREATE TABLE users_lists
(
    "id" SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    list_id INT REFERENCES todo_lists (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE todo_items
(
    "id" SERIAL PRIMARY KEY,
    "title" varchar(255) not null,
    "description" varchar(255),
    "done" boolean not null default false
);

CREATE TABLE lists_items
(
    "id" SERIAL PRIMARY KEY,
    list_id INT REFERENCES todo_lists (id) ON DELETE CASCADE NOT NULL,
    item_id INT REFERENCES todo_items (id) ON DELETE CASCADE NOT NULL
);
