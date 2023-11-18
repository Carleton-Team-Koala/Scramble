\c scramble_db;

-- Creates a table named "games" if it does not already exist.
-- The table has the following columns:
-- - GameID: Text data type
-- - Board: JSONB data type, not nullable
-- - LetterDistribution: JSONB data type, not nullable
-- - Players: JSONB data type
-- - CurrentPlayer: Text data type
-- - PlayerList: JSONB data type
-- - TotalMoves: Integer data type
-- - GameStarted: Boolean data type
CREATE TABLE IF NOT EXISTS games (
    GameID TEXT,
    Board JSONB NOT NULL,
    LetterDistribution JSONB NOT NULL,
    Players JSONB,
    CurrentPlayer TEXT,
    PlayerList JSONB,
    TotalMoves INTEGER,
    GameStarted BOOLEAN
);

/*
This script initializes the database by creating roles, granting privileges, and adding columns to the 'games' table.
*/

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

ALTER TABLE games
ADD COLUMN IF NOT EXISTS Player2Score INT;


/*
    This SQL query calculates the total score for each player in a game.
    It retrieves the player names and scores from the 'games' table and uses a CTE (Common Table Expression) to organize the data.
    The CTE 'PlayerScores' extracts the player names and scores from the 'games' table using JSON functions.
    The main query then combines the player names and scores from the CTE using UNION ALL, and calculates the total score for each player using the SUM function.
    The result is grouped by player name and ordered by the total score in descending order.
*/

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