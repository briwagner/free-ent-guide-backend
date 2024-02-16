package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/mlbapi"
	"log"
	"time"
)

type MLBGame struct {
	ID           uint      `json:"-"`
	GameID       int       `json:"id"`
	Gametime     time.Time `json:"gametime"`
	Description  string    `json:"description"`
	Status       string    `json:"status"`
	Inning       int       `json:"inning,omitempty"`
	Link         string    `json:"link,omitempty"`
	HomeID       uint      `json:"-"`
	Home         MLBTeam   `json:"home"`
	HomeScore    int       `json:"home_score"`
	VisitorID    uint      `json:"-"`
	Visitor      MLBTeam   `json:"visitor"`
	VisitorScore int       `json:"visitor_score"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

func (g *MLBGame) Create(q *modelstore.Queries) error {
	if g.GameID == 0 {
		return errors.New("invalid game ID")
	}
	if g.HomeID == 0 {
		return errors.New("invalid hometeam ID")
	}
	if g.VisitorID == 0 {
		return errors.New("invalid visitor team ID")
	}
	if g.UpdatedAt.IsZero() {
		g.UpdatedAt = time.Now()
	}

	id, err := q.MLBCreateGame(context.Background(), modelstore.MLBCreateGameParams{
		Gametime:     sql.NullTime{Time: g.Gametime, Valid: true},
		GameID:       int64(g.GameID),
		Description:  sql.NullString{String: g.Description, Valid: true},
		Status:       sql.NullString{String: g.Status, Valid: true},
		Link:         sql.NullString{String: g.Link, Valid: true},
		HomeID:       int64(g.HomeID),
		VisitorID:    int64(g.VisitorID),
		HomeScore:    sql.NullInt64{Int64: int64(g.HomeScore), Valid: true},
		VisitorScore: sql.NullInt64{Int64: int64(g.VisitorScore), Valid: true},
		UpdatedAt:    sql.NullTime{Time: g.UpdatedAt, Valid: true},
	})
	if err != nil {
		return err
	}

	g.ID = uint(id)
	return nil
}

func (mg *MLBGame) FindByGameID(q *modelstore.Queries, gameID int) error {
	row, err := q.MLBFindGameByID(context.Background(), int64(gameID))
	if err != nil {
		return err
	}

	mg.FromDB(row)
	return nil
}

func (mg *MLBGame) FromDB(row modelstore.MLBFindGameByIDRow) {
	mg.ID = uint(row.ID)
	mg.GameID = int(row.GameID)
	mg.Gametime = row.Gametime.Time
	mg.Description = row.Description.String
	mg.Status = row.Status.String
	mg.Link = row.Link.String
	mg.HomeID = uint(row.Homeid)
	mg.Home = MLBTeam{
		ID:     uint(row.Homeid),
		TeamID: int(row.Hometeamid),
		Name:   row.Homename.String,
	}
	mg.HomeScore = int(row.HomeScore.Int64)

	mg.VisitorID = uint(row.Awayid)
	mg.Visitor = MLBTeam{
		ID:     uint(row.Awayid),
		TeamID: int(row.Awayteamid),
		Name:   row.Awayname.String,
	}
	mg.VisitorScore = int(row.VisitorScore.Int64)
}

type MLBTeam struct {
	ID        uint      `json:"-"`
	TeamID    int       `json:"id"`
	Name      string    `json:"name"`
	Link      string    `json:"link,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (mt *MLBTeam) Create(q *modelstore.Queries) error {
	if mt.UpdatedAt.IsZero() {
		mt.UpdatedAt = time.Now()
	}
	id, err := q.MLBCreateTeam(context.Background(), modelstore.MLBCreateTeamParams{
		TeamID:    int64(mt.TeamID),
		Name:      sql.NullString{String: mt.Name, Valid: true},
		Link:      sql.NullString{String: mt.Link, Valid: true},
		UpdatedAt: sql.NullTime{Time: mt.UpdatedAt, Valid: true},
	})
	if err != nil {
		return err
	}

	mt.ID = uint(id)
	return nil
}

func (mt *MLBTeam) FindByTeamID(q *modelstore.Queries, team_id int) error {
	if team_id == 0 {
		return errors.New("invalid team id")
	}

	team, err := q.MLBFindTeamByID(context.Background(), int64(team_id))
	if err != nil {
		return err
	}

	mt.ID = uint(team.ID)
	mt.TeamID = int(team.TeamID)
	mt.Name = team.Name.String
	mt.Link = team.Link.String
	mt.UpdatedAt = team.UpdatedAt.Time

	return nil
}

type MLBGames []MLBGame

// LoadByDate fetches all games for the given date.
func (mgs *MLBGames) LoadByDate(d string, db Store) error {
	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		return err
	}
	// Set date to noon to allow capturing anything in PM that may show as next-day with UTC.
	dateFrom := time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, date.Location())
	dateTo := dateFrom.Add(time.Hour * 24)
	tx := db.Preload("Home").Preload("Visitor").Where("gametime BETWEEN ? AND ?", dateFrom.String(), dateTo.String()).Order("gametime").Find(mgs)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindByID searches for a game by Game ID.
func (g *MLBGame) FindByID(id string, db Store) error {
	tx := db.Preload("Home").Preload("Visitor").Where("game_id = ?", id).Find(g)
	if tx.Error != nil {
		return tx.Error
	}
	if g.ID == 0 {
		return errors.New("MLB game not found")
	}
	return nil
}

// ImportMLB calls to MLB api and saves new games to the DB.
func ImportMLB(q *modelstore.Queries, startDate time.Time) error {
	gameweek, err := mlbapi.ImportDates(startDate)
	if err != nil {
		return err
	}

	var count, countErrs int

	for _, day := range gameweek.Dates {
		log.Printf("importing MLB games for %s\n", day.Date)

		for _, g := range day.Games {
			game := MLBGame{
				GameID:      g.GameID,
				Gametime:    g.Time,
				Description: g.Description,
				Status:      g.Status.Status,
				Link:        g.Link,
			}

			var home MLBTeam
			err := home.FindByTeamID(q, g.Teams.Home.ID)
			if err != nil {
				// TODO create custom error and check for "team not found"

				// Try to create team.
				log.Printf("mlb creating home team %s\n", g.Teams.Home.Name)
				home.Name = g.Teams.Home.Name
				home.TeamID = g.Teams.Home.ID
				home.Link = g.Teams.Home.Link
				err = home.Create(q)

				if err != nil {
					log.Printf("mlb error finding home team %d\n", g.Teams.Home.ID)
					countErrs++
					continue
				}
			}
			game.HomeID = home.ID
			game.HomeScore = g.Teams.Home.Score

			var away MLBTeam
			err = away.FindByTeamID(q, g.Teams.Away.ID)
			if err != nil {
				// TODO create custom error and check for "team not found"

				// Try to create team.
				log.Printf("mlb creating away team %s\n", g.Teams.Away.Name)
				away.Name = g.Teams.Away.Name
				away.TeamID = g.Teams.Away.ID
				away.Link = g.Teams.Away.Link
				err = away.Create(q)

				if err != nil {
					log.Printf("mlb error finding away team %d\n", g.Teams.Away.ID)
					countErrs++
					continue
				}
			}
			game.VisitorID = away.ID
			game.VisitorScore = g.Teams.Away.Score

			err = game.Create(q)
			if err != nil {
				log.Printf("mlb error creating game %s", err)
				countErrs++
				continue
			}

			count++
		}
	}

	log.Printf("mlb import complete, adding %d games. Errors: %d", count, countErrs)

	return nil
}

// Use in Game update cycle, when checking playoff games that are not needed. The ID will be
// zero-d out, so we should delete the game.
var ErrorGameCanceled = errors.New("game is zeroed")

// GetUpdate makes an api request to get game update to
// merge with scheduled game info from the database. Called from handler.
func (g *MLBGame) GetUpdate() error {
	if g.GameID == 0 {
		return fmt.Errorf("invalid game id %d", g.GameID)
	}

	gameup, err := mlbapi.GetGameUpdate(g.Link)
	if err != nil {
		return err
	}

	g.HomeScore = gameup.HomeScore
	g.VisitorScore = gameup.VisitorScore
	g.Status = gameup.Status
	g.Inning = gameup.Inning

	return nil
}
