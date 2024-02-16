package mlbapi

import (
	"encoding/json"
	"os"
	"testing"

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

func TestUnmarshal_Update(t *testing.T) {
	data, err := os.ReadFile("testdata/mlb_gameupdate.json")
	if err != nil {
		panic(err)
	}

	g := MLBGameUpdate{}
	err = json.Unmarshal(data, &g)
	require.NoError(t, err)
	assert.Equal(t, int64(716608), g.GamePK)
	assert.Equal(t, "Final", g.Status)
	assert.Equal(t, 9, g.Inning)
	assert.Equal(t, 11, g.HomeScore)
	assert.Equal(t, 2, g.VisitorScore)
}

func TestGameUpdate(t *testing.T) {
	g, err := GetGameUpdate("/api/v1.1/game/716608/feed/live")
	require.NoError(t, err)
	assert.NotEqual(t, 716608, g.GamePK)
	// other vals same as above
}
