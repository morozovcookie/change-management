#!/bin/sh

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER controller WITH PASSWORD 'controller';
	CREATE DATABASE cm;
	GRANT CONNECT ON DATABASE cm TO controller;
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "cm" <<-EOSQL
	CREATE SCHEMA controller;
	GRANT ALL PRIVILEGES ON SCHEMA controller TO controller;
	GRANT DELETE, INSERT, SELECT, UPDATE ON ALL TABLES IN SCHEMA controller TO controller;
	GRANT SELECT, USAGE ON ALL SEQUENCES IN SCHEMA controller TO controller;
EOSQL
