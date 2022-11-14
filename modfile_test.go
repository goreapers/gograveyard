package gograveyard

import (
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		filename          string
		expectedDirect    int
		expectedIndirects int
	}{
		{
			filename:          "simple.go.mod",
			expectedDirect:    5,
			expectedIndirects: 1,
		},
		{
			filename:          "telegraf.go.mod",
			expectedDirect:    182,
			expectedIndirects: 243,
		},
	}

	for _, test := range tests {
		data, err := GoModBytes(filepath.Join("testdata", test.filename))
		if err != nil {
			t.Fatal(err)
		}

		modFile, err := Parse(data)
		if err != nil {
			t.Fatal(err)
		}

		if len(modFile.Direct) != test.expectedDirect {
			t.Fatalf("Expected %d dependencies but received %d",
				test.expectedDirect, len(modFile.Direct))
		}

		if len(modFile.Indirect) != test.expectedIndirects {
			t.Fatalf("Expected %d dependencies but received %d",
				test.expectedDirect, len(modFile.Indirect))
		}
	}
}

func TestReplace(t *testing.T) {
	data, err := GoModBytes(filepath.Join("testdata", "replace.go.mod"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	modFile, err := Parse(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	expected := []*Module{
		{
			Path:    "example.com/fork/net",
			Version: "v1.4.5",
		},
		{
			Path:    "example.com/fork/http",
			Version: "v1.19.0",
		},
		{
			Path:    "example.com/fork/http",
			Version: "v1.19.0",
		},
	}

	for _, e := range expected {
		var found bool
		for _, d := range modFile.Direct {
			if e.Path == d.Path && e.Version == d.Version {
				found = true
				continue
			}
		}

		if found != true {
			t.Log("Modfile parsed from testdata:")
			for i, d := range modFile.Direct {
				t.Logf("[%d] path: %s version %s", i, d.Path, d.Version)
			}
			t.Fatalf("expected path: %s and version: %s but missing from parsed mod file", e.Path, e.Version)
		}
	}
}
