package main

import (
	"gopkg.in/urfave/cli.v1"

	emtUtils "ethermint/cmd/utils"
)

func resetCmd(ctx *cli.Context) error {
	return emtUtils.ResetAll(ctx)
}
