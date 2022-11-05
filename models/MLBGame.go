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

type MLBGame struct {
	gorm.Model
	GameID      int       `json:"id" gorm:"uniqueIndex"`
	Link        string    `json:"link"`
	HomeID      uint      `json:"-"`
	Home        MLBTeam   `json:"home"`
	VisitorID   uint      `json:"-"`
	Visitor     MLBTeam   `json:"visitor"`
	Description string    `json:"description"`
	Gametime    time.Time `json:"gametime"`
}

type MLBTeam struct {
	gorm.Model
	TeamID int    `json:"id" gorm:"uniqueIndex"`
	Name   string `json:"name"`
	Link   string `json:"link"`
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

// GetUpdate makes an api request to get game update to
// merge with scheduled game info from the database.
func (g *MLBGame) GetUpdate() (MLBGameUpdate, error) {
	var gu MLBGameUpdate

	base := "https://statsapi.mlb.com"
	url := fmt.Sprintf("%s/%s", base, g.Link)
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

	// Set game from database.
	gu.Game = *g
	return gu, nil
}

// MLBGameUpdate wraps info on a scheduled game, live status (if applicable)
// and scores.
type MLBGameUpdate struct {
	ID           int64
	Game         MLBGame
	Status       string
	Inning       int64
	VisitorScore int64
	HomeScore    int64
}

// UnmarshalJSON is required to extract the minimal data needed on the frontend.
func (g *MLBGameUpdate) UnmarshalJSON(b []byte) error {
	var cg map[string]interface{}

	// Get ID.
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	err := dec.Decode(&cg)
	if err != nil {
		fmt.Println("error:", err)
	}
	id, err := cg["gamePk"].(json.Number).Int64()
	if err != nil {
		return err
	}
	g.ID = id

	// // Set status.
	gd := cg["gameData"].(map[string]interface{})
	st := gd["status"].(map[string]interface{})
	g.Status = st["detailedState"].(string)

	// // Set period.
	ld := cg["liveData"].(map[string]interface{})
	ls := ld["linescore"].(map[string]interface{})
	// If not live, the currentInning prop is not found.
	if _, ok := ls["currentInning"]; ok {
		inning, err := ls["currentInning"].(json.Number).Int64()
		if err != nil {
			return err
		}
		g.Inning = inning
	}

	// // Set scores.
	ts := ls["teams"].(map[string]interface{})
	ht := ts["home"].(map[string]interface{})
	if _, ok := ht["runs"]; ok {
		hg, err := ht["runs"].(json.Number).Int64()
		if err != nil {
			return err
		}
		g.HomeScore = hg
	}

	vt := ts["away"].(map[string]interface{})
	if _, ok := ht["runs"]; ok {
		vg, err := vt["runs"].(json.Number).Int64()
		if err != nil {
			return err
		}
		g.VisitorScore = vg
	}

	return nil
}
