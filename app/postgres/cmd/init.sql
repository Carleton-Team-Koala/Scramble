CREATE TABLE IF NOT EXISTS games (
    game_id UUID PRIMARY KEY,
    board JSONB,
    available_letters JSONB,
    players JSONB
);

CREATE ROLE KoalaAdmin WITH LOGIN PASSWORD 'KoalatyProduct';
ALTER ROLE KoalaAdmin CREATEDB;