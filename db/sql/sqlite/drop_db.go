package sqlite

var DropDatabase = `DROP DATABASE IF EXISTS $DB_NAME;
DROP USER IF EXISTS $DB_USER;`
