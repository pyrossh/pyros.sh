package utils

import (
	"context"

	"github.com/pyros2097/gromer"
)

func GetUrl(c context.Context) string {
	url := gromer.GetUrl(c)
	return "https://pyros.sh" + url.Path
}
