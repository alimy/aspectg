package version

import (
	"fmt"
	"github.com/alimy/aspectg/cmd"
	"github.com/spf13/cobra"
)

var (
	Version = "0.0.0"

	// GitHash Value will be set during build
	GitHash = "Not provided"

	// BuildTime Value will be set during build
	BuildTime = "Not provided"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "version of application",
		Long:  "version infomation for application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s \nBuildTime:%s\nGitHash:%s\n", Version, BuildTime, GitHash)
		},
	}

	// Register versionCmd as sub-command
	cmd.Register(versionCmd)
}
