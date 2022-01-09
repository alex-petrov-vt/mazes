package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

func newCreateCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new maze",
	}

    cmd.AddCommand(newCreateBinaryTreeCmd(out))
    cmd.AddCommand(newCreateSidewinderCmd(out))

	return cmd
}
