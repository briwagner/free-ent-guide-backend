package main

import (
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func handleBackup(tp TaskPayload, args []string) error {
	u, p, db := parseSqlString(tp.Cred.DB)

	cmd := exec.Command("mysqldump", "-u"+u, "-p"+p, db)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	data, err := io.ReadAll(stdout)
	if err != nil {
		return err
	}

	fn := "ent_backup_" + time.Now().Format("01-02-2006"+".sql")
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(data)

	return nil
}

func parseSqlString(s string) (u, p, db string) {
	r := regexp.MustCompile("^(.+)@")
	m := r.FindStringSubmatch(s)
	if len(m) == 0 {
		return u, p, db
	}

	pos := strings.Index(m[0], ":")
	if pos == 0 {
		return u, p, db
	}

	u = m[0][0:pos]
	p = m[0][pos+1 : len(m[0])-1]

	// Extract db name
	r2 := regexp.MustCompile(`\)/(.+)\?`)
	m2 := r2.FindStringSubmatch(s)
	if len(m2) < 2 {
		return u, p, db
	}

	db = m2[1] // the second match excludes the limiters
	return u, p, db
}
