package main

import (
	"os"

	"github.com/grafana/grok/cmd/generate"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:          "grok <command>",
		Short:        "A tool for working with Grafana objects from code",
		SilenceUsage: true,
	}

	rootCmd.AddCommand(generate.Command())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
