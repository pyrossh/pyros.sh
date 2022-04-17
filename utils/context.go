package utils

import (
	"context"

	"github.com/pyros2097/gromer"
)

const base = "https://pyros.sh"

func GetUrl(c context.Context) string {
	url := gromer.GetUrl(c)
	return base + url.Path
}

func GetImageUrl(s string) string {
	return base + s
}
