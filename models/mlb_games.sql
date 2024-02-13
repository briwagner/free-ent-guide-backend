-- name: MLBFindGameByID :one
SELECT * FROM mlb_games
  INNER JOIN mlb_teams mlbt ON (
    (mlb_games.home_id = mlbt.id)
    OR
    (mlb_games.visitor_id = mlbt.id)
  )
WHERE mlb_games.id = ? LIMIT 1;

-- name: MLBLoadGamesByDate :many
SELECT * FROM mlb_games
  INNER JOIN mlb_teams mlbt ON (
    (mlb_games.home_id = mlbt.id)
    OR
    (mlb_games.visitor_id = mlbt.id)
  )
WHERE mlb_games.gametime BETWEEN ? AND ?
ORDER BY mlb_games.gametime;
