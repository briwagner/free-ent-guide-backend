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
	GameID      int       `json:"id" gorm:"uniqueIndex"`
	Link        string    `json:"link"`
	HomeID      uint      `json:"-"`
	Home        NHLTeam   `json:"home"`
	VisitorID   uint      `json:"-"`
	Visitor     NHLTeam   `json:"visitor"`
	Description string    `json:"description"`
	Gametime    time.Time `json:"gametime"`
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
	d = fmt.Sprintf("%s%%", d)
	tx := db.Preload("Home").Preload("Visitor").Where("gametime LIKE ?", d).Find(ngs)
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

// GetUpdate makes an api request to get game update to
// merge with scheduled game info from the database.
func (g *NHLGame) GetUpdate() (NHLGameUpdate, error) {
	var gu NHLGameUpdate

	base := "http://statsapi.web.nhl.com"
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

// UnmarshalJSON is required to extract the minimal data needed on the frontend.
func (g *NHLGameUpdate) UnmarshalJSON(b []byte) error {
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

	// Set status.
	gd := cg["gameData"].(map[string]interface{})
	st := gd["status"].(map[string]interface{})
	g.Status = st["detailedState"].(string)

	// Set period.
	ld := cg["liveData"].(map[string]interface{})
	ls := ld["linescore"].(map[string]interface{})
	period, err := ls["currentPeriod"].(json.Number).Int64()
	if err != nil {
		return err
	}
	g.Period = period

	// Set scores.
	ts := ls["teams"].(map[string]interface{})
	ht := ts["home"].(map[string]interface{})
	hg, err := ht["goals"].(json.Number).Int64()
	if err != nil {
		return err
	}
	g.HomeScore = hg

	vt := ts["away"].(map[string]interface{})
	vg, err := vt["goals"].(json.Number).Int64()
	if err != nil {
		return err
	}
	g.VisitorScore = vg

	return nil
}
