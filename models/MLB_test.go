package models_test

import (
	"context"
	"free-ent-guide-backend/models"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// cache_test.go runs init here. To set:
// var Queries *modelstore.Queries

func TestDB_MLB(t *testing.T) {
	mlbWipe(t)

	t1 := models.MLBTeam{TeamID: 117, Name: "Houston Astros"}
	err := t1.Create(Queries)
	require.NoError(t, err)
	assert.NotEqual(t, uint(0), t1.ID)

	t2 := models.MLBTeam{TeamID: 143, Name: "Philadelphia Phillies"}
	err = t2.Create(Queries)
	require.NoError(t, err)
	require.NotEqual(t, uint(0), t2.ID)

	tq := models.MLBTeam{}
	err = tq.FindByTeamID(Queries, 117)
	require.NoError(t, err)
	assert.Equal(t, "Houston Astros", tq.Name)
	assert.Equal(t, int(117), tq.TeamID)

	gt := time.Date(2024, 2, 2, 12, 25, 54, 00, time.UTC)
	g := models.MLBGame{GameID: 444, Gametime: gt, Description: "demo game", Status: "Final", HomeID: t1.ID, HomeScore: 3, VisitorID: t2.ID, VisitorScore: 7}
	err = g.Create(Queries)
	require.NoError(t, err)
}

func TestGameUpdate(t *testing.T) {
	// TODO better way to seed this?
	mlbWipe(t)

	th := models.MLBTeam{TeamID: 111, Name: "Boston Red Sox"}
	err := th.Create(Queries)
	require.NoError(t, err)
	ta := models.MLBTeam{TeamID: 147, Name: "New York Yankees"}
	err = ta.Create(Queries)
	require.NoError(t, err)

	gt := time.Date(2024, 2, 2, 12, 25, 54, 00, time.UTC)
	g := models.MLBGame{GameID: 716628, Gametime: gt, Description: "Regular Season", Status: "In Progress", Link: "/api/v1.1/game/716628/feed/live", HomeID: th.ID, HomeScore: 0, VisitorID: ta.ID, VisitorScore: 0}
	err = g.Create(Queries)
	require.NoError(t, err)

	err = g.GetUpdate(&http.Client{Timeout: time.Second * 3})
	require.NoError(t, err)
	assert.Equal(t, 3, g.VisitorScore)
	assert.Equal(t, 2, g.HomeScore)
	assert.Equal(t, 9, g.Inning)
	assert.Equal(t, "Final", g.Status)
	assert.NotEqual(t, 0, g.ID)
}

func TestGamesByDate(t *testing.T) {
	g := &models.MLBGames{}
	err := g.LoadByDate(Queries, "2024-02-02")
	require.NoError(t, err)
	assert.NotEmpty(t, *g)
	for _, v := range *g {
		assert.NotNil(t, v.ID)
		assert.NotNil(t, v.GameID)
	}
}

func TestGameByID(t *testing.T) {
	t.Skipf("need a valid DB, i.e. switch to ent_v2")

	ctx := context.Background()
	team := models.MLBTeam{ID: 29}
	ti, err := time.Parse("2006-01-02", "2025-03-25")
	require.NoError(t, err)
	gameData, err := team.GamesByTeam(ctx, Queries, ti)
	require.NoError(t, err)

	assert.Equal(t, 4, len(gameData.NextGames))
}

func mlbWipe(t *testing.T) {
	err := Queries.MLBDeleteGames(context.Background())
	require.NoError(t, err)
	err = Queries.MLBDeleteTeams(context.Background())
	require.NoError(t, err)
}
