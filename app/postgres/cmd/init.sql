CREATE TABLE IF NOT EXISTS games (
    game_id UUID PRIMARY KEY,
    board character varying[][] NOT NULL,
    available_letters jsonb NOT NULL
    players JSONB
);

CREATE ROLE KoalaAdmin WITH LOGIN PASSWORD 'KoalatyProduct';
ALTER ROLE KoalaAdmin CREATEDB;