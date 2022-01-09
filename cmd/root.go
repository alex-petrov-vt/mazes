package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func newRootCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mazy",
		Short: "Mazy is a cli tool to explore concepts and algorithms related to mazes",
	}

    cmd.AddCommand(newCreateCmd(out))

    return cmd
}

func Execute() {
	if err := newRootCmd(os.Stdout).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
