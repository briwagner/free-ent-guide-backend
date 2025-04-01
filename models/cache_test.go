package models_test

import (
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/cred"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	Queries *modelstore.Queries
)

func init() {
	// This uses test DB. Change to ent_v2 for other operations.
	c := cred.Cred{DB: "ent_user:ent_password@tcp(127.0.0.1:3306)/ent_guide_test?charset=utf8mb4&parseTime=True&loc=Local"}
	st := models.Setup(c)

	Queries = modelstore.New(st)

	cWipe := &models.Caches{}
	_, err := cWipe.DropAll(Queries)
	if err != nil {
		panic(err)
	}
}

func Test_Cache(t *testing.T) {
	// Create
	now := time.Now()
	c1 := &models.Cache{
		Name:      "/v1/movies?zip=60048&date=2023-10-12",
		Value:     "cached values",
		UpdatedAt: time.Now(),
		Expires:   now.AddDate(0, 0, -2),
	}
	err := c1.Insert(Queries)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, c1.ID)

	// Seeds
	c2 := &models.Cache{
		Name:    "/v1/discover",
		Value:   "cached values",
		Expires: now.AddDate(0, 0, -2),
	}
	err = c2.Insert(Queries)
	require.NoError(t, err)
	c3 := &models.Cache{
		Name:    "/v1/tv-movies",
		Value:   "cached values",
		Expires: now.AddDate(0, 0, 4),
	}
	err = c3.Insert(Queries)
	require.NoError(t, err)
	c4 := &models.Cache{
		Name:    "/v1/tv-sports",
		Value:   "cached values",
		Expires: now.AddDate(0, 0, 1),
	}
	err = c4.Insert(Queries)
	require.NoError(t, err)

	// Get by name
	var cName models.Cache
	err = cName.GetByName("/v1/movies?zip=60048&date=2023-10-12", Queries)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, cName.ID)
	assert.Equal(t, "/v1/movies?zip=60048&date=2023-10-12", cName.Name)
	assert.NotEqual(t, "", cName.Value)
	assert.GreaterOrEqual(t, cName.UpdatedAt.Year(), 2020)
	assert.GreaterOrEqual(t, cName.Expires.Year(), 2020)

	// Get by ID
	var c models.Cache
	err = c.GetByID(c1.ID, Queries)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, c.ID)
	assert.Equal(t, "/v1/movies?zip=60048&date=2023-10-12", c.Name)
	assert.NotEqual(t, "", c.Value)
	assert.GreaterOrEqual(t, c.UpdatedAt.Year(), 2020)
	assert.GreaterOrEqual(t, c.Expires.Year(), 2020)

	// Get all
	cAll := &models.Caches{}
	err = cAll.ShowAll(Queries)
	assert.NoError(t, err)
	assert.Len(t, *cAll, 4)

	// Get stale
	cStale := &models.Caches{}
	err = cStale.ShowStale(time.Now(), Queries)
	assert.NoError(t, err)
	assert.Len(t, *cStale, 2)

	// Clear stale
	cClear := &models.Caches{}
	rows, err := cClear.DropStale(time.Now(), Queries)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), rows)

	// Wipe all
	cWipe := &models.Caches{}
	rows, err = cWipe.DropAll(Queries)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), rows)
}
