package model

import "strings"

type StoreImageMode int

const (
	None StoreImageMode = iota
	File
	GCS
)

func GetStoreImageModeFromString(src string) (mode StoreImageMode) {
	formattedSrc := strings.TrimSpace(strings.ToLower(src))

	if formattedSrc == "file" {
		return File
	}
	if formattedSrc == "gcs" {
		return GCS
	}

	return None
}
