package main

import (
	_ "music/internal/packed"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"

	"music/internal/cmd"
)

func main() {
	cmd, err := gcmd.NewFromObject(cmd.CMain{})
	if err != nil {
		panic(err)
	}
	cmd.Run(gctx.New())
}
