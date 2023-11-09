package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type NHLGame struct {
	gorm.Model
	GameID       int       `json:"id" gorm:"uniqueIndex:idx_gameid_gametime"`
	Link         string    `json:"link"`
	HomeID       uint      `json:"-"`
	Home         NHLTeam   `json:"home"`
	HomeScore    int       `json:"home_score" gorm:"home_score"`
	VisitorID    uint      `json:"-"`
	Visitor      NHLTeam   `json:"visitor"`
	VisitorScore int       `json:"visitor_score" gorm:"visitor_score"`
	Description  string    `json:"description"`
	Gametime     time.Time `json:"gametime" gorm:"uniqueIndex:idx_gameid_gametime"`
	Status       string    `json:"status" gorm:"status"`
}

type NHLTeam struct {
	gorm.Model
	TeamID int    `json:"id" gorm:"uniqueIndex"`
	Name   string `json:"name"`
	Link   string `json:"link"`
}

type NHLGames []NHLGame

// LoadByDate fetches all games for the given date.
func (ngs *NHLGames) LoadByDate(d string, db Store) error {
	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		return err
	}
	// Set date to noon to allow capturing anything in PM that may show as next-day with UTC.
	dateFrom := time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, date.Location())
	dateTo := dateFrom.Add(time.Hour * 24)
	tx := db.Preload("Home").Preload("Visitor").Where("gametime BETWEEN ? AND ?", dateFrom.String(), dateTo.String()).Order("gametime").Find(ngs)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindByID searches for a game by Game ID.
func (g *NHLGame) FindByID(id string, db Store) error {
	tx := db.Preload("Home").Preload("Visitor").Where("game_id = ?", id).Find(g)
	if tx.Error != nil {
		return tx.Error
	}
	if g.ID == 0 {
		return errors.New("NHL game not found")
	}
	return nil
}

// UpdateScore updates a DB record to set the scores for a completed game.
func (g *NHLGame) UpdateScore(db Store) error {
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

	err = db.Save(&g).Error
	if err != nil {
		return err
	}
	return nil
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
