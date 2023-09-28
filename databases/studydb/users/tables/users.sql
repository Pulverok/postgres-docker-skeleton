CREATE TABLE users.users (
    id bigserial NOT NULL,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT now()
);
