/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "minesweeper-cli",
	Short: "A deterministic minesweeper CLI utility.",
	Long: `You can simulate a game of minesweeper passing a seed or taking a generated seed by the program.
	Example usage:
		minesweeper-cli -seed 123456789 -w 12 -h 24 -act "<actions>"
	`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("seed", "s", "", "Seed to fill the randomness")
	rootCmd.Flags().StringP("actions", "a", "", "List of interactions with the current game")
	rootCmd.Flags().Int("width", 9, "Width of the board")
	rootCmd.Flags().Int("height", 9, "Height of the board")
	rootCmd.Flags().IntP("mines", "m", 10, "Total amount of mines for the game")
}
