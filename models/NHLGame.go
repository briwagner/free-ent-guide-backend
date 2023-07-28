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

// UpdateScore sets the scores for a completed game.
func (g *NHLGame) UpdateScore(db Store) error {
	up, err := g.GetUpdate()
	if err != nil {
		return err
	}
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
