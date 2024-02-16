-- name: MLBCreateTeam :execlastid
INSERT IGNORE INTO mlb_teams
  (team_id, name, link, updated_at)
VALUES
  (?, ?, ?, ?);

-- name: MLBFindTeamByID :one
SELECT * FROM mlb_teams
  WHERE team_id = ?;

-- name: MLBDeleteTeams :exec
DELETE FROM mlb_teams;