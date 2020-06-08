package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vindh1er/heimda11r-web/pkg/server"
)

var serverPort int
var serverBind string

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().IntVarP(&serverPort, "port", "p", 8080, "The port to serve the webserver on.")
	serverCmd.PersistentFlags().StringVarP(&serverBind, "bind", "b", "0.0.0.0", "The address to bind to.")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the webserver",
	Long:  `Run the webserver`,
	Run: func(cmd *cobra.Command, args []string) {
		server.StartServer(serverPort, serverBind)
	},
}
