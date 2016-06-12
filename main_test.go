package main

import "testing"

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
nil_value= "sql.NullString{}"

[string."notnull"]
go_type = "string"
nil_value= "\"\""
`

func TestLoadConfig1(t *testing.T) {
	c, err := loadConfig1(config1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", c)
}

func TestLoadConfig2(t *testing.T) {
	c, err := loadConfig2(config2)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range c {
		for n, m := range v {
			t.Logf("%s %s", k, n)
			t.Logf("  %s %s", m.GoType, m.NilValue)
		}
	}
}
