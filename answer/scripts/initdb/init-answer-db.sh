#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER answer_user WITH PASSWORD '$POSTGRES_ANSWER_USER_PASSWORD';
	CREATE DATABASE answer;

	ALTER DATABASE answer OWNER TO answer_user;
	GRANT ALL PRIVILEGES ON DATABASE answer TO answer_user;
EOSQL
