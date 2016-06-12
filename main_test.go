package main

import "testing"

func TestLoadConfig(t *testing.T) {
	c, err := loadConfig(config1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(c)
}
