// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: mlb_games.sql

package modelstore

import (
	"context"
	"database/sql"
)

const mLBCreateGame = `-- name: MLBCreateGame :execlastid
INSERT INTO mlb_games
  (gametime, game_id, description, status, link, home_id, visitor_id, home_score, visitor_score, updated_at)
  VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type MLBCreateGameParams struct {
	Gametime     sql.NullTime
	GameID       int64
	Description  sql.NullString
	Status       sql.NullString
	Link         sql.NullString
	HomeID       int64
	VisitorID    int64
	HomeScore    sql.NullInt64
	VisitorScore sql.NullInt64
	UpdatedAt    sql.NullTime
}

func (q *Queries) MLBCreateGame(ctx context.Context, arg MLBCreateGameParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, mLBCreateGame,
		arg.Gametime,
		arg.GameID,
		arg.Description,
		arg.Status,
		arg.Link,
		arg.HomeID,
		arg.VisitorID,
		arg.HomeScore,
		arg.VisitorScore,
		arg.UpdatedAt,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const mLBDeleteGames = `-- name: MLBDeleteGames :exec
DELETE FROM mlb_games
`

func (q *Queries) MLBDeleteGames(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, mLBDeleteGames)
	return err
}

const mLBFindGameByID = `-- name: MLBFindGameByID :one
SELECT mg.id, mg.game_id, mg.gametime, mg.description, mg.status, mg.link, mg.home_score, mg.visitor_score,
  ht.id AS homeID, ht.team_id AS homeTeamID, ht.name AS homeName,
  at.id AS awayID, at.team_id AS awayTeamID, at.name AS awayName
  FROM mlb_games AS mg
  INNER JOIN mlb_teams AS ht ON (mg.home_id = ht.id)
  INNER JOIN mlb_teams AS at ON (mg.visitor_id = at.id)
  WHERE game_id = ?
`

type MLBFindGameByIDRow struct {
	ID           int64
	GameID       int64
	Gametime     sql.NullTime
	Description  sql.NullString
	Status       sql.NullString
	Link         sql.NullString
	HomeScore    sql.NullInt64
	VisitorScore sql.NullInt64
	Homeid       int64
	Hometeamid   int64
	Homename     sql.NullString
	Awayid       int64
	Awayteamid   int64
	Awayname     sql.NullString
}

func (q *Queries) MLBFindGameByID(ctx context.Context, gameID int64) (MLBFindGameByIDRow, error) {
	row := q.db.QueryRowContext(ctx, mLBFindGameByID, gameID)
	var i MLBFindGameByIDRow
	err := row.Scan(
		&i.ID,
		&i.GameID,
		&i.Gametime,
		&i.Description,
		&i.Status,
		&i.Link,
		&i.HomeScore,
		&i.VisitorScore,
		&i.Homeid,
		&i.Hometeamid,
		&i.Homename,
		&i.Awayid,
		&i.Awayteamid,
		&i.Awayname,
	)
	return i, err
}

const mLBLoadGamesByDate = `-- name: MLBLoadGamesByDate :many
SELECT mg.id, mg.game_id, mg.gametime, mg.description, mg.status, mg.home_score, mg.visitor_score,
  ht.id AS homeID, ht.team_id AS homeTeamID, ht.name AS homeName,
  at.id AS awayID, at.team_id AS awayTeamID, at.name AS awayName
  FROM mlb_games AS mg
  INNER JOIN mlb_teams AS ht ON (ng.home_id = ht.id)
  INNER JOIN mlb_teams AS at ON (ng.visitor_id = at.id)
  WHERE mg.gametime BETWEEN ? AND ?
  ORDER BY mg.gametime
`

type MLBLoadGamesByDateParams struct {
	FromGametime sql.NullTime
	ToGametime   sql.NullTime
}

type MLBLoadGamesByDateRow struct {
	ID           int64
	GameID       int64
	Gametime     sql.NullTime
	Description  sql.NullString
	Status       sql.NullString
	HomeScore    sql.NullInt64
	VisitorScore sql.NullInt64
	Homeid       int64
	Hometeamid   int64
	Homename     sql.NullString
	Awayid       int64
	Awayteamid   int64
	Awayname     sql.NullString
}

func (q *Queries) MLBLoadGamesByDate(ctx context.Context, arg MLBLoadGamesByDateParams) ([]MLBLoadGamesByDateRow, error) {
	rows, err := q.db.QueryContext(ctx, mLBLoadGamesByDate, arg.FromGametime, arg.ToGametime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MLBLoadGamesByDateRow
	for rows.Next() {
		var i MLBLoadGamesByDateRow
		if err := rows.Scan(
			&i.ID,
			&i.GameID,
			&i.Gametime,
			&i.Description,
			&i.Status,
			&i.HomeScore,
			&i.VisitorScore,
			&i.Homeid,
			&i.Hometeamid,
			&i.Homename,
			&i.Awayid,
			&i.Awayteamid,
			&i.Awayname,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
