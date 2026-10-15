package main

import (
	"encoding/json"
	"testing"
)

func TestJSONValue(t *testing.T) {
	tests := []struct {
		name         string
		value        any
		databaseType string
		want         string
	}{
		{name: "numeric string", value: "12.34", databaseType: "NUMERIC", want: "12.34"},
		{name: "numeric bytes", value: []byte("12.34"), databaseType: "NUMERIC", want: "12.34"},
		{name: "json string", value: `{}`, databaseType: "JSONB", want: `{}`},
		{name: "text bytes", value: []byte("hello"), databaseType: "TEXT", want: `"hello"`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			encoded, err := json.Marshal(jsonValue(test.value, test.databaseType))
			if err != nil {
				t.Fatal(err)
			}
			if got := string(encoded); got != test.want {
				t.Errorf("jsonValue() = %s, want %s", got, test.want)
			}
		})
	}
}
