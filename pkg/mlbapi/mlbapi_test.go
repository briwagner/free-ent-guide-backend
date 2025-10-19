package mlbapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshal_Schedule(t *testing.T) {
	data, err := os.ReadFile("testdata/mlb_gameday.json")
	if err != nil {
		panic(err)
	}

	gameday := MLBGameDay{}
	err = json.Unmarshal(data, &gameday)
	require.NoError(t, err)

	date := "2023-09-12"
	assert.Len(t, gameday.Dates, 2, "includes two days of games")
	assert.Equal(t, gameday.Dates[0].Date, date, "parses time as raw string")
	assert.Len(t, gameday.Dates[0].Games, 17, "parses all games for a day")

	game := gameday.Dates[0].Games[0]
	assert.Equal(t, date, game.Time.Format("2006-01-02"), "parses time to generate date only")
	assert.Equal(t, date+" 17:35 UTC", game.Time.Format("2006-01-02 15:04 MST"), "parses game time")
	assert.Equal(t, "/api/v1.1/game/716628/feed/live", game.Link, "parses game link")
	assert.Equal(t, int(716628), game.GameID, "parses gameID")
	assert.Equal(t, "Regular Season", game.Description, "parses game description")
	assert.Equal(t, "Final", game.Status.Status, "parses game status")
	assert.Equal(t, 3, game.Teams.Away.Score)
	assert.Equal(t, "New York Yankees", game.Teams.Away.Name)
	assert.Equal(t, "/api/v1/teams/147", game.Teams.Away.Link)
	assert.Equal(t, 147, game.Teams.Away.ID)
	assert.Equal(t, 2, game.Teams.Home.Score)
	assert.Equal(t, "Boston Red Sox", game.Teams.Home.Name)
	assert.Equal(t, "/api/v1/teams/111", game.Teams.Home.Link)
	assert.Equal(t, 111, game.Teams.Home.ID)
}

func TestUnmarshal_Teams(t *testing.T) {
	by := []byte(`          {"teams": {
            "away": {
              "isWinner": true,
              "leagueRecord": {
                "losses": 72,
                "pct": ".500",
                "wins": 72
              },
              "score": 3,
              "seriesNumber": 47,
              "splitSquad": false,
              "team": {
                "id": 147,
                "link": "/api/v1/teams/147",
                "name": "New York Yankees"
              }
            },
            "home": {
              "isWinner": false,
              "leagueRecord": {
                "losses": 71,
                "pct": ".507",
                "wins": 73
              },
              "score": 2,
              "seriesNumber": 47,
              "splitSquad": false,
              "team": {
                "id": 111,
                "link": "/api/v1/teams/111",
                "name": "Boston Red Sox"
              }
            }
          }}`)

	var game MLBGame
	require.NoError(t, json.Unmarshal(by, &game))
	assert.Equal(t, "New York Yankees", game.Teams.Away.Name)
	assert.Equal(t, 147, game.Teams.Away.ID)
	assert.Equal(t, 3, game.Teams.Away.Score)
	assert.Equal(t, "Boston Red Sox", game.Teams.Home.Name)
	assert.Equal(t, 111, game.Teams.Home.ID)
	assert.Equal(t, 2, game.Teams.Home.Score)
	assert.Equal(t, "/api/v1/teams/111", game.Teams.Home.Link)

}

func TestUnmarshal_Update(t *testing.T) {
	data, err := os.ReadFile("testdata/mlb_gameupdate.json")
	require.NoError(t, err)

	g := MLBGameUpdate{}
	err = json.Unmarshal(data, &g)
	require.NoError(t, err)
	assert.Equal(t, int64(716608), g.GamePK)
	assert.Equal(t, "Final", g.Status)
	assert.Equal(t, 9, g.Inning)
	assert.Equal(t, 11, g.HomeScore)
	assert.Equal(t, 2, g.VisitorScore)

	data2, err := os.ReadFile("testdata/mlb_gameupdate_canceled.json")
	require.NoError(t, err)

	g2 := MLBGameUpdate{}
	err = json.Unmarshal(data2, &g2)
	require.Error(t, err)
	assert.Equal(t, int64(0), g2.GamePK)
}

func TestUnmarshal_Standings(t *testing.T) {
	by, err := os.ReadFile("testdata/mlb_standings.json")
	require.NoError(t, err)
	var st Standings
	require.NoError(t, json.Unmarshal(by, &st))

	assert.Len(t, st.Records, 3)

	alWest := st.Records[0]
	assert.Len(t, alWest.TeamRecords, 5)
	for _, team := range alWest.TeamRecords {
		assert.NotEqual(t, "", team.Team.Name)
	}

	yanks := alWest.TeamRecords[0]
	assert.Equal(t, "New York Yankees", yanks.Team.Name)
	assert.Equal(t, 147, yanks.Team.ID)
	assert.Equal(t, "1", yanks.DivisionRank)
	assert.Equal(t, "3", yanks.LeagueRank)
	assert.Equal(t, 3, yanks.LeagueRecord.Losses)
	assert.Equal(t, 6, yanks.LeagueRecord.Wins)
	assert.Equal(t, 0, yanks.LeagueRecord.Ties)
	assert.Equal(t, ".667", yanks.LeagueRecord.Percentage)
}

func TestGameUpdate(t *testing.T) {
	g, err := GetGameUpdate(&http.Client{Timeout: time.Second * 3}, "/api/v1.1/game/716608/feed/live")
	require.NoError(t, err)
	assert.NotEqual(t, 716608, g.GamePK)
	// other vals same as above
}

func TestGetStandings(t *testing.T) {
	t.Skip("live stats will vary")

	st, err := GetStandings(&http.Client{Timeout: time.Second * 3})
	require.Len(t, err, 2)
	require.NoError(t, err[0])
	assert.Len(t, st.Records, 6) // 3 for each league
	for _, div := range st.Records {
		assert.NotEqual(t, 0, div.Division.ID)
		assert.Equal(t, "regularSeason", div.StandingsType)
	}

	alWest := st.Records[0]
	assert.Len(t, alWest.TeamRecords, 5)
	for _, team := range alWest.TeamRecords {
		assert.NotEqual(t, "", team.Team.Name)
	}

	yanks := alWest.TeamRecords[0]
	assert.Equal(t, "New York Yankees", yanks.Team.Name)
	assert.Equal(t, 147, yanks.Team.ID)
	assert.Equal(t, "1", yanks.DivisionRank)
	assert.Equal(t, "3", yanks.LeagueRank)
	assert.Equal(t, 3, yanks.LeagueRecord.Losses)
	assert.Equal(t, 6, yanks.LeagueRecord.Wins)
	assert.Equal(t, 0, yanks.LeagueRecord.Ties)
	assert.Equal(t, ".667", yanks.LeagueRecord.Percentage)

	yanksRec, ok := st.RecordByTeam["New York Yankees"]
	assert.True(t, ok)
	assert.Equal(t, 5, len(yanksRec.TeamRecords))
	var found bool
	for _, rec := range yanksRec.TeamRecords {
		if rec.Team.ID == 147 {
			found = true
		}
	}
	assert.True(t, found)
}

func TestGetDivisions(t *testing.T) {
	divisions, err := GetDivisions()
	require.NoError(t, err)
	assert.Len(t, divisions, 6)
	alw, ok := divisions["200"]
	assert.True(t, ok)
	assert.Equal(t, "American League West", alw.Name)
	assert.Equal(t, "ALW", alw.Abbreviation)
}

func TestGetGames_Backoff(t *testing.T) {
	var tries int

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tries++
		if tries < 3 {
			time.Sleep(time.Second * 6)
			return
		}
		by, err := os.ReadFile("./testdata/mlb_gameday.json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(by))
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	cli := &http.Client{Timeout: time.Second * 2}

	ti, err := time.Parse("2006-01-02", "2025-08-11")
	require.NoError(t, err)

	gd, err := importDates(cli, ti, ts.URL)
	assert.NoError(t, err)
	assert.NotNil(t, gd)
	assert.Equal(t, 3, tries)
	t.Log(gd)
}
