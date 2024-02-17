package models_test

import (
	"context"
	"encoding/json"
	"free-ent-guide-backend/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser_DB(t *testing.T) {
	wipeUsers(t)

	email := "johndoe@email.com"
	pw := "password123"
	u := &models.User{Email: email, Password: pw}
	err := u.Create(Queries)
	require.NoError(t, err)
	assert.NotEqual(t, 0, u.ID)
	err = u.AddZip(Queries, int64(33161))
	require.NoError(t, err)
	err = u.AddZip(Queries, int64(20002))
	require.NoError(t, err)

	u2 := &models.User{Email: email}
	err = u2.FindByEmail(Queries)
	require.NoError(t, err)
	assert.NotEqual(t, 0, u2.ID)
	assert.Equal(t, email, u2.Email)
	assert.False(t, u2.CreatedAt.IsZero())
	assert.NotEmpty(t, u2.Data.Zips)
	assert.Len(t, u2.Data.Zips, 2)

	u3 := models.User{Email: email}
	ok := u3.CheckPasswordHash(Queries, pw)
	assert.True(t, ok)

	ok = u3.CheckPasswordHash(Queries, "bad password")
	assert.False(t, ok)

	err = u3.GetZips(Queries)
	require.NoError(t, err)
	assert.NotEmpty(t, u2.Data.Zips)
	assert.Equal(t, int64(33161), u2.Data.Zips[0])

	err = u3.ClearZips(Queries)
	require.NoError(t, err)
	assert.Empty(t, u3.Data.Zips)

	err = u3.AddZip(Queries, int64(90210))
	require.NoError(t, err)
	err = u3.AddZip(Queries, int64(92060))
	require.NoError(t, err)
	assert.Len(t, u3.Data.Zips, 2)

	err = u3.DeleteZip(Queries, int64(90210))
	require.NoError(t, err)
	assert.Len(t, u3.Data.Zips, 1)

	err = u3.DeleteZip(Queries, int64(12345))
	assert.Error(t, err)
}

func TestUserData(t *testing.T) {
	data := []byte(`{"zips": [33161, 20002]}`)
	ud := models.UserData{}
	err := json.Unmarshal(data, &ud)
	require.NoError(t, err)
	assert.Len(t, ud.Zips, 2)
}

func wipeUsers(t *testing.T) {
	err := Queries.UsersDelete(context.Background())
	require.NoError(t, err)
}
