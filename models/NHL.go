package models

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"free-ent-guide-backend/models/modelstore"
	"io"
	"net/http"
	"time"
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

func (nt *NHLTeam) FromDB(team modelstore.NhlTeam) {
	nt.ID = uint(team.ID)
	nt.TeamID = int(team.TeamID)
	nt.Name = team.Name.String
	nt.Link = team.Link.String
	nt.Tricode = team.Tricode.String
}

type NHLGame struct {
	ID           uint
	GameID       int       `json:"id"`
	Gametime     time.Time `json:"gametime"`
	Description  string    `json:"description"`
	Status       string    `json:"status"`
	Link         string    `json:"link"`
	HomeID       uint      `json:"-"`
	Home         *NHLTeam  `json:"home"`
	HomeScore    int       `json:"home_score"`
	VisitorID    uint      `json:"-"`
	Visitor      *NHLTeam  `json:"visitor"`
	VisitorScore int       `json:"visitor_score"`
	UpdatedAt    time.Time `json:"updated_at"`
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
	up, err := g.GetUpdate()
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

// GetLatestGames loads all games on the latest date found in the DB.
func GetLatestGames(q *modelstore.Queries) (NHLGames, error) {
	var games []NHLGame

	rows, err := q.NHLLatestGames(context.Background())
	if err != nil {
		return games, err
	}

	for _, row := range rows {
		new := NHLGame{GameID: int(row.ID), Gametime: row.Gametime.Time}
		games = append(games, new)
	}

	return games, nil
}

// GetUpdate makes an api request to get game update to
// merge with scheduled game info from the database.
func (g *NHLGame) GetUpdate() (NHLGameUpdate, error) {
	var gu NHLGameUpdate

	base := "https://api-web.nhle.com/v1/gamecenter"
	url := fmt.Sprintf("%s/%d/play-by-play", base, g.GameID)
	resp, err := http.Get(url)
	if err != nil {
		return gu, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return gu, err
	}

	err = json.Unmarshal(body, &gu)
	if err != nil {
		return gu, err
	}

	// Set game data from database.
	gu.Game = *g
	return gu, nil
}

// NHLGameUpdate wraps info on a scheduled game, live status (if applicable)
// and scores.
type NHLGameUpdate struct {
	ID           int64
	Game         NHLGame
	Status       string
	Period       int64
	VisitorScore int64
	HomeScore    int64
}

// As of 11-2023, NHL api has multiple versions of 'Final'.
// Previously, we used 'Final' in db and as signal to front-end.
func setGameState(st string) string {
	switch st {
	case "OFF", "OVER", "FINAL":
		return "Final"
	}
	return st
}

// UnmarshalJSON is required to extract the minimal data needed on the frontend.
func (g *NHLGameUpdate) UnmarshalJSON(b []byte) error {
	var cg map[string]interface{}

	// Get ID.
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	err := dec.Decode(&cg)
	if err != nil {
		return fmt.Errorf("error decoding: %w", err)
	}
	id, err := cg["id"].(json.Number).Int64()
	if err != nil {
		return fmt.Errorf("error parsing ID: %w", err)
	}
	g.ID = id

	st := cg["gameState"].(string)
	g.Status = setGameState(st)

	// Only check for live or past games, else this key is not found.
	if g.Status == "LIVE" || g.Status == "Final" {
		period, err := cg["period"].(json.Number).Int64()
		if err == nil {
			g.Period = period
		}
	}

	// Set scores.
	home := cg["homeTeam"].(map[string]interface{})
	hScore, err := home["score"].(json.Number).Int64()
	if err != nil {
		return err
	}
	g.HomeScore = hScore

	away := cg["awayTeam"].(map[string]interface{})
	aScore, err := away["score"].(json.Number).Int64()
	if err != nil {
		return err
	}
	g.VisitorScore = aScore

	return nil
}
