package main

import "github.com/BurntSushi/toml"

type ma struct {
	DBTypes  []string `toml:"db_types"`
	GoType   string   `toml:"go_type"`
	NilValue string   `toml:"nil_value"`
}

type mb struct {
	GoType   string `toml:"go_type"`
	NilValue string `toml:"nil_value"`
}

func loadConfig1(config string) (map[string]ma, error) {
	var t map[string]ma
	if _, err := toml.Decode(config, &t); err != nil {
		return nil, err
	}
	return t, nil
}

func loadConfig2(config string) (map[string]map[string]mb, error) {
	var t map[string]map[string]mb
	if _, err := toml.Decode(config, &t); err != nil {
		return nil, err
	}
	return t, nil
}
