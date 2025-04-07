package mlbapi

import (
	"bytes"
	"encoding/json"
	"math"
	"time"
)

type (
	MLBGameDay struct {
		Dates []struct {
			Date  string    `json:"date"`
			Games []MLBGame `json:"games"`
		} `json:"dates"`
	}

	MLBGame struct {
		Time        time.Time    `json:"gameDate"`
		Link        string       `json:"link"`
		GameID      int          `json:"gamePk"`
		Description string       `json:"seriesDescription"` // 'Regular Season', etc.
		Status      StatusString `json:"status"`
		Teams       MLBTeams     `json:"teams"`
	}

	StatusString struct {
		Status string
	}

	MLBTeams struct {
		Home MLBTeam `json:"home"`
		Away MLBTeam `json:"away"`
	}

	MLBTeam struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Link  string `json:"link"`
		Score int    `json:"score"` // populated only for schedule lookup
	}

	MLBGameUpdate struct {
		GamePK       int64
		Game         MLBGame
		Status       string
		Inning       int
		HomeScore    int
		VisitorScore int
	}

	DivisionPayload struct {
		Divisions []Division `json:"divisions"`
	}

	Division struct {
		Abbreviation string `json:"abbreviation"`
		Active       bool   `json:"active"`
		ID           int    `json:"id"`
		League       struct {
			ID   int    `json:"id"`
			Link string `json:"link"`
		} `json:"league"`
		Name      string `json:"name"`
		NameShort string `json:"nameShort"`
		Link      string `json:"link"`
	}

	// Standings
	Record struct { // by division
		StandingsType string `json:"standingsType"`
		Division      struct {
			ID   json.Number `json:"id"`
			Name string      `json:"name"`
		} `json:"division"`
		TeamRecords []struct {
			Team         MLBTeam `json:"team"`
			DivisionRank string  `json:"divisionRank"`
			LeagueRank   string  `json:"leagueRank"`
			LeagueRecord struct {
				Wins       int    `json:"wins"`
				Losses     int    `json:"losses"`
				Ties       int    `json:"ties"`
				Percentage string `json:"pct"`
			} `json:"leagueRecord"`
		} `json:"teamRecords"`
	}

	Standings struct { // by league
		Records []Record `json:"records"`

		// Build our own map to make this searchable.
		RecordByTeam map[string]*Record `json:"-"`
	}
)

func (ss *StatusString) UnmarshalJSON(data []byte) error {
	var vals map[string]interface{}
	if err := json.Unmarshal(data, &vals); err != nil {
		return err
	}
	state, ok := vals["detailedState"]
	if ok {
		ss.Status = state.(string)
	}

	return nil
}

func (t *MLBTeams) UnmarshalJSON(data []byte) error {
	var vals map[string]interface{}
	if err := json.Unmarshal(data, &vals); err != nil {
		return err
	}

	if h, ok := vals["home"]; ok {
		htValues := h.(map[string]interface{})

		var home MLBTeam
		by, err := json.Marshal(htValues["team"])
		if err != nil {
			return err
		}
		err = json.Unmarshal(by, &home)
		if err != nil {
			return err
		}
		t.Home = home

		if sc, ok := htValues["score"]; ok {
			t.Home.Score = int(math.Round(sc.(float64)))
		}
	}

	if aw, ok := vals["away"]; ok {
		awValues := aw.(map[string]interface{})

		var away MLBTeam
		by, err := json.Marshal(awValues["team"])
		if err != nil {
			return err
		}
		err = json.Unmarshal(by, &away)
		if err != nil {
			return err
		}
		t.Away = away

		if sc, ok := awValues["score"]; ok {
			t.Away.Score = int(math.Round(sc.(float64)))
		}
	}

	return nil
}

func (gu *MLBGameUpdate) UnmarshalJSON(data []byte) error {
	var vals map[string]interface{}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	err := dec.Decode(&vals)
	if err != nil {
		return err
	}

	id, err := vals["gamePk"].(json.Number).Int64()
	if err != nil {
		return err
	}
	gu.GamePK = id

	gd := vals["gameData"].(map[string]interface{})
	st := gd["status"].(map[string]interface{})
	// TODO add nil check
	gu.Status = st["detailedState"].(string)

	ld := vals["liveData"].(map[string]interface{})
	ls := ld["linescore"].(map[string]interface{})
	// If not live, the currentInning prop is not found.
	if _, ok := ls["currentInning"]; ok {
		inning, err := ls["currentInning"].(json.Number).Int64()
		if err != nil {
			return err
		}
		gu.Inning = int(inning)
	}

	ts := ls["teams"].(map[string]interface{})
	ht := ts["home"].(map[string]interface{})
	if _, ok := ht["runs"]; ok {
		hg, err := ht["runs"].(json.Number).Int64()
		if err != nil {
			return err
		}
		gu.HomeScore = int(hg)
	}

	vt := ts["away"].(map[string]interface{})
	if _, ok := ht["runs"]; ok {
		vg, err := vt["runs"].(json.Number).Int64()
		if err != nil {
			return err
		}
		gu.VisitorScore = int(vg)
	}
	return nil
}
