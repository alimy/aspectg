package cmd

import (
	"github.com/alimy/aspectg/cmd"
	"github.com/alimy/aspectg/module/build"
	"github.com/spf13/cobra"
	"github.com/unisx/logus"
)

func init() {
	buildCmd := &cobra.Command{
		Use:   "build",
		Short: "Generate AspectG AOP code",
		Long:  "build will generate AspectG defined code",
		Run:   buildRun,
	}

	// Register helloCmd as sub-command
	cmd.Register(buildCmd)
}

func buildRun(cmd *cobra.Command, args []string) {
	logus.InDevelopment()
	build.StartBuild()
}
