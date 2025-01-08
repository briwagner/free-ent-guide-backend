package main

import (
	"free-ent-guide-backend/pkg/cred"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Slack(t *testing.T) {
	msg := "test message"
	err := slackMessage(&cred.Cred{}, msg)
	require.NoError(t, err)
}

func Test_ParseSqlString(t *testing.T) {
	s := "ent_user:ent_password@tcp(127.0.0.1:3306)/ent_v2?charset=utf8mb4&parseTime=True&loc=Local"
	u, p, db := parseSqlString(s)
	assert.Equal(t, "ent_user", u)
	assert.Equal(t, "ent_password", p)
	assert.Equal(t, "ent_v2", db)
}
