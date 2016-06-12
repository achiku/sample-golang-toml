package main

import "github.com/BurntSushi/toml"

const config1 = `
[string_notnull]
db_types = ["character", "character varying", "text", "money"]
go_type = "string"
nil_value= "\"\""

[string_nullable]
db_types = ["character", "character varying", "text", "money"]
go_type = "sql.NullString"
nil_value= "\"\""

[time_notnull]
db_types = [
    "time with time zone",
	"time without time zone",
	"timestamp without time zone",
	"timestamp with time zone"
]
go_type = "time.Time"
nil_value= "0"

[bool_notnull]
db_types = ["boolean"]
go_type = "boole"
nil_value= "false"

[bool_nullable]
db_types = ["boolean"]
go_type = "boole"
nil_value= "false"
`

const config2 = `
[string."nullable"]
go_type = "sql.NullString"

[string."notnull"]
go_type = "string"
`

type ma struct {
	DBTypes  []string `toml:"db_types"`
	GoType   string   `toml:"go_type"`
	NilValue string   `toml:"nil_value"`
}

func loadConfig(config string) (map[string]ma, error) {
	var t map[string]ma
	if _, err := toml.Decode(config, &t); err != nil {
		return nil, err
	}
	return t, nil
}
