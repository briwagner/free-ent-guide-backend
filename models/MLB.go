package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/mlbapi"
	"log"
	"net/http"
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

func FindByGameID(q *modelstore.Queries, gameID int) (*MLBGame, error) {
	row, err := q.MLBFindGameByID(context.Background(), int64(gameID))
	if err != nil {
		return nil, err
	}

	g := fromDB(row)
	return g, nil
}

// UpdateScore pulls update from MLB api.
func (mg *MLBGame) UpdateScore(q *modelstore.Queries, client *http.Client) error {
	up, err := mlbapi.GetGameUpdate(client, mg.Link)
	if err != nil {
		return err
	}

	if up.Status != "Final" {
		log.Printf("game not finished: %d, %s\n", up.GamePK, up.Status)
		return nil
	}

	mg.Status = up.Status
	mg.HomeScore = up.HomeScore
	mg.VisitorScore = up.VisitorScore
	mg.UpdatedAt = time.Now()

	return mg.UpdateScoreV2(q)
}

func (mg *MLBGame) UpdateScoreV2(q *modelstore.Queries) error {
	if mg.GameID == 0 {
		return errors.New("invalid game ID")
	}

	return q.MLBUpdateScore(context.Background(), modelstore.MLBUpdateScoreParams{
		HomeScore:    sql.NullInt64{Valid: true, Int64: int64(mg.HomeScore)},
		VisitorScore: sql.NullInt64{Valid: true, Int64: int64(mg.VisitorScore)},
		Status:       sql.NullString{Valid: true, String: mg.Status},
		GameID:       int64(mg.GameID),
	})
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
func (mgs *MLBGames) LoadByDate(q *modelstore.Queries, d string) error {
	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		return fmt.Errorf("mlb error parsing date: %w", err)
	}
	// Set date to noon to allow capturing anything in PM that may show as next-day with UTC.
	dateFrom := time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, date.Location())
	dateTo := dateFrom.Add(time.Hour * 24)

	rows, err := q.MLBLoadGamesByDate(context.Background(), modelstore.MLBLoadGamesByDateParams{
		FromGametime: sql.NullTime{Time: dateFrom, Valid: true},
		ToGametime:   sql.NullTime{Time: dateTo, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("error fetching MLB games: %w", err)
	}

	for _, row := range rows {
		*mgs = append(*mgs, MLBGame{
			ID:          uint(row.ID),
			GameID:      int(row.GameID),
			Gametime:    row.Gametime.Time,
			Description: row.Description.String,
			Status:      row.Status.String,
			Link:        row.Link.String,
			HomeID:      uint(row.Homeid),
			Home: MLBTeam{
				ID:     uint(row.Homeid),
				Name:   row.Homename.String,
				TeamID: int(row.Hometeamid),
			},
			VisitorID: uint(row.Awayid),
			Visitor: MLBTeam{
				ID:     uint(row.Awayid),
				Name:   row.Awayname.String,
				TeamID: int(row.Awayteamid),
			},
			HomeScore:    int(row.HomeScore.Int64),
			VisitorScore: int(row.VisitorScore.Int64),
		})
	}

	return nil
}

// ImportMLB calls to MLB api and saves new games to the DB.
func ImportMLB(q *modelstore.Queries, client *http.Client, startDate time.Time) (string, error) {
	gameweek, err := mlbapi.ImportDates(client, startDate)
	if err != nil {
		return "", err
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

	ret := fmt.Sprintf("mlb import complete, adding %d games. Errors: %d", count, countErrs)

	return ret, nil
}

// MLBGetLatestGames loads all games on the latest date found in the DB.
func MLBGetLatestGames(q *modelstore.Queries) (MLBGames, error) {
	var games []MLBGame
	rows, err := q.MLBLatestGames(context.Background())
	if err != nil {
		return games, err
	}

	for _, row := range rows {
		new := MLBGame{ID: uint(row.ID), Gametime: row.Gametime.Time}
		games = append(games, new)
	}

	return games, nil
}

// Use in Game update cycle, when checking playoff games that are not needed. The ID will be
// zero-d out, so we should delete the game.
var ErrorGameCanceled = errors.New("game is zeroed")

// GetUpdate makes an api request to get game update to
// merge with scheduled game info from the database. Called from handler.
func (g *MLBGame) GetUpdate(client *http.Client) error {
	if g.GameID == 0 {
		return fmt.Errorf("invalid game id %d", g.GameID)
	}

	gameup, err := mlbapi.GetGameUpdate(client, g.Link)
	if err != nil {
		return err
	}

	g.HomeScore = gameup.HomeScore
	g.VisitorScore = gameup.VisitorScore
	g.Status = gameup.Status
	g.Inning = gameup.Inning

	return nil
}

type MLBTeamData struct {
	Team      *MLBTeam
	NextGames []MLBGame
	PastGames []MLBGame
	Standings *mlbapi.Record
}

// GamesByTeam fetches next and last 5 games for the team, either home or visitor.
func (t *MLBTeam) GamesByTeam(ctx context.Context, q *modelstore.Queries, d time.Time) (*MLBTeamData, error) {
	games, err := q.MLBNextGamesByTeam(ctx, modelstore.MLBNextGamesByTeamParams{
		Gametime: sql.NullTime{Time: d, Valid: true},
		TeamID:   int64(t.TeamID),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get coming games: %s", err)
	}

	td := MLBTeamData{Team: t}

	for i, g := range games {
		td.NextGames = append(td.NextGames, fromDBGamesByTeam(g.MlbGame, g.Hometeamid, g.Awayteamid, g.Homename, g.Awayname))

		// Grab the team data. More than name??
		if i == 0 {
			if g.Hometeamid == int64(t.TeamID) {
				td.Team.Name = g.Homename.String
			} else if g.Awayteamid == int64(t.TeamID) {
				td.Team.Name = g.Awayname.String
			}
		}
	}

	gamesPast, err := q.MLBLastGamesByTeam(ctx, modelstore.MLBLastGamesByTeamParams{
		Gametime: sql.NullTime{Time: d, Valid: true},
		TeamID:   int64(t.TeamID),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get past games: %s", err)
	}

	for i, g := range gamesPast {
		td.PastGames = append(td.PastGames, fromDBGamesByTeam(g.MlbGame, g.Hometeamid, g.Awayteamid, g.Homename, g.Awayname))

		// Try again, if neeed?
		if i == 0 && td.Team.Name == "" {
			if g.Hometeamid == int64(t.TeamID) {
				td.Team.Name = g.Homename.String
			} else if g.Awayteamid == int64(t.TeamID) {
				td.Team.Name = g.Awayname.String
			}
		}
	}

	return &td, nil
}

// GetStandings fetches the relevant division standings for this team.
func (t *MLBTeam) GetStandings(client *http.Client) (*mlbapi.Record, error) {
	st, errs := mlbapi.GetStandings(client)
	if err := errors.Join(errs...); err != nil {
		return nil, err
	}

	standings, ok := st.RecordByTeam[t.Name]
	if !ok {
		return nil, fmt.Errorf("standings not found for %s", t.Name)
	}

	return standings, nil
}

// Convert DB row to go struct.
func fromDBGamesByTeam(g modelstore.MlbGame, homeTeamID, awayTeamID int64, homeName, awayName sql.NullString) MLBGame {
	game := MLBGame{
		ID:          uint(g.ID),
		GameID:      int(g.GameID),
		Gametime:    g.Gametime.Time,
		Description: g.Description.String,
		Status:      g.Status.String,
		Link:        g.Link.String,
		HomeID:      uint(g.HomeID),
		Home: MLBTeam{
			ID:     uint(g.HomeID),
			TeamID: int(homeTeamID),
			Name:   homeName.String,
		},
		HomeScore: int(g.HomeScore.Int64),
		VisitorID: uint(g.VisitorID),
		Visitor: MLBTeam{
			ID:     uint(g.VisitorID),
			TeamID: int(awayTeamID),
			Name:   awayName.String,
		},
		VisitorScore: int(g.VisitorScore.Int64),
	}

	return game
}

// fromDB converts the sqlc type into the general MLBGame type.
// TODO can this be a more generic method? This doesn't work for NextGamesByTeam, for example.
func fromDB(row modelstore.MLBFindGameByIDRow) *MLBGame {
	var mg MLBGame
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

	return &mg
}
