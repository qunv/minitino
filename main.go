package main

import (
	"context"
	"github.com/qunv/minitino/boot"
)

func main() {
	ctx := context.Background()
	boot.New(ctx).Run()
}
