package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xie392/restful-api/version"
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "host-api",
	Short: "host-api 后端API",
	Long:  "host-api 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "print host-api version")
}
