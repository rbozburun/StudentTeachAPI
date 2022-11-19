CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY NOT NULL,
    role_name CHARACTER VARYING(45) NOT NULL  
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,
    name CHARACTER VARYING(45) NOT NULL,
    surname CHARACTER VARYING(45) NOT NULL,
    username CHARACTER VARYING(45) NOT NULL,
    email CHARACTER VARYING(45) NOT NULL,
    role_id INTEGER REFERENCES roles(id)
);

