package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Slack(t *testing.T) {
	msg := "test message"
	err := slackMessage(msg)
	require.NoError(t, err)
}
