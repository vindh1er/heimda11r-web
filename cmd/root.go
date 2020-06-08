package cmd

import (
	"flag"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/klog"
)

var (
	version string
	commit  string
)

func init() {
	klog.InitFlags(nil)
	flag.Parse()
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}

var rootCmd = &cobra.Command{
	Use:   "heimda11r-web",
	Short: "heimda11r-web",
	Long:  `A webserver. Mostly`,
	Run: func(cmd *cobra.Command, args []string) {
		klog.Error("You must specify a sub-command.")
		err := cmd.Help()
		if err != nil {
			klog.Error(err)
		}
		os.Exit(1)
	},
}

// Execute the stuff
func Execute(VERSION string, COMMIT string) {
	version = VERSION
	commit = COMMIT
	if err := rootCmd.Execute(); err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}
