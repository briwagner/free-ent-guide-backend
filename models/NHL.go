package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/nhlapi"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

type NHLTeam struct {
	ID          uint      `json:"-"`
	TeamID      int       `json:"id"`
	FranchiseID int       `json:"franchiseId,omitempty"`
	Tricode     string    `json:"triCode,omitempty"`
	Name        string    `json:"fullName"`
	Link        string    `json:"link,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func (nt *NHLTeam) Create(q *modelstore.Queries) error {
	if nt.TeamID == 0 {
		return errors.New("invalid team ID")
	}
	if nt.UpdatedAt.IsZero() {
		nt.UpdatedAt = time.Now()
	}
	id, err := q.NHLCreateTeam(context.Background(), modelstore.NHLCreateTeamParams{
		Name:        sql.NullString{String: nt.Name, Valid: true},
		Link:        sql.NullString{String: nt.Link, Valid: true},
		TeamID:      int64(nt.TeamID),
		FranchiseID: int64(nt.FranchiseID),
		Tricode:     sql.NullString{String: nt.Tricode, Valid: true},
		UpdatedAt:   sql.NullTime{Time: nt.UpdatedAt, Valid: true},
	})
	if err != nil {
		return err
	}
	nt.ID = uint(id)
	return nil
}

func (nt *NHLTeam) FindByTeamID(q *modelstore.Queries, teamID int64) error {
	if teamID == 0 {
		return errors.New("invalid team ID")
	}
	res, err := q.NHLFindTeamByID(context.Background(), teamID)
	if err != nil {
		return err
	}
	nt.FromDB(res)
	return nil
}

func (nt *NHLTeam) FindOrCreateByTeamID(q *modelstore.Queries, teamID int64) error {
	err := nt.FindByTeamID(q, teamID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		// Create
		idStr := strconv.FormatInt(teamID, 10)
		newteam, err := nhlapi.GetTeam(idStr)
		if err != nil {
			return err
		}

		nt.TeamID = newteam.ID
		nt.Name = newteam.Name
		nt.Tricode = newteam.Tricode

		err = nt.Create(q)
		if err != nil {
			return err
		}
	}

	return nil
}

// FromDB converts types to return a team we use for the API.
func (nt *NHLTeam) FromDB(team modelstore.NhlTeam) {
	nt.ID = uint(team.ID)
	nt.TeamID = int(team.TeamID)
	nt.Name = team.Name.String
	nt.Link = team.Link.String
	nt.Tricode = team.Tricode.String
}

// This is sent back to frontend.
type NHLGameUpdate struct {
	GameID       int     `json:"ID"`
	Game         NHLGame `json:"Game"`
	Status       string  `json:"Status"`
	Period       int     `json:"Period"`
	HomeScore    int     `json:"HomeScore"`
	VisitorScore int     `json:"VisitorScore"`
}

func (ngu *NHLGameUpdate) FromGame(ng NHLGame) {
	ngu.GameID = ng.GameID
	ngu.Game = ng
	ngu.Status = ng.Status
	ngu.Period = ng.Period
	ngu.HomeScore = ng.HomeScore
	ngu.VisitorScore = ng.VisitorScore
}

type NHLGame struct {
	ID           uint
	GameID       int       `json:"id"`
	Gametime     time.Time `json:"gametime"`
	Description  string    `json:"description"`
	Status       string    `json:"status"`
	Period       int       `json:"period,omitempty"`
	Link         string    `json:"link,omitempty"`
	HomeID       uint      `json:"-"`
	Home         *NHLTeam  `json:"home"`
	HomeScore    int       `json:"home_score"`
	VisitorID    uint      `json:"-"`
	Visitor      *NHLTeam  `json:"visitor"`
	VisitorScore int       `json:"visitor_score"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

func (ng *NHLGame) Create(q *modelstore.Queries) error {
	if ng.GameID == 0 {
		return errors.New("invalid game ID")
	}
	if ng.UpdatedAt.IsZero() {
		ng.UpdatedAt = time.Now()
	}
	id, err := q.NHLCreateGame(context.Background(), modelstore.NHLCreateGameParams{
		GameID:       sql.NullInt64{Int64: int64(ng.GameID), Valid: true},
		Gametime:     sql.NullTime{Time: ng.Gametime, Valid: true},
		Description:  sql.NullString{String: ng.Description, Valid: true},
		Status:       sql.NullString{String: ng.Status, Valid: true},
		Link:         sql.NullString{String: ng.Link, Valid: true},
		HomeID:       int64(ng.HomeID),
		VisitorID:    int64(ng.VisitorID),
		HomeScore:    sql.NullInt64{Int64: int64(ng.HomeScore), Valid: true},
		VisitorScore: sql.NullInt64{Int64: int64(ng.VisitorScore), Valid: true},
		UpdatedAt:    sql.NullTime{Time: ng.UpdatedAt, Valid: true},
	})
	if err != nil {
		return err
	}
	ng.ID = uint(id)
	return nil
}

// FindByID searches for a game by Game ID.
func (g *NHLGame) FindByGameID(q *modelstore.Queries, id int) error {
	if id == 0 {
		return errors.New("invalid game ID")
	}
	res, err := q.NHLFindGameByID(context.Background(), sql.NullInt64{Int64: int64(id), Valid: true})
	if err != nil {
		return err
	}

	g.ID = uint(res.ID)
	g.GameID = int(res.GameID.Int64)
	g.Gametime = res.Gametime.Time
	g.Description = res.Description.String
	g.Status = res.Status.String
	g.HomeID = uint(res.Homeid)
	g.VisitorID = uint(res.Awayid)
	g.Home = &NHLTeam{ID: uint(res.Homeid), Name: res.Homename.String}
	g.Visitor = &NHLTeam{ID: uint(res.Awayid), Name: res.Awayname.String}
	return nil
}

// UpdateScorev2 saves to the DB.
func (g *NHLGame) UpdateScorev2(q *modelstore.Queries) error {
	if g.GameID == 0 {
		return errors.New("invalid game ID")
	}
	return q.NHLUpdateScore(context.Background(), modelstore.NHLUpdateScoreParams{
		HomeScore:    sql.NullInt64{Int64: int64(g.HomeScore), Valid: true},
		VisitorScore: sql.NullInt64{Int64: int64(g.VisitorScore), Valid: true},
		GameID:       sql.NullInt64{Int64: int64(g.GameID), Valid: true},
		Status:       sql.NullString{String: g.Status, Valid: true},
	})
}

// UpdateScore updates a DB record to set the scores for a completed game.
func (g *NHLGame) UpdateScore(q *modelstore.Queries) error {
	up, err := nhlapi.GetUpdate(g.GameID)
	if err != nil {
		return err
	}

	// `Final` is now `OFF`, aka official.
	if up.Status != "Final" {
		return fmt.Errorf("update failed; game not finished %d", up.ID)
	}
	g.HomeScore = int(up.HomeScore)
	g.VisitorScore = int(up.VisitorScore)
	g.Status = up.Status

	return g.UpdateScorev2(q)
}

func (g *NHLGame) FromDB(game modelstore.NhlGame) {
	g.ID = uint(game.ID)
	g.Gametime = game.Gametime.Time
	g.GameID = int(game.GameID.Int64)
	g.Description = game.Description.String
	g.Status = game.Status.String
	g.Link = game.Link.String
	g.HomeID = uint(game.HomeID)
	g.VisitorID = uint(game.VisitorID)
	g.HomeScore = int(game.HomeScore.Int64)
	g.VisitorScore = int(game.VisitorScore.Int64)
}

type NHLGames []NHLGame

// LoadByDate fetches all games for the given date.
func (ngs *NHLGames) LoadByDate(q *modelstore.Queries, datestr string) error {
	date, err := time.Parse("2006-01-02", datestr)
	if err != nil {
		return err
	}
	// Set date to noon to allow capturing anything in PM that may show as next-day with UTC.
	dateFrom := time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, date.Location())
	dateTo := dateFrom.Add(time.Hour * 24)

	rows, err := q.NHLLoadGamesByDate(context.Background(), modelstore.NHLLoadGamesByDateParams{
		FromGametime: sql.NullTime{Time: dateFrom, Valid: true},
		ToGametime:   sql.NullTime{Time: dateTo, Valid: true},
	})
	if err != nil {
		return err
	}

	for _, row := range rows {
		g := NHLGame{
			ID:           uint(row.ID),
			GameID:       int(row.GameID.Int64),
			Gametime:     row.Gametime.Time,
			Description:  row.Description.String,
			Status:       row.Status.String,
			HomeScore:    int(row.HomeScore.Int64),
			VisitorScore: int(row.VisitorScore.Int64),
		}
		home := NHLTeam{
			ID:     uint(row.Homeid),
			TeamID: int(row.Hometeamid),
			Name:   row.Homename.String,
		}
		g.Home = &home
		g.HomeID = home.ID

		away := NHLTeam{
			ID:     uint(row.Awayid),
			TeamID: int(row.Awayteamid),
			Name:   row.Awayname.String,
		}
		g.Visitor = &away
		g.VisitorID = away.ID

		*ngs = append(*ngs, g)
	}
	return nil
}

// LoadByDateIncomplete fetches all games for the given date and not 'Final, OFF' (official).
func (ngs *NHLGames) LoadByDateIncomplete(q *modelstore.Queries, datestr string) error {
	date, err := time.Parse("2006-01-02", datestr)
	if err != nil {
		return err
	}
	// Set date to noon to allow capturing anything in PM that may show as next-day with UTC.
	dateFrom := time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, date.Location())
	dateTo := dateFrom.Add(time.Hour * 24)

	rows, err := q.NHLLoadIncompleteGamesByDate(context.Background(), modelstore.NHLLoadIncompleteGamesByDateParams{
		FromGametime: sql.NullTime{Time: dateFrom, Valid: true},
		ToGametime:   sql.NullTime{Time: dateTo, Valid: true},
	})
	if err != nil {
		return err
	}

	for _, row := range rows {
		g := NHLGame{
			ID:           uint(row.ID),
			GameID:       int(row.GameID.Int64),
			Gametime:     row.Gametime.Time,
			Description:  row.Description.String,
			Status:       row.Status.String,
			HomeScore:    int(row.HomeScore.Int64),
			VisitorScore: int(row.VisitorScore.Int64),
		}
		home := NHLTeam{
			ID:     uint(row.Homeid),
			TeamID: int(row.Hometeamid),
			Name:   row.Homename.String,
		}
		g.Home = &home
		g.HomeID = home.ID

		away := NHLTeam{
			ID:     uint(row.Awayid),
			TeamID: int(row.Awayteamid),
			Name:   row.Awayname.String,
		}
		g.Visitor = &away
		g.VisitorID = away.ID

		*ngs = append(*ngs, g)
	}
	return nil
}

// NHLGetNextGameday tries to find games for the date passed,
// and if not will continue searching on the next day until
// reaching a limit of days searched.
func NHLGetNextGameday(q *modelstore.Queries, date time.Time) (NHLGames, error) {
	tries := 4

	var gs NHLGames
	var err error
	for i := 0; i < tries; i++ {
		err = gs.LoadByDate(q, date.Format("2006-01-02"))
		if err != nil {
			return gs, err
		}
		if len(gs) > 0 {
			return gs, nil
		}
		date = date.Add(24 * time.Hour)
	}

	return gs, err
}

// NHLGetLatestGames loads all games on the latest date found in the DB.
func NHLGetLatestGames(q *modelstore.Queries) (NHLGames, error) {
	var games []NHLGame

	rows, err := q.NHLLatestGames(context.Background())
	if err != nil {
		return games, err
	}

	for _, row := range rows {
		new := NHLGame{ID: uint(row.ID), Gametime: row.Gametime.Time}
		games = append(games, new)
	}

	return games, nil
}

// ImportNHL fetches a week of games from NHL API and converts to struct
// storing them in DB.
func ImportNHL(q *modelstore.Queries, startDate string) (string, error) {
	gameweek, err := nhlapi.ImportNHL(startDate)
	if err != nil {
		return err.Error(), err
	}

	var (
		count, countErrs int
		sb               strings.Builder
	)

	for _, days := range gameweek.Days {
		log.Printf("import %d NHL games for %s\n", days.NumberOfGames, days.Date)
		daysGames := days.Games

		for _, g := range daysGames {
			var game NHLGame
			game.GameID = g.ID
			game.Gametime = g.Gametime
			game.Description = g.GameType.Type
			game.Status = g.GameState

			// Map MLB team IDs to my database IDs.
			home := &NHLTeam{}
			err := home.FindOrCreateByTeamID(q, int64(g.Home.ID))
			if err != nil {
				countErrs++
				log.Printf("error finding home team %d for game %d: %s", g.Home.ID, g.ID, err)
				continue
			}
			game.HomeID = home.ID

			away := &NHLTeam{}
			err = away.FindOrCreateByTeamID(q, int64(g.Away.ID))
			if err != nil {
				countErrs++
				log.Printf("error finding away team %d for game %d: %s", g.Away.ID, g.ID, err)
				continue
			}
			game.VisitorID = away.ID

			err = game.Create(q)
			if err != nil {
				countErrs++
				// Ignore duplicate entry in logs.
				if err.(*mysql.MySQLError).Number != 1062 {
					log.Printf("error creating game %d: %s", g.ID, err)
				}
				log.Printf("some error %s", err)
				continue
			}

			// fmt.Printf("got game %v\n", game.GameID)
			// fmt.Printf("home ID %d | %d\n", game.HomeID, g.Home.ID)
			// fmt.Printf("away ID %d | %d\n", game.VisitorID, g.Away.ID)
			// fmt.Printf("gametime %v\n\n", game.Gametime)

			count++
		}
		sb.WriteString(fmt.Sprintf("nhl import complete for %s, adding %d games. Errors: %d\n", days.Date, count, countErrs))
	}

	return sb.String(), nil
}
