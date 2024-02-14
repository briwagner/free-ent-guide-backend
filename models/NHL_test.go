package models_test

import (
	"free-ent-guide-backend/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// cache_test.go runs init here to set:
// var Queries *modelstore.Queries

func TestDB_NHLTeam(t *testing.T) {
	t1 := models.NHLTeam{Name: "Buffalo Testteam", Link: "/nhl/test-team", TeamID: 99}
	err := t1.Create(Queries)
	require.NoError(t, err)
	assert.NotEqual(t, 0, t1.ID)
	t.Log(t1)

	t2 := models.NHLTeam{}
	err = t2.FindByTeamID(Queries, 7)
	require.NoError(t, err)
	assert.Equal(t, "Buffalo Sabres", t2.Name)
}

func TestDB_NHLGame(t *testing.T) {
	t.Skip()

	// TODO find a way to wipe DB every time and import from JSON.
	ng := models.NHLGame{
		GameID:       1414,
		Gametime:     time.Date(2024, 2, 2, 12, 25, 54, 00, time.UTC),
		Description:  "test game",
		Status:       "test status",
		Link:         "test link",
		HomeID:       17,
		VisitorID:    21,
		HomeScore:    4,
		VisitorScore: 3,
	}
	err := ng.Create(Queries)
	require.NoError(t, err)
	assert.NotEqual(t, 0, ng.ID)

	ng2 := &models.NHLGame{}
	err = ng2.FindByGameID(Queries, ng.GameID)
	require.NoError(t, err)
	assert.Equal(t, 4, ng2.HomeScore)
	assert.Equal(t, 3, ng2.VisitorScore)
	t.Log(ng2)

	ng2.HomeScore = 5
	ng2.VisitorScore = 7
	ng2.Status = "OFF"
	err = ng2.UpdateScorev2(Queries)
	require.NoError(t, err)

	ngs := models.NHLGames{}
	err = ngs.LoadByDate(Queries, time.Date(2024, 1, 11, 0, 0, 0, 0, time.UTC).Format("2006-01-02"))
	require.NoError(t, err)
	assert.Len(t, ngs, 26)
}
