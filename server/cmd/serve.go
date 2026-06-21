package cmd

import (
	"fmt"

	"github.com/Chad-Glazier/controlward/httpserver"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the http server",
	Long: `Starts an HTTP server for control ward. Upon starting, a URL to the 
API documentation will be logged.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !httpserver.ValidPort(httpserver.Port) {
			return fmt.Errorf(
				"--port argument %s is not a valid port number",
				httpserver.Port,
			)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		httpserver.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(
		&httpserver.Port,
		"port", "p",
		httpserver.Port,
		"set the port to listen on",
	)
}
