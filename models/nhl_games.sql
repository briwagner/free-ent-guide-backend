-- name: NHLCreateGame :execlastid
INSERT INTO nhl_games
  (gametime, game_id, description, status, link, home_id, visitor_id, home_score, visitor_score, updated_at)
  VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: NHLFindGameByID :one
SELECT ng.id, ng.game_id, ng.gametime, ng.description, ng.status, ng.home_score, ng.visitor_score,
  ht.id AS homeID, ht.team_id AS homeTeamID, ht.name AS homeName,
  at.id AS awayID, at.team_id AS awayTeamID, at.name AS awayName
  FROM nhl_games AS ng
  INNER JOIN nhl_teams AS ht ON (ng.home_id = ht.id)
  INNER JOIN nhl_teams AS at ON (ng.visitor_id = at.id)
  WHERE game_id = ?;

-- name: NHLUpdateScore :exec
UPDATE nhl_games
  SET home_score = ?, visitor_score = ?, status = ?
  WHERE game_id = ?;

-- name: NHLLoadGamesByDate :many
SELECT ng.id, ng.game_id, ng.gametime, ng.description, ng.status, ng.home_score, ng.visitor_score,
  ht.id AS homeID, ht.team_id AS homeTeamID, ht.name AS homeName,
  at.id AS awayID, at.team_id AS awayTeamID, at.name AS awayName
  FROM nhl_games AS ng
  INNER JOIN nhl_teams AS ht ON (ng.home_id = ht.id)
  INNER JOIN nhl_teams AS at ON (ng.visitor_id = at.id)
  WHERE ng.gametime BETWEEN ? AND ?
  ORDER BY ng.gametime;

-- name: NHLLatestGames :many
SELECT id, gametime
  FROM nhl_games
  WHERE DATE(gametime) = (
    SELECT DATE(gametime)
    FROM nhl_games
    ORDER BY gametime
    DESC LIMIT 1
  );