package main

import (
	"context"
	"github.com/qunv/minitino/app"
)

func main() {
	ctx := context.Background()
	app.New(ctx).Run()
}
