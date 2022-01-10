package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

type createOpts struct {
    height *uint
    width *uint
}

func newCreateCmd(out io.Writer) *cobra.Command {

    o := createOpts{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new maze",
	}

    o.height = cmd.PersistentFlags().Uint("height", 10, "height of the maze to be created")
    o.width = cmd.PersistentFlags().Uint("width", 10, "width of the maze to be created")

    cmd.AddCommand(newCreateBinaryTreeCmd(o, out))
    cmd.AddCommand(newCreateSidewinderCmd(o, out))


	return cmd
}
