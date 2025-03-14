// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: nhl_teams.sql

package modelstore

import (
	"context"
	"database/sql"
)

const nHLCreateTeam = `-- name: NHLCreateTeam :execlastid
INSERT IGNORE INTO nhl_teams
  (team_id, name, link, tricode, franchise_id, updated_at)
VALUES
  (?, ?, ?, ?, ?, ?)
`

type NHLCreateTeamParams struct {
	TeamID      int64
	Name        sql.NullString
	Link        sql.NullString
	Tricode     sql.NullString
	FranchiseID int64
	UpdatedAt   sql.NullTime
}

func (q *Queries) NHLCreateTeam(ctx context.Context, arg NHLCreateTeamParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, nHLCreateTeam,
		arg.TeamID,
		arg.Name,
		arg.Link,
		arg.Tricode,
		arg.FranchiseID,
		arg.UpdatedAt,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const nHLFindTeamByID = `-- name: NHLFindTeamByID :one
SELECT id, updated_at, team_id, franchise_id, tricode, name, link FROM nhl_teams
  WHERE team_id = ?
`

func (q *Queries) NHLFindTeamByID(ctx context.Context, teamID int64) (NhlTeam, error) {
	row := q.db.QueryRowContext(ctx, nHLFindTeamByID, teamID)
	var i NhlTeam
	err := row.Scan(
		&i.ID,
		&i.UpdatedAt,
		&i.TeamID,
		&i.FranchiseID,
		&i.Tricode,
		&i.Name,
		&i.Link,
	)
	return i, err
}
