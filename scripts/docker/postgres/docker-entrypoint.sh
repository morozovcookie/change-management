#!/bin/sh

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER controller WITH PASSWORD 'controller';
	CREATE DATABASE cm;
	GRANT ALL PRIVILEGES ON DATABASE cm TO controller;
EOSQL


psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "cm" <<-EOSQL
	CREATE SCHEMA controller;
EOSQL
