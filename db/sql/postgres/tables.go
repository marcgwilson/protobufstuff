package postgres

var createTables = `CREATE TABLE resource (
    -- id SERIAL PRIMARY KEY,
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    url VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE request (
    -- id SERIAL PRIMARY KEY,
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    resource_id INTEGER REFERENCES resource(id) ON DELETE CASCADE NOT NULL,
    url VARCHAR NOT NULL,
    headers VARCHAR,
    body VARCHAR,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE response (
    -- id SERIAL PRIMARY KEY,
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    request_id INTEGER REFERENCES request(id) ON DELETE CASCADE NOT NULL,
    headers VARCHAR,
    body VARCHAR,
    temporal VARCHAR,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION set_updated_at_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON resource
FOR EACH ROW
EXECUTE PROCEDURE set_updated_at_timestamp();

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON request
FOR EACH ROW
EXECUTE PROCEDURE set_updated_at_timestamp();

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON response
FOR EACH ROW
EXECUTE PROCEDURE set_updated_at_timestamp();

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO $DB_USER;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO $DB_USER;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO $DB_USER;`
