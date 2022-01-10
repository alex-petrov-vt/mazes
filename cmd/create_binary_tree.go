package cmd

import (
	"fmt"
	"io"
	"time"

	"github.com/alex-petrov-vt/mazes/algorithms"
	"github.com/alex-petrov-vt/mazes/grid"
	"github.com/spf13/cobra"
)

func newCreateBinaryTreeCmd(o createOpts, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "binary-tree",
		Short: "Create a new maze using binary tree creation algoritm",
		Run: func(cmd *cobra.Command, args []string) {
			g := grid.NewGrid(int(*o.height), int(*o.width))
			algorithms.CreateBinaryTreeMaze(g, time.Now().UnixNano())
            fmt.Fprintf(out, "%s", g)
		},
	}

	return cmd
}
