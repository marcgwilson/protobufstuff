package sqlite

var CreateTables = `CREATE TABLE IF NOT EXISTS resource (
    id INTEGER PRIMARY KEY ASC,
    url VARCHAR NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS request (
    id INTEGER PRIMARY KEY ASC,
    resource_id INTEGER REFERENCES resource(id) ON DELETE CASCADE NOT NULL,
    url VARCHAR NOT NULL,
    headers VARCHAR,
    body VARCHAR,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS response (
    id INTEGER PRIMARY KEY ASC,
    request_id INTEGER REFERENCES request(id) ON DELETE CASCADE NOT NULL,
    headers VARCHAR,
    body VARCHAR,
    temporal VARCHAR,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`

var CreateTriggers = `CREATE TRIGGER set_updated_at_timestamp_resource
AFTER UPDATE On resource
BEGIN
   UPDATE resource SET updated_at = STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW') WHERE id = NEW.id;
END;

CREATE TRIGGER set_updated_at_timestamp_request
AFTER UPDATE On request
BEGIN
   UPDATE request SET updated_at = STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW') WHERE id = NEW.id;
END;

CREATE TRIGGER set_updated_at_timestamp_response
AFTER UPDATE On response
BEGIN
   UPDATE response SET updated_at = STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW') WHERE id = NEW.id;
END;`

var CreateExtras = `CREATE OR REPLACE FUNCTION set_updated_at_timestamp()
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
