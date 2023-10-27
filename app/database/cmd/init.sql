\c scramble_db;

CREATE TABLE IF NOT EXISTS games (
    GameID TEXT,
    Board JSONB NOT NULL,
    LetterDistribution JSONB NOT NULL,
    Players JSONB
);

CREATE ROLE KoalaAdmin 
LOGIN
PASSWORD 'KoalatyProduct';

GRANT ALL PRIVILEGES ON DATABASE scramble_db to KoalaAdmin;