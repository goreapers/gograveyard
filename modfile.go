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

type replaceInfo struct {
	TargetPath      string
	TargetVersion   string
	Originalversion string
}

func Parse(data []byte) (*ModFile, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))

	var require bool
	var f ModFile
	replace := make(map[string]replaceInfo)

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "replace") {
			parseReplace(scanner.Text(), replace)
			continue
		}

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
			fmt.Printf("warning: invalid module found, received %s \n", scanner.Text())
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
			continue
		}

		f.Direct = append(f.Direct, &mod)
	}

	// Update the paths and versions based on any replace directives found
	replaceMods(f.Indirect, replace)
	replaceMods(f.Direct, replace)

	return &f, nil
}

// parseReplace will parse the `replace directive` found in the mod file
// e.g. `replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5`
func parseReplace(line string, replace map[string]replaceInfo) {
	replaceTxt := strings.Split(strings.TrimSpace(line), " ")

	var originalPath string
	var info replaceInfo

	const middle = 3

	for i, r := range replaceTxt {
		if r == "replace" || r == "=>" {
			continue
		}

		if VerifySemanticVersion(r) {
			// The version for the path on the left side of `=>` will have an index less than 3
			if i < middle {
				info.Originalversion = r
				continue
			}
			// The version for the path on the right side of `=>` will have an index greater than 3
			info.TargetVersion = r
			continue
		}

		if i < middle {
			originalPath = r
			continue
		}
		info.TargetPath = r
	}

	replace[originalPath] = info
}

// replaceMods will update a module path and version if a matching path and version are found
func replaceMods(mods []*Module, replace map[string]replaceInfo) {
	for _, m := range mods {
		if val, ok := replace[m.Path]; ok {
			if val.Originalversion != "" && val.Originalversion != m.Version {
				continue
			}
			m.Path = val.TargetPath
			m.Version = val.TargetVersion
		}
	}
}
