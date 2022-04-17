package utils

import (
	"context"

	"github.com/pyros2097/gromer"
	"pyros.sh/assets"
)

const base = "https://pyros.sh"

func GetUrl(c context.Context) string {
	url := gromer.GetUrl(c)
	return base + url.Path
}

func GetImageUrl(path string) string {
	return base + gromer.GetAssetUrl(assets.FS, path)
}
