package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"free-ent-guide-backend/pkg/cred"
	"io"
	"net/http"
	"strings"
)

type SlackBody struct {
	Text string `json:"text"`
}

func slackMessage(c *cred.Cred, msg string) error {
	if c.SlackURL == "" {
		return nil
	}

	var b strings.Builder
	_, err := b.WriteString(msg)
	if err != nil {
		return fmt.Errorf("error writing to slack %w", err)
	}

	var sb SlackBody
	sb.Text = b.String()
	data, err := json.Marshal(sb)
	if err != nil {
		return err
	}

	cli := &http.Client{}
	resp, err := cli.Post(c.SlackURL, "applicaton/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		by, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("slackbot error resp, failed to read body %d", resp.StatusCode)
		}
		return fmt.Errorf("slackbot error resp %d: %s", resp.StatusCode, string(by))
	}

	return nil
}
