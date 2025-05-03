-- name: MLBCreateGame :execlastid
INSERT INTO mlb_games
  (gametime, game_id, description, status, link, home_id, visitor_id, home_score, visitor_score, updated_at)
  VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: MLBFindGameByID :one
SELECT mg.id, mg.game_id, mg.gametime, mg.description, mg.status, mg.link, mg.home_score, mg.visitor_score,
  ht.id AS homeID, ht.team_id AS homeTeamID, ht.name AS homeName,
  at.id AS awayID, at.team_id AS awayTeamID, at.name AS awayName
  FROM mlb_games AS mg
  INNER JOIN mlb_teams AS ht ON (mg.home_id = ht.id)
  INNER JOIN mlb_teams AS at ON (mg.visitor_id = at.id)
  WHERE game_id = ?;

-- name: MLBUpdateScore :exec
UPDATE mlb_games
  SET home_score = ?, visitor_score = ?, status = ?
  WHERE game_id = ?;

-- name: MLBLoadGamesByDate :many
SELECT mg.id, mg.game_id, mg.gametime, mg.description, mg.status, mg.link, mg.home_score, mg.visitor_score,
		ht.id AS homeID, ht.team_id AS homeTeamID, ht.name AS homeName,
		at.id AS awayID, at.team_id AS awayTeamID, at.name AS awayName
	FROM mlb_games AS mg
	INNER JOIN mlb_teams AS ht ON (mg.home_id = ht.id)
	INNER JOIN mlb_teams AS at ON (mg.visitor_id = at.id)
	WHERE mg.gametime BETWEEN ? AND ?
	ORDER BY mg.gametime;

-- name: MLBLoadIncompleteGamesByDate :many
SELECT mg.id, mg.game_id, mg.gametime, mg.description, mg.status, mg.link, mg.home_score, mg.visitor_score,
		ht.id AS homeID, ht.team_id AS homeTeamID, ht.name AS homeName,
		at.id AS awayID, at.team_id AS awayTeamID, at.name AS awayName
	FROM mlb_games AS mg
	INNER JOIN mlb_teams AS ht ON (mg.home_id = ht.id)
	INNER JOIN mlb_teams AS at ON (mg.visitor_id = at.id)
	WHERE mg.gametime BETWEEN ? AND ?
	AND (mg.status NOT IN ("Final","Postponed","Completed Early"))
	ORDER BY mg.gametime;

-- name: MLBLatestGames :many
SELECT id, gametime
  FROM mlb_games
  WHERE DATE(gametime) = (
    SELECT DATE(gametime)
    FROM mlb_games
    ORDER BY gametime
    DESC LIMIT 1
  );

-- name: MLBDeleteGames :exec
DELETE FROM mlb_games;

-- name: MLBNextGamesByTeam :many
SELECT sqlc.embed(mlb_games),
		ht.id AS homeID, ht.team_id AS homeTeamID, ht.name AS homeName,
		at.id AS awayID, at.team_id AS awayTeamID, at.name AS awayName
	FROM mlb_games
	INNER JOIN mlb_teams AS ht ON (mlb_games.home_id = ht.id)
	INNER JOIN mlb_teams AS at ON (mlb_games.visitor_id = at.id)
	WHERE mlb_games.gametime >= sqlc.arg(gametime)
	AND (at.team_id = sqlc.arg(team_id) OR ht.team_id = sqlc.arg(team_id))
	ORDER BY mlb_games.gametime
	LIMIT 5;

-- name: MLBLastGamesByTeam :many
SELECT sqlc.embed(mlb_games),
		ht.id AS homeID, ht.team_id AS homeTeamID, ht.name AS homeName,
		at.id AS awayID, at.team_id AS awayTeamID, at.name AS awayName
	FROM mlb_games
	INNER JOIN mlb_teams AS ht ON (mlb_games.home_id = ht.id)
	INNER JOIN mlb_teams AS at ON (mlb_games.visitor_id = at.id)
	WHERE mlb_games.gametime < sqlc.arg(gametime)
	AND (at.team_id = sqlc.arg(team_id) OR ht.team_id = sqlc.arg(team_id))
	ORDER BY mlb_games.gametime DESC
	LIMIT 5;