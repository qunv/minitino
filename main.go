package main

import (
	"context"
	"embed"
	"github.com/qunv/minitino/app"
)

//go:embed _assets
//go:embed _templates
var fs embed.FS

func main() {
	ctx := context.Background()
	app.New(ctx, fs).Run()
}
