package main

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	payload := ReadConfig()
	if payload.User == "" {
		t.Fatalf("No user set in config")
	}
	if payload.Repo == "" {
		t.Fatalf("No repo set in config")
	}
}

func TestMain(m *testing.M) {
	m.Run()
	return
}
