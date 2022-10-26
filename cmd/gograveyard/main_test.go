package main

import (
	"errors"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	if err := parse([]string{"../../go.mod"}); err != nil {
		t.Fail()
	}
}

func TestParse_NoArgs(t *testing.T) {
	if err := parse([]string{}); !errors.Is(err, errNoFileArg) {
		t.Fail()
	}
}

func TestParse_TooManyArgs(t *testing.T) {
	if err := parse([]string{"fake", "invalid"}); !errors.Is(err, errTooManyFiles) {
		t.Fail()
	}
}

func TestParse_DoesNotExist(t *testing.T) {
	if err := parse([]string{"fake"}); err == nil {
		t.Fail()
	}
}

func TestHelp(t *testing.T) {
	help := helpString()

	if help == "" {
		t.Fatalf("all help output is missing")
	}

	if !strings.Contains(help, "  help") {
		t.Fatalf("help missing")
	}

	if !strings.Contains(help, "  parse") {
		t.Fatalf("parse missing")
	}

	if !strings.Contains(help, "  version") {
		t.Fatalf("version missing")
	}

	if !strings.Contains(help, "--help") {
		t.Fatalf("help flag missing")
	}

	if !strings.Contains(help, "--json") {
		t.Fatalf("json flag missing")
	}
}

func TestVersion(t *testing.T) {
	version := versionString()

	if version == "" {
		t.Fatalf("version is empty")
	}

	if !strings.HasPrefix(version, "gograveyard (v") {
		t.Fatalf("version not prefixed correctly: '%s'", version)
	}

	if strings.Count(version, ".") != 2 {
		t.Fatalf("version does not contain major & minor separator: '%s'", version)
	}
}
