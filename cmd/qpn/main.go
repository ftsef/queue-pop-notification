package main

import (
	"context"
	"fmt"
	"os"
	"queue-pop-notification/internal/service"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "qpn",
		Short: "qpn - Queue Pop Notification - Notifies when a queue pop is detected",
		Long:  `Sends a Discord notification when a queue pop is detected in World of Warcraft.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			service.Run(context.Background(), cfgFile)
		},
	}
)

func main() {

	cobra.MousetrapHelpText = ""
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "Path to the configuration file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
