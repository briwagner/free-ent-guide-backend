package models

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis"
)

var cacheClient *redis.Client

func init() {
	cacheClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:7000",
		Password: "",
		DB:       0,
	})
}

func TestGetDbKey(t *testing.T) {
	n := "johnDoe"
	u := User{Name: n}
	got := u.getDbKey()
	if fmt.Sprintf("user:%s", n) != got {
		t.Errorf("Got %s, want %s", got, n)
	}
}

func TestGetUserName(t *testing.T) {
	n := "johnDoe"
	u := User{Name: n}
	got := u.GetUserName()
	if n != got {
		t.Errorf("Got %s, want %s", got, n)
	}
}

func TestPing(t *testing.T) {
	cmd := cacheClient.Ping()
	if cmd.Err() != nil {
		t.Errorf("Cannot ping redis %v", cmd.Err())
	}
}

func TestCreate(t *testing.T) {
	u := User{Name: "johnDoe", Password: "password"}
	err := u.Create(cacheClient)
	if err != nil {
		t.Errorf("User not created %s", err.Error())
	}

	if len(u.Zips) != 0 {
		t.Errorf("New user should have no zips. Has %d", len(u.Zips))
	}

	t.Cleanup(func() {
		cacheClient.FlushAll()
	})
}

func TestCheckPasswordHash(t *testing.T) {
	u := User{Name: "johnDoe", Password: "password"}
	u.Create(cacheClient)
	p := "password"
	check := u.CheckPasswordHash(cacheClient, p)
	if !check {
		t.Errorf("User password check failed using: %s.", p)
	}

	p = "passw0rd"
	check2 := u.CheckPasswordHash(cacheClient, p)
	if check2 {
		t.Errorf("User password check should fail using: %s", p)
	}

	t.Cleanup(func() {
		cacheClient.FlushAll()
	})
}

func TestGetZips(t *testing.T) {
	u := User{Name: "johnDoe", Password: "password"}
	u.Create(cacheClient)

	z := "11223"
	u.AddZip(cacheClient, z)
	u.GetZips(cacheClient)
	if u.Zips[0] != z {
		t.Errorf("Zip %s not added to user", z)
	}

	t.Cleanup(func() {
		cacheClient.FlushAll()
	})
}

func TestDeleteZip(t *testing.T) {
	u := User{Name: "johnDoe", Password: "password"}
	u.Create(cacheClient)

	z := "11223"
	z2 := "22122"
	u.AddZip(cacheClient, z)
	u.AddZip(cacheClient, z2)
	u.GetZips(cacheClient)
	if u.Zips[0] != z {
		t.Errorf("Zip %s not added to user", z)
	}

	u.DeleteZip(cacheClient, z2)
	u.GetZips(cacheClient)
	if len(u.Zips) != 1 {
		t.Errorf("User zip not deleted: %s", z2)
	}

	t.Cleanup(func() {
		cacheClient.FlushAll()
	})
}

func TestClearZips(t *testing.T) {
	u := User{Name: "johnDoe", Password: "password"}
	u.Create(cacheClient)

	z := "11223"
	z2 := "22122"
	u.AddZip(cacheClient, z)
	u.AddZip(cacheClient, z2)
	u.GetZips(cacheClient)
	if u.Zips[0] != z {
		t.Errorf("Zip %s not added to user", z)
	}

	u.ClearZips(cacheClient)
	u.GetZips(cacheClient)
	if len(u.Zips) != 0 {
		t.Errorf("User zips not cleared. Has %d zip", len(u.Zips))
	}

	t.Cleanup(func() {
		cacheClient.FlushAll()
	})
}
