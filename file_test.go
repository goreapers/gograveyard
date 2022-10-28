package gograveyard

import (
	"testing"
)

func TestGoModRead(t *testing.T) {
	if _, err := GoModBytes("go.mod"); err != nil {
		t.Fail()
	}
}

func TestGoModRead_doesNotExist(t *testing.T) {
	if _, err := GoModBytes("FileDoesNotExist"); err == nil {
		t.Fail()
	}
}
