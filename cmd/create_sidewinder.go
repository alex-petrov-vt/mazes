package cmd

import (
	"fmt"
	"io"
	"time"

	"github.com/alex-petrov-vt/mazes/algorithms"
	"github.com/alex-petrov-vt/mazes/grid"
	"github.com/spf13/cobra"
)

func newCreateSidewinderCmd(o createOpts, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sidewinder",
		Short: "Create a new maze using sidewinder creation algorithm",
		Run: func(cmd *cobra.Command, args []string) {
			g := grid.NewGrid(int(*o.height), int(*o.width))
			algorithms.CreateSidewinderMaze(g, time.Now().UnixNano())
            fmt.Fprintf(out, "%s", g)
		},
	}

	return cmd
}
