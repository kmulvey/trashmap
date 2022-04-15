DROP TABLE areas;
DROP TABLE users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(128) NOT NULL,
    contact_allowed boolean DEFAULT false
);

CREATE TABLE areas (
    id SERIAL PRIMARY KEY,
    user_id SERIAL,
    polygon GEOMETRY,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE INDEX areas_polygon_idx ON areas USING GIST (polygon);
