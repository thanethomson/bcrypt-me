package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var (
	flagCost int
)

var rootCmd = &cobra.Command{
	Use:   "bcrypt-me",
	Short: "A simple CLI for generating and validating bcrypt hashes",
}

var genCmd = &cobra.Command{
	Use:   "gen [password]",
	Short: "Generate a bcrypt hash of [password] with the given parameters",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hash, err := bcrypt.GenerateFromPassword([]byte(args[0]), flagCost)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(hash))
	},
}

var checkCmd = &cobra.Command{
	Use:   "check [hash] [password]",
	Short: "Validate the given password against the specified bcrypt hash",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := bcrypt.CompareHashAndPassword([]byte(args[0]), []byte(args[1]))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Password matches!")
	},
}

func main() {
	genCmd.PersistentFlags().IntVarP(&flagCost, "cost", "c", 12, fmt.Sprintf("the cost of the hash; min %d, max %d", bcrypt.MinCost, bcrypt.MaxCost))
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(checkCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
