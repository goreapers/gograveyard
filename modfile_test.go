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
			t.Log(err)
			t.Fail()
		}

		modFile, err := Parse(data)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		if len(modFile.Direct) != test.expectedDirect {
			t.Logf("Expected %d dependencies but received %d",
				test.expectedDirect, len(modFile.Direct))
			t.Fail()
		}

		if len(modFile.Indirect) != test.expectedIndirects {
			t.Logf("Expected %d dependencies but received %d",
				test.expectedDirect, len(modFile.Indirect))
			t.Fail()
		}
	}
}
