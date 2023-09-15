CREATE TABLE IF NOT EXISTS pessoas (
    id uuid PRIMARY KEY,
    nome varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
);

CREATE EXTENSION IF NOT EXISTS pg_trgm SCHEMA pg_catalog;

