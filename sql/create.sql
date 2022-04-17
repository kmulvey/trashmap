DROP TABLE IF EXISTS areas;
DROP TABLE IF EXISTS auth;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(128) NOT NULL,
    contact_allowed boolean DEFAULT false
);

CREATE TABLE auth (
    auth_token varchar(32) PRIMARY KEY,
    user_id SERIAL,
	created timestamp NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE areas (
    id SERIAL PRIMARY KEY,
    user_id SERIAL,
    polygon GEOMETRY,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE INDEX areas_polygon_idx ON areas USING GIST (polygon);
