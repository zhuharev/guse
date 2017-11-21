package parser

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	testData := `#                   comment

  
Cb            users.get                     users.Get
`
	routes, err := Parse(strings.NewReader(testData))
	if err != nil {
		t.Fatal(err)
	}
	if len(routes) != 1 {
		t.Fatalf("routes not found")
	}
	if routes[0].Method != "Cb" {
		t.Fatalf("method not recognized")
	}
	if routes[0].Name != "users.get" {
		t.Fatalf("name not recognized")
	}
	if routes[0].Handler != "users.Get" {
		t.Fail()
	}
}
