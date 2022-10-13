package main

import (
	"context"
	"embed"
	"github.com/qunv/minitino/app"
	"github.com/qunv/minitino/app/config"
)

//go:embed _assets
//go:embed _templates
var fs embed.FS

func main() {
	ctx := context.Background()
	app.New(ctx, fs, config.LoadConfig()).Run()
}
