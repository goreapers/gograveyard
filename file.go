package gograveyard

import (
	"fmt"
	"os"
)

func GoModBytes(pathToGoMod string) ([]byte, error) {
	bytes, err := os.ReadFile(pathToGoMod)
	if err != nil {
		return nil, fmt.Errorf("failed reading file: %w", err)
	}

	return bytes, nil
}
