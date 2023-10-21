\c scramble_db;

CREATE TABLE IF NOT EXISTS games (
    GameID int,
    Board text[15][15] NOT NULL,
    LetterDistribution jsonb NOT NULL,
    Players JSONB
);

CREATE ROLE KoalaAdmin 
LOGIN
PASSWORD 'KoalatyProduct';

GRANT ALL PRIVILEGES ON DATABASE scramble_db to KoalaAdmin;