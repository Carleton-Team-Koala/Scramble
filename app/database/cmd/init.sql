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

CREATE ROLE ReadOnlyUser 
WITH LOGIN NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION 
PASSWORD 'Kwyjibo';


GRANT ALL PRIVILEGES ON DATABASE scramble_db to KoalaAdmin;

GRANT CONNECT ON DATABASE scramble_db TO ReadOnlyUser;
GRANT USAGE ON SCHEMA public TO ReadOnlyUser;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO ReadOnlyUser;



ALTER TABLE games
ADD COLUMN IF NOT EXISTS Player1Score INT;
ADD COLUMN IF NOT EXISTS Player2Score INT;



WITH PlayerScores AS (
    SELECT
        GameID,
        (Players->0->>'name') AS Player1Name,
        (Players->0->>'score')::INT AS Player1Score,
        (Players->1->>'name') AS Player2Name,
        (Players->1->>'score')::INT AS Player2Score
    FROM games
)

SELECT
    PlayerName,
    SUM(PlayerScore) AS TotalScore
FROM (
    SELECT
        GameID,
        Player1Name AS PlayerName,
        Player1Score AS PlayerScore
    FROM PlayerScores

    UNION ALL

    SELECT
        GameID,
        Player2Name AS PlayerName,
        Player2Score AS PlayerScore
    FROM PlayerScores
) AS AllPlayerScores
GROUP BY PlayerName
ORDER BY TotalScore DESC;