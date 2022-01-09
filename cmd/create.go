package cmd

import (
	"fmt"
	"time"

	"github.com/alex-petrov-vt/mazes/algorithms"
	"github.com/alex-petrov-vt/mazes/grid"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new maze",
	Run: func(cmd *cobra.Command, args []string) {
		g := grid.NewGrid(5,5)
        algorithms.CreateBinaryTreeMaze(g, time.Now().UnixNano())
        fmt.Printf("%s", g)
	},
}
