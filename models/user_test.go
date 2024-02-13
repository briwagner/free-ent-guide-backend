package models

import (
	"testing"

	"free-ent-guide-backend/pkg/cred"
)

var AppStore Store

func init() {
	c := cred.Cred{DB: "gorm:password@tcp(127.0.0.1:3306)/demo_gorm_test?charset=utf8mb4&parseTime=True&loc=Local"}
	AppStore, _ = Setup(c)
}

func TestGetUserName(t *testing.T) {
	n := "johnDoe"
	u := User{Name: n}
	got := u.GetUserName()
	if n != got {
		t.Errorf("Got %s, want %s", got, n)
	}
}

func TestCreate(t *testing.T) {
	u := User{Name: "johnDoe", Password: "password", Zips: []UserZip{}}
	err := u.Create(AppStore)
	if err != nil {
		t.Errorf("User not created %s", err.Error())
	}

	u2 := User{Name: "Faye Dunnaway", Password: "fdlk09835j", Zips: []UserZip{{Zip: "33161"}}}
	err = u2.Create(AppStore)
	if err != nil {
		t.Errorf("User2 not created %s", err.Error())
	}

	if len(u2.Zips) != 1 {
		t.Errorf("New user should have one zip. Has %d", len(u2.Zips))
	}

	// @todo: drop records? drop tables?
}

func TestLoadUserByName(t *testing.T) {
	u := User{Name: "johnDoe", Password: "password", Zips: []UserZip{}}
	err := u.Create(AppStore)
	if err != nil {
		t.Errorf("User not created %s", err.Error())
	}

	u2 := User{Name: "johnDoe"}
	err = u2.LoadByName(AppStore)
	if err != nil {
		t.Errorf("User2 not created %s", err.Error())
	}

	if u2.ID == 0 {
		t.Errorf("Failed to load user by name")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	u := User{Name: "johnDoe", Password: "password"}
	u.Create(AppStore)
	p := "password"
	check := u.CheckPasswordHash(AppStore, p)
	if !check {
		t.Errorf("User password check failed using: %s.", p)
	}

	p = "passw0rd"
	check2 := u.CheckPasswordHash(AppStore, p)
	if check2 {
		t.Errorf("User password check should fail using: %s", p)
	}

	// @todo: cleanup records.
}

func TestZip(t *testing.T) {
	u := User{Name: "johnDoe", Password: "password", Zips: []UserZip{{Zip: "33161"}}}
	err := u.Create(AppStore)
	if err != nil {
		panic(err)
	}

	z := "11223"
	u.AddZip(AppStore, z)
	u.GetZips(AppStore)

	u2 := User{}
	AppStore.Preload("Zips").Find(&u2, u.ID)
	if len(u2.Zips) != 2 {
		t.Errorf("Failed to add user zips. Has %d, want %d", len(u2.Zips), 2)
	}

	ok, _ := u2.HasZip(AppStore, z)
	if !ok {
		t.Errorf("HasZip failed to find zip %s", z)
	}

	z2 := "32211"
	ok2, _ := u2.HasZip(AppStore, z2)
	if ok2 {
		t.Errorf("HasZip found zip that shouldn't be %s", z2)
	}
}

func TestDeleteZip(t *testing.T) {
	u := User{Name: "johnDoe", Password: "password", Zips: []UserZip{{Zip: "33161"}, {Zip: "20002"}}}
	err := u.Create(AppStore)
	if err != nil {
		panic(err)
	}

	if len(u.Zips) != 2 {
		t.Errorf("Failed to add save user with zips.")
		return
	}

	zd := "20002"
	u.DeleteZip(AppStore, zd)
	if len(u.Zips) != 1 {
		t.Errorf("Failed to delete zip. User has %d, want 1", len(u.Zips))
	}
	if u.Zips[0].Zip == zd {
		t.Error("DeleteZip failed to remove correct zip")
	}
}

func TestClearZips(t *testing.T) {
	u := User{Name: "johnDoe", Password: "password", Zips: []UserZip{{Zip: "33161"}, {Zip: "20002"}}}
	u.Create(AppStore)

	u.ClearZips(AppStore)
	if len(u.Zips) != 0 {
		t.Errorf("User zips not cleared. Has %d zip", len(u.Zips))
	}
}
