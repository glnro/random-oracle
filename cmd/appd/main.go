package main

import (
	"fmt"
	"os"

	"github.com/glnro/oracle/app"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/glnro/oracle/cmd/appd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
