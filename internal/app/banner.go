package app

import (
	"os"
	"errors"
)

func getBanner() string {
	if _, err := os.Stat("./banner.txt"); errors.Is(err, os.ErrNotExist) {
		return "" // No banner.
	}

	banner, err := os.ReadFile("./banner.txt")
	if err != nil {
		// Errors don't really matter; the hoster will just have to figure out the problem.
		return ""
	}

	return string(banner)
}
