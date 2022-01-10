package cmd

import (
	"fmt"
	"io"
	"time"

	"github.com/alex-petrov-vt/mazes/algorithms"
	"github.com/alex-petrov-vt/mazes/grid"
	"github.com/spf13/cobra"
)

func newCreateBinaryTreeCmd(out io.Writer) *cobra.Command {

    var height *uint
    var width *uint

	cmd := &cobra.Command{
		Use:   "binary-tree",
		Short: "Create a new maze using binary tree creation algoritm",
		Run: func(cmd *cobra.Command, args []string) {
			g := grid.NewGrid(int(*height), int(*width))
			algorithms.CreateBinaryTreeMaze(g, time.Now().UnixNano())
			fmt.Printf("%s", g)
		},
	}

    height = cmd.Flags().Uint("height", 10, "height of the maze to be created")
    width = cmd.Flags().Uint("width", 10, "width of the maze to be created")


	return cmd
}
