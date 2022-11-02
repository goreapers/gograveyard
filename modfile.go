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
	Path     string
	Version  string
	Indirect bool
}

func Parse(data []byte) (*ModFile, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))

	var require bool
	var f ModFile

	for scanner.Scan() {
		switch s := scanner.Text(); s {
		case "require (":
			require = true
		case ")":
			require = false
		default:
			if !require {
				continue
			}

			// If in a require section, parse the module
			// A module is required to have a path and version, the "indirect" comment is optional
			var mod Module
			m := strings.Split(strings.TrimSpace(s), " ")
			const pathAndVersion = 2
			if len(m) == pathAndVersion {
				mod.Path = m[0]
				mod.Version = m[1]
			}
			var indirect bool
			if len(m) == 4 && fmt.Sprintf("%s %s", m[2], m[3]) == "// indirect" {
				indirect = true
			}
			if len(m) >= pathAndVersion {
				if indirect {
					f.Indirect = append(f.Indirect, &mod)
				} else {
					f.Direct = append(f.Direct, &mod)
				}
			}
		}
	}

	return &f, nil
}
