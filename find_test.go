package cleanprops

import (
	"fmt"
	"testing"
)

func TestFindSimple(t *testing.T) {
	props := LoadProps("testdata/find.properties")
	result := Find(props.SortedKeys(), "testdata/Example.java")

	if len(result) != 2 {
		t.Errorf("expected to find %d keys", 2)
	}
}

func TestFind(t *testing.T) {
	props := LoadProps("testdata/find.properties")
	result := Find(props.SortedKeys(), "testdata/Camel.xml")

	if len(result) != 4 {
		t.Errorf("expected to find %d keys but was %d", 4, len(result))
		fmt.Println(result)
	}
}