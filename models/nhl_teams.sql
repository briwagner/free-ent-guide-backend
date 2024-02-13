-- name: NHLCreateTeam :execlastid
INSERT IGNORE INTO nhl_teams
  (team_id, name, link, tricode, franchise_id, updated_at)
VALUES
  (?, ?, ?, ?, ?, ?);

-- name: NHLFindTeamByID :one
SELECT * FROM nhl_teams
  WHERE team_id = ?;