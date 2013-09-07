package redis

import (
	"testing"
)

func initredis() *Redis {
	r := &Redis{}
	if err := r.Connect(); err != nil {
		panic(err)
	}
	return r
}

func TestAuth(t *testing.T) {
	r := initredis()
	defer r.Quit()
	if err := r.Auth("123"); err != nil {
	} else {
		t.Fatal(err.Error())
	}
}

func TestSelect(t *testing.T) {
	r := initredis()
	defer r.Quit()
	if err := r.Select(1); err != nil {
		t.Fatal(err.Error())
	}
	if err := r.Select(65535); err != nil {
	} else {
		t.Fatal(err.Error())
	}
}

func TestEcho(t *testing.T) {
	r := initredis()
	defer r.Quit()
	if message, err := r.Echo("Hello"); err != nil {
		t.Fatal(err.Error())
	} else if message != "Hello" {
		t.Fatal(message)
	}
}

func TestPing(t *testing.T) {
	r := initredis()
	defer r.Quit()
	if _, err := r.Ping(); err != nil {
		t.Fatal(err.Error())
	}
}

func TestQuit(t *testing.T) {
	r := initredis()
	defer r.Quit()
	if err := r.Quit(); err != nil {
		t.Fatal(err.Error())
	}
}