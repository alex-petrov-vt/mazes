package cmd

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/alex-petrov-vt/mazes/algorithms"
	"github.com/alex-petrov-vt/mazes/distances"
	"github.com/alex-petrov-vt/mazes/grid"
	"github.com/spf13/cobra"
)

type createOpts struct {
	height        uint
	width         uint
	showDistances bool
	algorithm     string
}

func newCreateCmd(out io.Writer) *cobra.Command {

	o := createOpts{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new maze",
		Run: func(cmd *cobra.Command, args []string) {
			g := grid.NewGrid(int(o.height), int(o.width))

			switch o.algorithm {
			case algorithms.BinaryTree.String():
				algorithms.CreateBinaryTreeMaze(g, time.Now().UnixNano())
			case algorithms.Sidewinder.String():
				algorithms.CreateSidewinderMaze(g, time.Now().UnixNano())
			default:
				fmt.Fprintf(out, "unrecognized algorithm: %s. Allowed values: %s. Using %s instead",
					o.algorithm, strings.Join(algorithms.Algorithms(), ", "), algorithms.BinaryTree)
			}

			if o.showDistances {
				dg := distances.NewDistanceGrid(g)
				d := distances.FindDistances(dg.Grid.Cells[0][0])
				dg.Distances = d
				fmt.Fprintf(out, "%s", dg)
			} else {
				fmt.Fprintf(out, "%s", g)
			}
		},
	}

	cmd.Flags().UintVar(&o.height, "height", 10, "height of the maze to be created")
	cmd.Flags().UintVar(&o.width, "width", 10, "width of the maze to be created")
	cmd.Flags().BoolVarP(&o.showDistances, "distances", "d", false, "show distances on the maze")
	cmd.Flags().StringVarP(&o.algorithm, "algo", "a", string(algorithms.BinaryTree),
		fmt.Sprintf("algorithm to use to generate the maze. Allowed values: %s", strings.Join(algorithms.Algorithms(), ", ")))

	return cmd
}
