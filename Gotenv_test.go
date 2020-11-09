package gotenv

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	env := strings.NewReader("a=1\nb=2\nc=3")
	expected := dotenv{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	result := dotenv{}
	if err := parse(env, result); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}
