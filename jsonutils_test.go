package main

import (
	"testing"
)

const WANTED = `{
  "title": "example glossary"
}`

const INPUT = `
{
"title": "example glossary"
}`

func TestStringToPrettyJSON (t *testing.T) {
	want := WANTED
	have := stringToPrettyJSON(INPUT)

	if want != have {
		t.Errorf("got %s, want %s", have, want)
	}
}