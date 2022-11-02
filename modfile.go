package gograveyard

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type ModFile struct {
	Direct   []*Module
	Indirect []*Module
}

type Module struct {
	Path    string
	Version string
}

func Parse(data []byte) (*ModFile, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))

	var require bool
	var f ModFile

	for scanner.Scan() {
		switch scanner.Text() {
		case "require (":
			require = true

			continue
		case ")":
			require = false

			continue
		}

		if !require {
			continue
		}

		// If in a require section, parse the module
		// A module is required to have a path and version, the "indirect" comment is optional
		m := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		const pathAndVersion = 2
		// Skip if invalid module, missing required path or version
		if len(m) < pathAndVersion {
			continue
		}

		mod := Module{
			Path:    m[0],
			Version: m[1],
		}

		var indirect bool
		if len(m) == 4 && fmt.Sprintf("%s %s", m[2], m[3]) == "// indirect" {
			indirect = true
		}

		if indirect {
			f.Indirect = append(f.Indirect, &mod)
		} else {
			f.Direct = append(f.Direct, &mod)
		}
	}

	return &f, nil
}
