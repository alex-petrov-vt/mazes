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
	cmd := &cobra.Command{
		Use:   "binary-tree",
		Short: "Create a new maze using binary tree creation algoritm",
		Run: func(cmd *cobra.Command, args []string) {
			g := grid.NewGrid(5, 5)
			algorithms.CreateBinaryTreeMaze(g, time.Now().UnixNano())
			fmt.Printf("%s", g)
		},
	}

	return cmd
}
