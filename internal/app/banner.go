package app

import (
	"errors"
	"os"
)

func getBanner(bannerPath string) string {
	if _, err := os.Stat(bannerPath); errors.Is(err, os.ErrNotExist) {
		return "" // No banner.
	}

	banner, err := os.ReadFile(bannerPath)
	if err != nil {
		// Errors don't really matter; the hoster will just have to figure out the problem.
		return ""
	}

	return string(banner)
}
